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
```sh
fly launch
# use the same region!
fly volumes create myapp_data --region fra --size 1
```

```toml
# add mounts and env to toml:
[mounts]
source="myapp_data"
destination="/data"

[env]
DB_PATH="/data"
```

```sh
fly deploy

fly logs | grep DB_PATH

fly ssh sftp shell
$ put db/sqlite3.db /data/sqlite3.db

flyctl ssh console
$ ls -l /data

flyctl apps open
```