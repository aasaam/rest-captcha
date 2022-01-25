<div align="center">
  <h1>
    REST Captcha
  </h1>
  <p>
    Simple in memory multi language captcha generator server
  </p>
  <p>
    <a href="https://github.com/aasaam/rest-captcha/actions/workflows/build.yml" target="_blank"><img src="https://github.com/aasaam/rest-captcha/actions/workflows/build.yml/badge.svg" alt="build" /></a>
    <a href="https://goreportcard.com/report/github.com/aasaam/rest-captcha"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/aasaam/rest-captcha"></a>
    <a href="https://hub.docker.com/r/aasaam/rest-captcha" target="_blank"><img src="https://img.shields.io/docker/image-size/aasaam/rest-captcha?label=docker%20image" alt="docker" /></a>
    <a href="https://github.com/aasaam/rest-captcha/actions/workflows/docs.yml" target="_blank"><img src="https://github.com/aasaam/rest-captcha/actions/workflows/docs.yml/badge.svg" alt="docs" /></a>
    <a href="https://github.com/aasaam/rest-captcha/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/aasaam/rest-captcha"></a>
  </p>
</div>

## Guide

For see available options

```bash
$ docker run --rm aasaam/rest-captcha -h

# Usage of ./rest-captcha:
#   -auth-password string
#         Basic authentication password
#   -auth-username string
#         Basic authentication username
#   -base-url string
#         Base URL for routes (default "/")
#   -listen string
#         Application listen address (default "0.0.0.0:4000")
#   -return-value
#         Return value on generation
```

It's generate captcha image via ID, base64 encoded image and value of captcha:

```bash
curl -X POST -H 'Content-type: application/json' -d '{"lang":"fa","ttl":30, "level": "1", "quality": 10}' http://rest-captcha:4000/new
```

```bash
curl -X POST -H 'Content-type: application/json' -d '{"id":"UNIQUE_IDENTIFIER","value":999999}' http://rest-captcha:4000/solve
```

## Languages

Currently following language are supported:

- `en` English (It's default/fallback language for invalid language code)
- `fa` Persian
- `ar` Arabic

<div>
  <p align="center">
    <a href="https://aasaam.com" title="aasaam software development group">
      <img alt="aasaam software development group" width="64" src="https://raw.githubusercontent.com/aasaam/information/master/logo/aasaam.svg">
    </a>
    <br />
    aasaam software development group
  </p>
</div>
