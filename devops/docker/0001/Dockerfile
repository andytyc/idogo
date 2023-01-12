FROM ubuntu:20.04

RUN apt-get update && apt-get install -y tzdata \
    && rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Shanghai

####

ARG RUN_DIR
ENV RUN_DIR=${RUN_DIR:-/home/watrix/darwin/run}
WORKDIR $RUN_DIR

ARG APP_SERVER_NAME
ENV APP_SERVER_NAME=${APP_SERVER_NAME:-unknown}

# EXPOSE 12222

CMD ["pro"]
ENTRYPOINT ["/docker-entrypoint.sh"]
