# car-project

## Changelog

- **v1**: developing

## Description

This is a personal side project to record vehicle information.

Using Technical:

- Golang
  - [Go Clean Architecture](https://github.com/bxcodec/go-clean-arch)
    - [Gin-Gonic](https://gin-gonic.com/)
    - [GORM](https://gorm.io)
  - [Mockery](https://github.com/vektra/mockery)
  - [Swaggo](https://github.com/swaggo/swag)
  - [Golangci-lint](https://github.com/golangci/golangci-lint)
- Typescript
  - [Quasar SPA Application](https://quasar.dev/quasar-cli-webpack/developing-spa/introduction)
- Mysql 5.7

## System Requirement

- Golang 1.16+
- Node 12+
- Quasar CLI 1.3.2+

\*Make Sure you have set **GOROOT**/bin, **GOPATH**/bin into your env **PATH\***

## Install

```bash
$ make install
# Includes
  $ go get
  $ make mockery-install
  $ make swaggo-install
  $ make lint-install
  $ make frontend-install
```

## Test Golang

```bash
# quick test
$ make test
# test and open cover report
$ make test-cover
```

## Before Run

```bash
$ make before-run-backend
# Includes
  $ make mockery
  $ make swaggo
  $ make lint
  $ make test

$ make before-run-frontend
# Includes
  $ make frontend-build
```

## Run Application

```bash
$ make run
```

## Deploy To FreeBSD@RPi3

```bash
$ make build-freebsd-rpi3
$ scp dist/car-project_freebsd_rpi3 => REMOTE_SERVER
$ scp config.yaml => REMOTE_SERVER
$ scp app/.quasar-project/dist/spa/* => REMOTE_SERVER/www
```
