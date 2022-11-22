FROM golang:latest AS builder

WORKDIR /go/src/github.com/elephant-insurance/ms-template
ADD . /go/src/github.com/elephant-insurance/ms-template/

RUN uname -a; go version; go env ; go build -mod=vendor -o ms-template .


FROM alpine:latest
WORKDIR /usr/local/ms-template
COPY --from=builder /go/src/github.com/elephant-insurance/ms-template/ms-template .
COPY --from=builder /go/src/github.com/elephant-insurance/ms-template/config.yml .
COPY --from=builder /go/src/github.com/elephant-insurance/ms-template/config-tokens.yml .

RUN apk add --no-cache bash && \
  apk add --no-cache curl && \
  apk add --no-cache libc6-compat

CMD /usr/local/ms-template/ms-template

EXPOSE 4000