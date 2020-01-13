# FROM golang:onbuild
FROM golang:1.13.3

WORKDIR /go/src/storage-service
COPY ./cmd/storage-svc .
COPY ./internal .

# RUN go get -d -v ./storage-svc
# RUN go build -o storage-service 
# RUN go install ./storage-svc

CMD ["./storage-svc"]