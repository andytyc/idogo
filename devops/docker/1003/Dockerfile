# FROM andytyc/ubuntu:16.04
# FROM ubuntu:20.04
FROM golang:1.16-alpine

# cd /app
WORKDIR /app

# copy code to /app
COPY . .

# build
RUN go build -o /idogo

# run
CMD [ "/idogo" ]