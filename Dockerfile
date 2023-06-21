FROM centos:7
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir /usr/local/cp-audio
RUN mkdir /usr/local/cp-audio/bin
RUN mkdir /usr/local/cp-audio/etc
COPY bin /usr/local/cp-audio/bin
COPY container_start.sh /usr/local/cp-audio/container_start.sh

WORKDIR /usr/local/cp-audio
