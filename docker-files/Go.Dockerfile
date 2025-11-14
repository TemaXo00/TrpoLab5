FROM golang:1.25.4-alpine

WORKDIR /app

ARG MODULE_PATH
ARG SERVICE_PATH

COPY ${MODULE_PATH}/go.mod ./
COPY ${MODULE_PATH}/go.sum ./
RUN go mod download

COPY ${SERVICE_PATH}/ ./

RUN go build -o main .

CMD ["/app/main"]
