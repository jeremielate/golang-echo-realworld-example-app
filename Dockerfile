FROM docker.io/library/golang:1-alpine AS build
WORKDIR /src
COPY . .
RUN go build -buildvcs=false

FROM docker.io/library/alpine:3.16.1
RUN apk add ca-certificates
COPY --from=build /src/golang-echo-realworld-example-app /usr/bin/realworld

ENTRYPOINT ["/usr/bin/realworld"]
