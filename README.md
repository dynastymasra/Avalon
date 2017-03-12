# Avalon

[![Go](https://img.shields.io/badge/go-1.8.0-00E5E6.svg)](https://golang.org/)
[![Glide](https://img.shields.io/badge/glide-0.12.3-CFBDB1.svg)](https://glide.sh/)
[![Elasticsearch](https://img.shields.io/badge/elasticsearch-5.2.2-FED10A.svg)](https://www.elastic.co/products/elasticsearch)
[![Postgres](https://img.shields.io/badge/postgres-9.6.2-2C5687.svg)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/license-MIT-44897A.svg)](https://github.com/dynastymasra/Avalon/blob/master/LICENSE)

Avalon is simple project CRUD Golang (Go Programming Language) with database Postgres and Elasticsearch

- Documentation this service in [Wiki](https://github.com/dynastymasra/Avalon/wiki)

## Libraries
Use glide command for install all dependencies required this application.
  - Use command `glide install` for install all dependency.
  - Use command `glide up` for update all dependency.

  ## How To Run and Deploy

### *Local*

Use commad go `go run main.go` in root folder for run this application.

-----

### *Docker*

#### Build Docker

Avalon uses a Docker pattern called [Build Container](https://medium.com/@alexeiled/docker-pattern-the-build-container-b0d0e86ad601#.ixjx51u42) in order to make sure the deployed artifact 
can be run instantly and ephemerally. The way it does this is by using separate container for two separate deployment steps:

  - `docker build -f docker/Dockerfile.build -t $(IMAGE) .` This command will build the images.
  - `docker create $(IMAGE) /bin/bash -xe ./start.sh build` Run image base on image above.
  - `docker start -a $(CONTAINER)` Start container.
  - `docker cp $(CONTAINER):/binary .` Copy all file **build go** inside folder ***binary*** to ***host***.
  - `docker build -f docker/Dockerfile -t $(IMAGE) .` This will build image with binary file.

#### Run Docker
```
 docker run --name avalon -d -p <port>:<port> -e ADDRESS=:8080 -e MODE=debug -e <environment> coralteam/avalon:<version>
```  

## Makefile

Makefile used for development propose only, for make easy for build image develop testing. have var `REPOSITORY` default `coralteam/avalon` and 
`VERSION` default `develop`, use command `make build` for create docker image.

## Environment Variable

+ `ADDRESS` - Address application is used default is `:8080`
+ `GIN_MODE` - Mode application is used `debug` or `release`, default is `release`
  - `debug` - Application will run with debug mode, show application full log.
  - `release` - Application will run with release mode, show less log.
+ `POSTGRES_USERNAME` - Postgres username, default is `postgres`
+ `POSTGRES_PASSWORD` - Postgres password, default is `root`
+ `POSTGRES_ADDRESS` - Postgres address, default is `192.168.99.100:5432`
+ `POSTGRES_DATABASE` - Default postgres database name is used, default is `avalon`
+ `POSTGRES_LOGGING` - Show log postgres database, default is `false`