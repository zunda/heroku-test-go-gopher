# go-gopher
A minimal Go app to be run as a Heroku container registry app

## Run
### Locally
The server can be accessed at http://localhost:3000/

```
$ go run main.go
```

or

```
$ docker build -t go-gopher .
$ docker run -p 3000:3000 go-gopher
```

## As a Heroku app
### Slug

## Resources
https://github.com/heroku/go-getting-started/blob/master/Dockerfile and edited.
