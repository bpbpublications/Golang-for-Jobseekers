FROM golang:1.19 as source
WORKDIR /app
ADD go.mod .
RUN go mod download
ADD . .
RUN go build -o app .

FROM debian
COPY --from=source /app/app /app
CMD ["/app"]
