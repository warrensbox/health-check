# FROM alpine:3.12

# ARG GOLANG_VERSION=1.14.3

# #we need the go version installed from apk to bootstrap the custom version built from source
# RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates && update-ca-certificates

# RUN wget https://dl.google.com/go/go$GOLANG_VERSION.src.tar.gz && tar -C /usr/local -xzf go$GOLANG_VERSION.src.tar.gz

# RUN cd /usr/local/go/src && ./make.bash

# ENV PATH=$PATH:/usr/local/go/bin

# RUN rm go$GOLANG_VERSION.src.tar.gz

# #we delete the apk installed version to avoid conflict
# RUN apk del go

# RUN go version

FROM alpine:latest

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin ${GOPATH}/pkg ${GOPATH}/go/src/health-check

# Install Glide
#RUN go get -u github.com/Masterminds/glide/...

ADD . ${GOPATH}/go/src/health-check

WORKDIR ${GOPATH}/go/src/health-check

RUN ls
RUN go version

RUN go mod download
RUN go get -d -t -v ./...
#RUN go install -v ./...

RUN make

ENTRYPOINT ["./health-check"]