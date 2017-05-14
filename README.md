# cepinator

Brazilian zip codes (CEP) microservice.

It gets CEP's from [viacep] and cache them in a redis store.
It's simple and really fast.

You can run it within Docker:

```console
docker -d run -p 3000:3000 caarlos0/cepinator
```

Or download a binary from [releases](https://github.com/caarlos0/cepinator/releases)
and execute it.

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.
We appreciate your contribution. Please refer to our [contributing guidelines](CONTRIBUTING.md).

[![Release](https://img.shields.io/github/release/caarlos0/cepinator.svg?style=flat-square)](https://github.com/caarlos0/cepinator/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/caarlos0/cepinator.svg?style=flat-square)](https://travis-ci.org/caarlos0/cepinator)
[![Coverage Status](https://img.shields.io/codecov/c/github/caarlos0/cepinator/master.svg?style=flat-square)](https://codecov.io/gh/caarlos0/cepinator)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/caarlos0/cepinator)
[![Go Report Card](https://goreportcard.com/badge/github.com/caarlos0/cepinator?style=flat-square)](https://goreportcard.com/report/github.com/caarlos0/cepinator)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

## Configuration

Configuration is done via environment variables:

- `PORT`: port to bind to, defaults to 3000;
- `REDIS_URL`: redis URL to use, defaults to `:6379`.

[viacep]: http://viacep.com.br/
