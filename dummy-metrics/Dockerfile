FROM golang:1.19 as build

WORKDIR /go/src/dummy-metrics
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/dummy-metrics

FROM gcr.io/distroless/static-debian11

LABEL org.opencontainers.image.source https://github.com/zoetrope/dummy-metrics

COPY --from=build /go/bin/dummy-metrics /
CMD ["/dummy-metrics"]
