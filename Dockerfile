FROM golang:latest as builder

LABEL maintainer = "Matheus Carmo (a.k.a Carmel) <mematheuslc@gmail.com>"

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o guiomar_exec .

# Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /guiomar
COPY --from=builder /app/guiomar_exec .

ENTRYPOINT /guiomar/guiomar_exec

EXPOSE 3000

