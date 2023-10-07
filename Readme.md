# Sponsors API

- [Sponsors API](#sponsors-api)
  - [Deploying](#deploying)
  - [Usage](#usage)
    - [`/markdown`](#markdown)
    - [`/json`](#json)
    - [`/txt`](#txt)
    - [`/`](#)


Sponsors API is a [GitHub Sponsors](https://github.com/sponsors) server for displaying your current sponsor avatars in your project Readme. It looks like this:

[<img src="https://sponsors.mranftl.com/avatar/0" width="35">](https://sponsors.mranftl.com/profile/0)
[<img src="https://sponsors.mranftl.com/avatar/1" width="35">](https://sponsors.mranftl.com/profile/1)
[<img src="https://sponsors.mranftl.com/avatar/2" width="35">](https://sponsors.mranftl.com/profile/2)
[<img src="https://sponsors.mranftl.com/avatar/3" width="35">](https://sponsors.mranftl.com/profile/3)
[<img src="https://sponsors.mranftl.com/avatar/4" width="35">](https://sponsors.mranftl.com/profile/4)
[<img src="https://sponsors.mranftl.com/avatar/5" width="35">](https://sponsors.mranftl.com/profile/5)
[<img src="https://sponsors.mranftl.com/avatar/6" width="35">](https://sponsors.mranftl.com/profile/6)
[<img src="https://sponsors.mranftl.com/avatar/7" width="35">](https://sponsors.mranftl.com/profile/7)
[<img src="https://sponsors.mranftl.com/avatar/8" width="35">](https://sponsors.mranftl.com/profile/8)
[<img src="https://sponsors.mranftl.com/avatar/9" width="35">](https://sponsors.mranftl.com/profile/9)
[<img src="https://sponsors.mranftl.com/avatar/10" width="35">](https://sponsors.mranftl.com/profile/10)
[<img src="https://sponsors.mranftl.com/avatar/11" width="35">](https://sponsors.mranftl.com/profile/11)
[<img src="https://sponsors.mranftl.com/avatar/12" width="35">](https://sponsors.mranftl.com/profile/12)
[<img src="https://sponsors.mranftl.com/avatar/13" width="35">](https://sponsors.mranftl.com/profile/13)
[<img src="https://sponsors.mranftl.com/avatar/14" width="35">](https://sponsors.mranftl.com/profile/14)
[<img src="https://sponsors.mranftl.com/avatar/15" width="35">](https://sponsors.mranftl.com/profile/15)
[<img src="https://sponsors.mranftl.com/avatar/16" width="35">](https://sponsors.mranftl.com/profile/16)
[<img src="https://sponsors.mranftl.com/avatar/17" width="35">](https://sponsors.mranftl.com/profile/17)
[<img src="https://sponsors.mranftl.com/avatar/18" width="35">](https://sponsors.mranftl.com/profile/18)
[<img src="https://sponsors.mranftl.com/avatar/19" width="35">](https://sponsors.mranftl.com/profile/19)
[<img src="https://sponsors.mranftl.com/avatar/20" width="35">](https://sponsors.mranftl.com/profile/20)
[<img src="https://sponsors.mranftl.com/avatar/21" width="35">](https://sponsors.mranftl.com/profile/21)
[<img src="https://sponsors.mranftl.com/avatar/22" width="35">](https://sponsors.mranftl.com/profile/22)
[<img src="https://sponsors.mranftl.com/avatar/23" width="35">](https://sponsors.mranftl.com/profile/23)
[<img src="https://sponsors.mranftl.com/avatar/24" width="35">](https://sponsors.mranftl.com/profile/24)
[<img src="https://sponsors.mranftl.com/avatar/25" width="35">](https://sponsors.mranftl.com/profile/25)
[<img src="https://sponsors.mranftl.com/avatar/26" width="35">](https://sponsors.mranftl.com/profile/26)
[<img src="https://sponsors.mranftl.com/avatar/27" width="35">](https://sponsors.mranftl.com/profile/27)
[<img src="https://sponsors.mranftl.com/avatar/28" width="35">](https://sponsors.mranftl.com/profile/28)
[<img src="https://sponsors.mranftl.com/avatar/29" width="35">](https://sponsors.mranftl.com/profile/29)
[<img src="https://sponsors.mranftl.com/avatar/30" width="35">](https://sponsors.mranftl.com/profile/30)
[<img src="https://sponsors.mranftl.com/avatar/31" width="35">](https://sponsors.mranftl.com/profile/31)
[<img src="https://sponsors.mranftl.com/avatar/32" width="35">](https://sponsors.mranftl.com/profile/32)
[<img src="https://sponsors.mranftl.com/avatar/33" width="35">](https://sponsors.mranftl.com/profile/33)
[<img src="https://sponsors.mranftl.com/avatar/34" width="35">](https://sponsors.mranftl.com/profile/34)
[<img src="https://sponsors.mranftl.com/avatar/35" width="35">](https://sponsors.mranftl.com/profile/35)
[<img src="https://sponsors.mranftl.com/avatar/36" width="35">](https://sponsors.mranftl.com/profile/36)
[<img src="https://sponsors.mranftl.com/avatar/37" width="35">](https://sponsors.mranftl.com/profile/37)
[<img src="https://sponsors.mranftl.com/avatar/38" width="35">](https://sponsors.mranftl.com/profile/38)
[<img src="https://sponsors.mranftl.com/avatar/39" width="35">](https://sponsors.mranftl.com/profile/39)
[<img src="https://sponsors.mranftl.com/avatar/40" width="35">](https://sponsors.mranftl.com/profile/40)
[<img src="https://sponsors.mranftl.com/avatar/41" width="35">](https://sponsors.mranftl.com/profile/41)
[<img src="https://sponsors.mranftl.com/avatar/42" width="35">](https://sponsors.mranftl.com/profile/42)
[<img src="https://sponsors.mranftl.com/avatar/43" width="35">](https://sponsors.mranftl.com/profile/43)
[<img src="https://sponsors.mranftl.com/avatar/44" width="35">](https://sponsors.mranftl.com/profile/44)
[<img src="https://sponsors.mranftl.com/avatar/45" width="35">](https://sponsors.mranftl.com/profile/45)
[<img src="https://sponsors.mranftl.com/avatar/46" width="35">](https://sponsors.mranftl.com/profile/46)
[<img src="https://sponsors.mranftl.com/avatar/47" width="35">](https://sponsors.mranftl.com/profile/47)
[<img src="https://sponsors.mranftl.com/avatar/48" width="35">](https://sponsors.mranftl.com/profile/48)
[<img src="https://sponsors.mranftl.com/avatar/49" width="35">](https://sponsors.mranftl.com/profile/49)
[<img src="https://sponsors.mranftl.com/avatar/50" width="35">](https://sponsors.mranftl.com/profile/50)
[<img src="https://sponsors.mranftl.com/avatar/51" width="35">](https://sponsors.mranftl.com/profile/51)
[<img src="https://sponsors.mranftl.com/avatar/52" width="35">](https://sponsors.mranftl.com/profile/52)
[<img src="https://sponsors.mranftl.com/avatar/53" width="35">](https://sponsors.mranftl.com/profile/53)
[<img src="https://sponsors.mranftl.com/avatar/54" width="35">](https://sponsors.mranftl.com/profile/54)
[<img src="https://sponsors.mranftl.com/avatar/55" width="35">](https://sponsors.mranftl.com/profile/55)
[<img src="https://sponsors.mranftl.com/avatar/56" width="35">](https://sponsors.mranftl.com/profile/56)
[<img src="https://sponsors.mranftl.com/avatar/57" width="35">](https://sponsors.mranftl.com/profile/57)
[<img src="https://sponsors.mranftl.com/avatar/58" width="35">](https://sponsors.mranftl.com/profile/58)
[<img src="https://sponsors.mranftl.com/avatar/59" width="35">](https://sponsors.mranftl.com/profile/59)
[<img src="https://sponsors.mranftl.com/avatar/60" width="35">](https://sponsors.mranftl.com/profile/60)
[<img src="https://sponsors.mranftl.com/avatar/61" width="35">](https://sponsors.mranftl.com/profile/61)
[<img src="https://sponsors.mranftl.com/avatar/62" width="35">](https://sponsors.mranftl.com/profile/62)
[<img src="https://sponsors.mranftl.com/avatar/63" width="35">](https://sponsors.mranftl.com/profile/63)
[<img src="https://sponsors.mranftl.com/avatar/64" width="35">](https://sponsors.mranftl.com/profile/64)
[<img src="https://sponsors.mranftl.com/avatar/65" width="35">](https://sponsors.mranftl.com/profile/65)
[<img src="https://sponsors.mranftl.com/avatar/66" width="35">](https://sponsors.mranftl.com/profile/66)
[<img src="https://sponsors.mranftl.com/avatar/67" width="35">](https://sponsors.mranftl.com/profile/67)
[<img src="https://sponsors.mranftl.com/avatar/68" width="35">](https://sponsors.mranftl.com/profile/68)

## Deploying

We build and publish a minimal docker image via GitHub Actions.

The following environment variables are supported:

- `GITHUB_TOKEN` the GitHub API token (org read scope is required)
- `PORT` the server port (defaults to 3000)
- `URL` the url to your endpoint such as "https://sponsors.myhost.com" (optional)
- `CACHE_TTL` the cache TTL (go duration, sponsors are cached for an hour by default (`1h`))

## Usage

* Visit the `/markdown` path for the markdown to copy/paste into readmes
* For convenience, the following additional debug endpoints are provided:
  * `/` serves a minimal HTML with imgs and links to the sponsor profiles
  * `/json` serves the current sponsors as JSON like this: ``````

### `/markdown`

Visit the `/markdown` path for the markdown to copy/paste into readmes

### `/json`

```json
{
    "login": "your_username",
    "sponsors": [
        {
            "login": "your_sponsor",
            "url": "https://..."
        },
        [...]
    ]
}
```

### `/txt`

List in the following format:
```txt
<login> <url>
```

### `/`

Serves a minimal HTML with sponsors images + links to the sponsor profiles