// Package sponsors provides GitHub sponsors management.
package sponsors

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shurcooL/githubv4"
)

// pixel is a png used for missing avatars.
var pixel []byte

// initialize gray pixel for missing avatar responses.
func init() {
	var buf bytes.Buffer
	r := image.Rect(0, 0, 1, 1)
	img := image.NewRGBA(r)
	c := color.RGBA{0xF8, 0xF9, 0xFA, 0xFF}
	draw.Draw(img, r, &image.Uniform{c}, image.ZP, draw.Src)
	png.Encode(&buf, img)
	pixel = buf.Bytes()
}

// Sponsor model.
type SponsorEntity struct {
	Typename string `graphql:"typename :__typename"`

	User struct {
		// Login name of the sponsor.
		Login string

		// AvatarURL of the sponsor.
		AvatarURL string
	} `graphql:"... on User"`

	Organization struct {
		// Login name of the sponsor.
		Login string

		// AvatarURL of the sponsor.
		AvatarURL string
	} `graphql:"... on Organization"`
}

func (se SponsorEntity) GetLogin() string {
	if se.Typename == "User" {
		return se.User.Login
	}
	return se.Organization.Login
}

func (se SponsorEntity) GetAvatarURL() string {
	if se.Typename == "User" {
		return se.User.AvatarURL
	}
	return se.Organization.AvatarURL
}

// Server manager.
type Server struct {
	// URL is the url of the server.
	URL string

	// Client is the github client.
	Client *githubv4.Client

	// CacheTTL is the duration until the cache expires.
	CacheTTL time.Duration

	// cache
	mu             sync.Mutex
	cacheTimestamp time.Time
	cache          []SponsorEntity

	// viewer login name (the user who's token is used to fetch the sponsors)
	login string
}

// ServeHTTP implementation.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	path := r.URL.Path

	// logging
	start := time.Now()
	log.Printf("%s %s", r.Method, path)
	defer func() {
		log.Printf("%s %s -> %s", r.Method, path, time.Since(start))
	}()

	// prime cache
	err := s.primeCache(ctx)
	if err != nil {
		log.Printf("error priming cache: %s", err)
		http.Error(w, "Error fetching sponsors", http.StatusInternalServerError)
		return
	}

	// routing
	switch {
	case path == "/":
		s.serveHTML(w, r)
	case strings.HasPrefix(path, "/avatar"):
		s.serveAvatar(w, r)
	case strings.HasPrefix(path, "/profile"):
		s.serveProfile(w, r)
	case strings.HasPrefix(path, "/txt"):
		s.serveTXT(w, r)
	case strings.HasPrefix(path, "/json"):
		s.serveJSON(w, r)
	case strings.HasPrefix(path, "/markdown"):
		s.serveMarkdown(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotImplemented)
	}
}

// serveMarkdown serves a list of markdown links which you can copy/paste into your Readme.
func (s *Server) serveMarkdown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/markdown")
	for i := 0; i < 100; i++ {
		fmt.Fprintf(w, `[<img src="%s/avatar/%d" width="35">](%s/profile/%d)`, s.URL, i, s.URL, i)
		fmt.Fprintf(w, "\n")
	}
}

// serveMarkdown serves a list of markdown links which you can copy/paste into your Readme.
func (s *Server) serveTXT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	for _, se := range s.cache {
		fmt.Fprintf(w, `%s %s`, se.GetLogin(), se.GetAvatarURL())
		fmt.Fprintf(w, "\n")
	}
}

func (s *Server) serveHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>sponsors</title>
  </head>
  <body>`)

	for _, se := range s.cache {
		fmt.Fprintf(w, `<a href="https://github.com/%s/" target="_blank"><img src="%s" alt="%s" width="35" /></a>`, se.GetLogin(), se.GetAvatarURL(), se.GetLogin())
	}

	fmt.Fprintf(w, `</body>
