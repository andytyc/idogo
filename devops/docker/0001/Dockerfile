# FROM andytyc/ubuntu:16.04
FROM ubuntu:20.04

RUN apt-get update && apt-get install -y tzdata \
    && rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Shanghai

########
# build
########

# cd /app
WORKDIR /app

# copy code to /app
COPY . .

# build
RUN go build -o /idogo

# run
CMD [ "/idogo" ]