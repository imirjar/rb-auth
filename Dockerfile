FROM golang:1.24
WORKDIR /app
COPY . /app/
# COPY db/reports /db
# COPY bin/diver .
# COPY config/diver .
RUN go build -o rb-auth ./cmd/main.go
CMD ["./rb-auth"]