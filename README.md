# go-gopher
A minimal Go app to be run as a Heroku container registry app

## Running and accessing the app
### Locally
```
$ go run main.go &
$ curl http://localhost:3000
```

(`fg` and `^C` to stop)

or

```
$ docker build -t go-gopher .
$ docker run -p 3000:3000 go-gopher &
$ curl http://localhost:3000
```

(`docker ps` and `docker stop <container-id>` to stop)

## As a Heroku app
### Slug

## Shell access
### Locally
```
$ docker run -it go-gopher bash
```

## Resources
- Example `Dockerfile` for a go Heroku app: https://github.com/heroku/go-getting-started/blob/master/Dockerfile
