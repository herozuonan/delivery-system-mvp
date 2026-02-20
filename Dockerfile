FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o /out/delivery ./cmd/delivery

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/delivery /delivery
EXPOSE 8080
ENV ADDR=":8080"
ENTRYPOINT ["/delivery"]
