[![Build Status](https://travis-ci.org/DragonSSS/simple-go-starter.svg?branch=master)](https://travis-ci.org/DragonSSS/simple-go-starter)

# simple-go-starter

## Intro

## Infrastructure

[Makefile](https://github.com/DragonSSS/simple-go-starter/blob/master/Makefile) takes care of:

* Run code linter with golangci-lint docker
* Download dependencies using Go modules
* Build binary into build/bin/
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
  * Build cloud audition docker image locally
  * Load local built image into k8s cluster
  * Deploy cloud audition app
  * Check the log of running pod
  * Deploy a k8s service for cloud audition app
  * Forward local traffic into the k8s service by kubectl port-forward
  * Check health of app by curl
  