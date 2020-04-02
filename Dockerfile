FROM reg.miz.so/library/centos:7
MAINTAINER "sdyx-basic &lt;jaden@hyx.com>"

WORKDIR /app

RUN rm /etc/localtime && \
ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ADD ./main /app/main

RUN mkdir -p  /app/conf
RUN mkdir logs
RUN mkdir /data && \
cd /data && \
mkdir logs

CMD ["./main", "-conf", "conf/config.json"]
