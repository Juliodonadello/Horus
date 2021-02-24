ARG GO_VERSION=1.15

FROM golang:${GO_VERSION} AS builder


WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

FROM scratch AS final
# Compilo mi imagen
COPY --from=builder /app /app
COPY ./docs ./docs
COPY ./certs ./
EXPOSE 443
EXPOSE 80
# Configuro que programa inicia el contenedor cuando arranca
ENTRYPOINT ["/app"]