FROM golang:1.18 as build
COPY . /src
WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o kvs github.com/go-programs/key-value-store/cmd/server

FROM scratch
COPY --from=build /src/kvs .
EXPOSE 8080
CMD ["/kvs"]
