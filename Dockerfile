FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go build -o app
ARG PORT=8085
ENV PORT=${PORT}
EXPOSE ${PORT}
CMD ["./app"]