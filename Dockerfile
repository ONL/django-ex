FROM golang as builder

WORKDIR /go/src/github.com/onl/interactivemaps/
COPY . .
RUN go get -v github.com/gorilla/mux
RUN go build

FROM fedora:latest
RUN dnf upgrade -y && \
    dnf install -y \
       dumb-init \
    && dnf clean all

COPY static /opt/bin/static
COPY templates /opt/bin/templates
COPY --from=builder /go/src/github.com/onl/interactivemaps/interactivemaps /opt/bin/interactivemaps

ENV PASSWORD Test123
EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/opt/bin/interactivemaps"]
