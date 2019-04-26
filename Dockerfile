FROM golang:1.10

ADD . /go/src/golang-rest-api-best-practice
WORKDIR /go/src/golang-rest-api-best-practice
RUN go get -d github.com/gorilla/mux
RUN go get -d go.mongodb.org/mongo-driver/bson

CMD ["go","run","main.go"]