</html>`)
}

func (s *Server) serveJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"login":"%s","sponsors":[`, s.login)

	for i, se := range s.cache {
		fmt.Fprintf(w, `{"login":"%s","url":"%s"}`, se.GetLogin(), se.GetAvatarURL())
		if i < len(s.cache)-1 {
			fmt.Fprintf(w, ",")
		}
	}
	fmt.Fprintf(w, "]}")
}

// serveAvatar redirects to a sponsor's avatar image.
func (s *Server) serveAvatar(w http.ResponseWriter, r *http.Request) {
	// /avatar/{index}
	index := strings.Replace(r.URL.Path, "/avatar/", "", 1)
	n, err := strconv.Atoi(index)
	if err != nil {
		log.Printf("error parsing index: %s", err)
		http.Error(w, "Sponsor index must be a number", http.StatusBadRequest)
		return
	}

	// check index bounds
	if n > len(s.cache)-1 {
		w.Header().Set("Content-Type", "image/png")
		io.Copy(w, bytes.NewReader(pixel))
		return
	}

	// redirect to avatar
	sponsor := s.cache[n]
	w.Header().Set("Location", sponsor.GetAvatarURL())
	w.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprintf(w, "Redirecting to %s", sponsor.GetAvatarURL())
}

// serveProfile redirects to a sponsor's profile.
func (s *Server) serveProfile(w http.ResponseWriter, r *http.Request) {
	// /profile/{index}
	index := strings.Replace(r.URL.Path, "/profile/", "", 1)
	n, err := strconv.Atoi(index)
	if err != nil {
		log.Printf("error parsing index: %s", err)
		http.Error(w, "Sponsor index must be a number", http.StatusBadRequest)
		return
	}

	// check index bounds
	if n > len(s.cache)-1 {
		// redirect to viewer sponsor profile site
		url := fmt.Sprintf("https://github.com/sponsors/%s", s.login)
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "Redirecting to %s", url)
		return
	}

	// redirect to the sponsors profile page
	sponsor := s.cache[n]
	url := fmt.Sprintf("https://github.com/%s", sponsor.GetLogin())
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprintf(w, "Redirecting to %s", url)
}

// primeCache implementation.
func (s *Server) primeCache(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// check ttl
	if time.Since(s.cacheTimestamp) <= s.CacheTTL {
		return nil
	}

	// fetch
	log.Printf("cache miss, fetching sponsors")
	sponsors, login, err := s.getSponsors(ctx)
	if err != nil {
		return err
	}

	s.cache = sponsors
	s.cacheTimestamp = time.Now()
	s.login = login
	return nil
}

// getSponsors implementation.
func (s *Server) getSponsors(ctx context.Context) ([]SponsorEntity, string /* viewer login name */, error) {
	var sponsors []SponsorEntity
	var q sponsorships
	var cursor string

	for {
		err := s.Client.Query(ctx, &q, map[string]interface{}{
			"cursor": githubv4.String(cursor),
		})

		if err != nil {
			return nil, "", err
		}

		for _, edge := range q.Viewer.SponsorshipsAsMaintainer.Edges {
			sponsor := edge.Node.SponsorEntity
			sponsors = append(sponsors, sponsor)
		}

		if !q.Viewer.SponsorshipsAsMaintainer.PageInfo.HasNextPage {
			break
		}

		cursor = q.Viewer.SponsorshipsAsMaintainer.PageInfo.EndCursor
	}

	return sponsors, q.Viewer.Login, nil
}

// sponsorships query.
type sponsorships struct {
	Viewer struct {
		Login                    string
		SponsorshipsAsMaintainer struct {
			PageInfo struct {
				EndCursor   string
				HasNextPage bool
			}

			Edges []struct {
				Node struct {
					SponsorEntity SponsorEntity
				}
				Cursor string
			}
		} `graphql:"sponsorshipsAsMaintainer(first: 100, after: $cursor)"`
	}
}
