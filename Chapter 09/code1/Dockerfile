FROM golang:1.19
WORKDIR /app
ADD . .
RUN go get
RUN go build -o app .
CMD ["/app/app"]
