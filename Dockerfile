FROM golang:1.25-alpine AS build

COPY ./ /build/

WORKDIR /build/

RUN apk add git
RUN go mod download

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
    go build -o ./fdb

FROM alpine

COPY --from=build /build/fdb /bin/fdb

ENTRYPOINT ["/bin/fdb"]

CMD ["-c", "/var/lib/fdb"]