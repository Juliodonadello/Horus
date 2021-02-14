ARG GO_VERSION=1.15

FROM golang:${GO_VERSION}-alpine AS builder


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
EXPOSE 5000

# Configuro que programa inicia el contenedor cuando arranca
ENTRYPOINT ["/app"]