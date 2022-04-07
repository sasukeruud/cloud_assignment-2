FROM golang:1.17-alpine as Builder

LABEL stage=builder

# Set up execution environment in container's GOPATH
WORKDIR /go/src/app/cmd

# Copy relevant folders into container
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum
COPY ./cmd /go/src/app/cmd
COPY ./handlers /go/src/app/handlers
COPY ./readJson /go/src/app/readJson

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

# To get the time zone data
FROM alpine:latest as alpine-with-tz
RUN apk --no-cache add tzdata zip
WORKDIR /usr/share/zoneinfo

#Compressing the zone data
RUN zip -q -r -0 /zoneinfo.zip .

# Final container
FROM scratch AS final

LABEL maintainer="robinru@stud.ntnu.no"

# Root as working directory to copy compiled file to
WORKDIR /

# Retrieve binary from builder container
COPY --from=builder /src/app/cmd/main .

# Setting time zone data
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine-with-tz /zoneinfo.zip /
ENV TZ=Europe/Berlin

# Fetching the cert hints.
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 80

# Instantiate server
CMD ["./main"]