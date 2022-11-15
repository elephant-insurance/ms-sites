FROM golang:latest AS builder
COPY ../ms-dnc .

WORKDIR /go/src/github.com/elephant-insurance/ms-sites
ADD . /go/src/github.com/elephant-insurance/ms-sites/

RUN uname -a; go version; go env ; go build -mod=vendor -o ms-sites .

FROM alpine:latest
WORKDIR /usr/local/ms-sites

COPY --from=builder /go/src/github.com/elephant-insurance/ms-sites/ms-sites .
COPY --from=builder /go/src/github.com/elephant-insurance/ms-sites/config.yml .
COPY --from=builder /go/src/github.com/elephant-insurance/ms-sites/config-tokens.yml .


RUN apk add --no-cache bash && \
  apk add --no-cache curl && \
  apk add --no-cache libc6-compat

CMD /usr/local/ms-sites/ms-sites

EXPOSE 4000
