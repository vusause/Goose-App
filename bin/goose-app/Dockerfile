FROM golang:latest as builder
COPY main.go /go/src/github.com/vusause/goose-app/
WORKDIR /go/src/github.com/vusause/goose-app/
RUN go get -d -v github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /go/src/github.com/vusause/goose-app/main /app/
WORKDIR /app
CMD ["./main"]
