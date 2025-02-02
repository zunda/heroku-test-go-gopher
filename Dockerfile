# https://github.com/heroku/go-getting-started/blob/main/Dockerfile
FROM heroku/heroku:22-build as build

COPY . /app
WORKDIR /app

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl -s https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xvz -C /tmp/buildpack/heroku/go

#Execute Buildpack
RUN STACK=heroku-22 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:20

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
USER heroku
CMD /app/bin/go-gopher
