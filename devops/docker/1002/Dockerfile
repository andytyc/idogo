FROM golang:1.16-alpine

# cd /app
WORKDIR /app

# copy code to /app
COPY . .

# build
RUN go build -o /idogo

# run command
CMD [ "/idogo -n=2" ]