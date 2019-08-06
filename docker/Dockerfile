# build first the image based on Dockerfile-openresty from https://github.com/opentracing-contrib/nginx-opentracing
FROM openresty-opentracing:latest
MAINTAINER Juan Matias Kungfu de la Camara Beovide <juanmatias@gmail.com>

RUN mkdir -p /var/log/nginx
RUN adduser --system --no-create-home --shell /bin/false --group --disabled-login nginx
