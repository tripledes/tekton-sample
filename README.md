# Tekton sample

Application to learn and demo [Tekton](https://tekton.dev) pipelines

## Building

```shell
$ go test ./pkg/api && go build
```

## Running it locally

```shell
$ podman-compose up --force-recreate --abort-on-container-exit --build -d
...
$ curl localhost:8080/quotes/all
...
$ curl localhost:8080/quotes/one
...
```
