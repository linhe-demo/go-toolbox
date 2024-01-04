FROM alpine:3.12.0

MAINTAINER linhe

RUN mkdir -p "/data/app" && ln -sf /dev/stdout /data/app/service.log

WORKDIR "/data/app"

COPY toolbox /data/app/toolbox

CMD ["./toolbox", "-f", "etc/toolbox-api.yaml"]