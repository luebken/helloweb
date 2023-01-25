# Go
```
go run cmd/main/main.go

go build cmd/main/main.go
./main

curl localhost:8080
```

# Docker

```
docker image build -t helloweb .
docker container run -p 8080:8080 helloweb

curl localhost:8080
```

# Fly.io

```
flyctl apps create helloweb-01
sed -i '' 's/APP_NAME/helloweb-01/g' fly.toml
fly deploy
fly open
```