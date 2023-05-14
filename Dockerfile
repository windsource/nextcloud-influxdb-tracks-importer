FROM golang:1.14-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base linux-headers
WORKDIR /go/src/
COPY . .
RUN go get -d -v ./...  
RUN go test ./... && go build -a -o app .

FROM alpine:latest  
LABEL org.opencontainers.image.source = "https://github.com/windsource/nextcloud-influxdb-tracks-importer" 
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /go/src/app .

CMD ["./app"] 
