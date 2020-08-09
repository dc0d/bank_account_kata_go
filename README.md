# Bank Account Kata in Go

Personal notes on practicing [Bank Account Kata](https://github.com/sandromancuso/Bank-kata) in [Go](https://golang.org/).

> Status: Parked 🚧
>
> _The focus goal here was only exploration and playing around with [IDD (Interaction-Driven Design)](https://codurance.com/2017/12/08/introducing-idd/) and proxy-based mocks - no dedication to anything other than exploration and only sketches and drafts_.

----


## Prerequisites

Install [golangci-lint](https://github.com/golangci/golangci-lint).

```
$ go get -u github.com/onsi/ginkgo/ginkgo
$ go get -u github.com/matryer/moq
$ go mod tidy
```


## Run Tests

Generate mocks:

```
$ make gen
```

Run Unit Tests:

```
$ make
```


## Linting

Linting needs `golangci-lint` to be installed.

```
$ make lint
```
