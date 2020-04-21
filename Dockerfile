FROM golang:latest as builder
LABEL maintainer "Rafael Barros <rafaelbarros.df@gmail.com>"

WORKDIR /go/src/calc-soma
COPY /src/calc-soma .
RUN go build -ldflags="-s -w" -o /go/calc-soma

FROM scratch

COPY --from=builder /go/calc-soma /go/calc-soma

CMD ["/go/calc-soma"]