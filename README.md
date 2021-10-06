[![Build Status](https://app.travis-ci.com/DragonSSS/simple-go-starter.svg?branch=master)](https://app.travis-ci.com/DragonSSS/simple-go-starter)

# simple-go-starter

## Intro

## Architecture

## APIs

The Go swaggo plugin is used to generate API doc with annotations in the code.

The yaml of swagger api file is at [swagger.yaml](https://github.com/DragonSSS/simple-go-starter/blob/master/docs/swagger.yaml). Please review it with [online swagger editor](https://editor.swagger.io/)

## Infrastructure

[Makefile](https://github.com/DragonSSS/simple-go-starter/blob/master/Makefile) takes care of:

* Run code linter with golangci-lint docker
* Download dependencies using Go modules
* Build binary into `build/bin/`
* Build docker image using [Dockerfile](https://github.com/DragonSSS/simple-go-starter/blob/master/Dockerfile)
* Run unit test and generate test coverage report into coverage.out
* Generate swagger API doc
* Clean the compiled binary

The repo is integrated with Travis CI pipeline with [travis.yml](https://github.com/DragonSSS/simple-go-starter/blob/master/.travis.yml), which supports stages:

* Lint (make lint)
* Unit test (make test)
* Build binary (make build)
* Deploy the local built image on k8s cluster, check health of web service
  * Install k8s cluster on the fly using k8s Kind (running k8s cluster into container)
  * Build go-starter docker image locally
  * Load local built image into k8s cluster
  * Deploy go-starter app
  * Check the log of running pod
  * Deploy a k8s service for go-starter
  * Forward local traffic into the k8s service by kubectl port-forward
  * Check health of go-starter by curl
