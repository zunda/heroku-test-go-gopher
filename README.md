# go-gopher
A minimal Go app to be run as a Heroku container registry app

## Running and accessing the app
### Local
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
```
$ heroku apps:create
$ git push heroku master
$ heroku apps:open
```

### Container
```
$ heroku update beta
$ heroku plugins:install @heroku-cli/plugin-manifest
$ heroku apps:create --manifest
$ heroku stack:set container
$ git push heroku master
$ heroku apps:open
```

To get Heroku CLI back to stable release:

```
$ heroku update stable
$ heroku plugins:remove manifest
```

## Shell access
### Local
```
$ docker run -it go-gopher bash
```

## Resources
- [Example `Dockerfile` for a go Heroku app](https://github.com/heroku/go-getting-started/blob/master/Dockerfile)
- Heroku Dev Center
  - [Container Registry & Runtime (Docker Deploys)](https://devcenter.heroku.com/articles/container-registry-and-runtime)
  - [Docker Builds with heroku.yml](https://devcenter.heroku.com/articles/docker-builds-heroku-yml)

## License
- [`favicon.ico`](https://golang.org/favicon.ico): Creative Commons Attribution 3.0 License
- Others: [MIT](LICENSE)
