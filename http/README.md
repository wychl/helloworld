# hello-world

## build

```sh
docker build -t  wanyanchengli/hello-world:latest .
```

## push

```sh
docker push wanyanchengli/hello-world:latest
```

## run

```sh
docker run --rm  -p 8080:8080 wanyanchengli/hello-world:latest
```