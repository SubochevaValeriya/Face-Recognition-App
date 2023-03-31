FROM golang:1.18-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executablew
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o face-recognition-app ./cmd/main.go

CMD ["./face-recognition-app"]