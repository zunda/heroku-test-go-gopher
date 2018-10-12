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
Creating app... done, stack is container
$ heroku stack
=== Available Stacks
  cedar-14
* container
  heroku-16
  heroku-18
$ git push heroku master
Counting objects: 38, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (17/17), done.
Writing objects: 100% (38/38), 4.62 KiB | 2.31 MiB/s, done.
Total 38 (delta 16), reused 38 (delta 16)
remote: Compressing source files... done.
remote: Building source:
remote: === Fetching app code
remote: =!= Build failed due to an error:
remote:
remote: =!= generate step: can't load manifest before running validations
remote:
remote: If this persists, please contact us at https://help.heroku.com/.
remote: Verifying deploy...
remote:
remote: !	Push rejected to obscure-shelf-55765.
remote:
To https://git.heroku.com/obscure-shelf-55765.git
 ! [remote rejected] master -> master (pre-receive hook declined)
```

(WAT)

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
