FROM registry.cn-hangzhou.aliyuncs.com/suosi/base-busybox:uclibc-tzdata

WORKDIR /app

COPY app ./deploy/docker/docker-entry.sh ./config/*-docker.yaml ./

CMD ["sh", "./docker-entry.sh"]
