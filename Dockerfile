FROM golang:1.17.3-alpine AS builder
RUN apk update && \
    apk add --no-cache build-base imagemagick imagemagick-dev && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /go/lgtm-generator
COPY . .
RUN go install

FROM alpine
RUN apk update && \
    apk add --no-cache imagemagick imagemagick-dev && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /go/bin/lgtm-generator /app/lgtm-generator
CMD /app/lgtm-generator
