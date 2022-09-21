FROM golang:1.19-buster AS builder

ARG APP_RELATIVE_PATH

COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN GOPROXY=https://goproxy.cn make build

FROM ubuntu:20.04

ARG APP_RELATIVE_PATH

# apt 更换源
RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

RUN apt-get install apt-transport-https
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
#COPY /src/app/${APP_RELATIVE_PATH}/bin /app
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/configs /data/conf

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["./server", "-conf", "/data/conf"]
