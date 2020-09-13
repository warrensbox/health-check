FROM alpine:latest

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin ${GOPATH}/pkg ${GOPATH}/go/src/health-check

ADD . ${GOPATH}/go/src/health-check

WORKDIR ${GOPATH}/go/src/health-check

RUN ls
RUN go version

RUN go mod download
RUN go get -d -t -v ./...

RUN make
ENTRYPOINT ["./health-check"]