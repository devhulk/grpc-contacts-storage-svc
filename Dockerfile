# FROM golang:onbuild
FROM golang:1.13.3

WORKDIR /go/src/storage-service
COPY . .
COPY ./cmd/storage-svc .

# RUN go get -d -v ./storage-svc
# RUN go build -o storage-service 
# RUN go install ./storage-svc

CMD ["./storage-svc"]