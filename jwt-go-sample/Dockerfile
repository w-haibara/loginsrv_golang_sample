FROM golang:alpine AS jwt-go-builder

USER root
WORKDIR /jwt-go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add make

COPY src ./
RUN make test && make

FROM scratch

WORKDIR /jwt-go

COPY --from=jwt-go-builder /jwt-go/jwt-go-sample .
COPY secret.key .
CMD ["./jwt-go-sample"]
