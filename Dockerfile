FROM golang:1.18.9-alpine

LABEL maintainer="Johan Book"
LABEL title="file-receiver-server"
LABEL description="File receiver server"

WORKDIR /app
COPY . .
RUN ./scripts/build

CMD ["./main"]

EXPOSE 8080
