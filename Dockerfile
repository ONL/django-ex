FROM golang as builder

WORKDIR /go/src/github.com/onl/interactivemaps/
COPY . .
RUN go get -v github.com/gorilla/mux
RUN go build

FROM scratch
COPY static /bin/static
COPY templates /bin/templates
COPY --from=builder /go/src/github.com/onl/interactivemaps/interactivemaps /bin/interactivemaps

ENV PASSWORD Test123

ENTRYPOINT ["/bin/interactivemaps"]
