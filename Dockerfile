FROM golang:1.17.3-alpine

RUN apk update && \
    apk add --no-cache build-base imagemagick imagemagick-dev && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/lgtm-generator
COPY . .
RUN go install
CMD /go/bin/lgtm-generator
