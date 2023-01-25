FROM golang:1.18-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build cmd/main/main.go

FROM scratch
COPY --from=build /src/main /bin/main
ENTRYPOINT ["/bin/main"]