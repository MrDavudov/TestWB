FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o testwb ./cmd/main.go

CMD ["./testwb"]

# docker build -t testwb .