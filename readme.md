# Database
```
sqlite3 db/sqlite3.db < db/database.sql
```

# Go
```
export DB_PATH=db
go run cmd/main/main.go

go build cmd/main/main.go
./main

curl localhost:8080
```

# Docker local
```
export DOCKER_DEFAULT_PLATFORM=linux/amd64
docker image build -t helloweb .
export DB_PATH=db
docker container run -p 8080:8080 -e DB_PATH=/mnt -v $PWD/db:/mnt helloweb

curl localhost:8080

Fly.io
```
fly launch --auto-confirm
fly deploy
fly volumes create myapp_data --region fra --size 1
```