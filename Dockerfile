FROM golang:1.11.2-alpine AS build

RUN apk add git

WORKDIR /go/src/github.com/kastenhq/zipcode

ENV GO111MODULE=on

COPY cmd cmd
COPY pkg pkg
COPY go.mod .
COPY go.sum .

RUN go mod download && \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go install -installsuffix "static" -ldflags "-X ${PKG}/pkg/version.VERSION=${VERSION}" \
    ./cmd/zipcode

FROM alpine AS zipcode

COPY --from=build /go/bin/zipcode /bin/zipcode

ENTRYPOINT ["/bin/zipcode"]
