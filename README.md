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

## Configuration

Configuration is done via environment variables:

- `PORT`: port to bind to, defaults to 3000;
- `REDIS_URL`: redis URL to use, defaults to `:6379`.
