FROM golang:1.20-alpine
WORKDIR /
COPY . .
RUN go build -o spoty-incubator
EXPOSE 8080
CMD ["./spoty-incubator"]
