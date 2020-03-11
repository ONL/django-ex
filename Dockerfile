FROM golang as builder

WORKDIR /go/src/github.com/onl/interactive-maps/
COPY . .
RUN go get -v github.com/gorilla/mux
RUN go build

FROM scratch
WORKDIR /bin
COPY --from=builder /go/src/app/bin/abc /bin/interactivemaps

ENV PASSWORD Test123

ENTRYPOINT ["/bin/interactivemaps"]
CMD ["--conf", "/etc/Caddyfile", "--agree=true"]
