FROM golang:1.16-alpine as build
RUN apk add --no-cache git musl-dev gcc

COPY .gitconfig /root/.gitconfig

WORKDIR /app
COPY . .

ARG VERSION
ENV VERSION $VERSION
ENV GOPRIVATE=github.com/synspective
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X github.com/synspective/syns-platform-backend-sample-rest/version=${VERSION}" -o ./bin/sample ./cmd/sample/main.go

# running image
FROM alpine

WORKDIR /go

ARG PROJECT_ID
ENV PROJECT_ID $PROJECT_ID

COPY --from=build /app/bin /go/bin

ENTRYPOINT ["./bin/sample"]
