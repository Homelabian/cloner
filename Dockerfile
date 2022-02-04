FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
RUN go build -o /cloner

FROM alpine
WORKDIR /
COPY --from=build /cloner /bin/cloner

ENTRYPOINT [ "/bin/cloner" ]