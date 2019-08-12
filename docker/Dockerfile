FROM golang:1.12.6 AS builder

WORKDIR $GOPATH/src/github.com/rugwirobaker/larissa

ENV GO111MODULE=on

COPY . .

ARG VERSION="unset"

RUN DATE="$(date -u +%Y-%m-%d-%H:%M:%S-%Z)" && GO111MODULE=on CGO_ENABLED=0 GOPROXY="https://proxy.golang.org" go build -ldflags "-X github.com/gomods/athens/pkg/build.version=$VERSION -X github.com/gomods/athens/pkg/build.buildDate=$DATE" -o /bin/larissa ./cmd/.

FROM alpine

COPY --from=builder /bin/larissa /bin/larissa

RUN apk add --update tini

EXPOSE 3000

ENTRYPOINT [ "/sbin/tini", "--" ]

CMD ["larissa"]
