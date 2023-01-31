# helloweb
A small scaffold for building a web application with Go and Sqlite.


## Database
```sh
sqlite3 db/sqlite3.db < db/database.sql
```

## Go
```sh
export DB_PATH=db

go run cmd/main/main.go
# or
go build cmd/main/main.go
./main

curl localhost:8080
```

## Docker local
```sh
export DOCKER_DEFAULT_PLATFORM=linux/amd64
docker image build -t helloweb .
export DB_PATH=db
docker container run -p 8080:8080 -e DB_PATH=/mnt -v $PWD/db:/mnt helloweb

curl localhost:8080
```

## Fly.io
Getting started:
```sh
flyctl auth login

fly launch --generate-name --copy-config --region fra
fly volumes create myapp_data --region fra --size 1
fly deploy
fly ssh sftp shell
$ put db/sqlite3.db /data/sqlite3.db

fly logs
fly apps restart $(fly info --json | jq -r .Name)

fly apps open
```

Debug/Troubleshoot/Other:
```sh
fly doctor

flyctl ssh console
$ ls -l /data

flyctl auth docker
docker tag helloweb:latest registry.fly.io/helloweb:v1
docker push registry.fly.io/helloweb:v1

cat ~/.fly/config.yml

flyctl destroy
```