FROM golang:1.11 as build
WORKDIR /go/src/github.com/jukeizu/chat
COPY Makefile go.mod go.sum ./
RUN make deps
ADD . .
RUN make build-linux
RUN echo "jukeizu:x:100:101:/" > passwd

FROM scratch
COPY --from=build /go/src/github.com/jukeizu/chat/passwd /etc/passwd
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build --chown=100:101 /go/src/github.com/jukeizu/chat/bin/chat .
USER jukeizu
ENTRYPOINT ["./chat"]
