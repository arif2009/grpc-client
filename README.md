Sample codes for developing a REST back-end service.

# Required tools to run
- Go >= 1.16
- [Docker](https://www.docker.com/products/docker-desktop)

# How to run server
- via docker compose (requires gcloud.json)
    ```
    docker compose up -d
    ```
- directly (requires gcloud to be logged in)
    ```
    go run cmd/sample/main.go
    ```
