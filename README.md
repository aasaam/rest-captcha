<div align="center">
  <h1>
    REST Captcha
  </h1>
  <p>
    Simple in memory multi language captcha generator server
  </p>
  <p>
    <a href="https://gitlab.com/aasaam/rest-captcha/-/pipelines"><img alt="CI Pipeline" src="https://gitlab.com/aasaam/rest-captcha/badges/master/pipeline.svg"></a>
    <a href="https://gitlab.com/aasaam/rest-captcha/"><img alt="Code Coverage" src="https://gitlab.com/aasaam/rest-captcha/badges/master/coverage.svg"></a>
    <a href="https://goreportcard.com/report/github.com/aasaam/rest-captcha"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/aasaam/rest-captcha"></a>
    <a href="https://hub.docker.com/r/aasaam/rest-captcha" target="_blank"><img src="https://img.shields.io/docker/image-size/aasaam/rest-captcha?label=docker%20image" alt="docker" /></a>
    <a href="https://quay.io/repository/aasaam/rest-captcha" target="_blank"><img src="https://img.shields.io/badge/docker%20image-quay.io-blue" alt="quay.io" /></a>
    <a href="https://github.com/aasaam/rest-captcha/actions/workflows/docs.yml" target="_blank"><img src="https://github.com/aasaam/rest-captcha/actions/workflows/docs.yml/badge.svg" alt="docs" /></a>
    <a href="https://github.com/aasaam/rest-captcha/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/aasaam/rest-captcha"></a>
  </p>
</div>

## Guide

For see available options

```bash
docker run --rm aasaam/rest-captcha -h
# will show available options
```

It's generate captcha image via ID, base64 encoded image and value of captcha:

```bash
curl -X POST -H 'Content-type: application/json' -d '{"lang":"fa","ttl":30, "level": "1"}' http://rest-captcha:4000/new
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
