FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main ./cmd/*.go

FROM alpine

WORKDIR /workdir

COPY --from=builder /build/configs /workdir/configs 
COPY --from=builder /build/main /workdir/main


ENTRYPOINT ["./main"]