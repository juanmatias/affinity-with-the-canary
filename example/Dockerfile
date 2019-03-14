FROM alpine:latest
MAINTAINER Juan Matias Kungfu de la Camara Beovide "juan.delacamara@3xmgroup.com"

# In case no version is passed
ARG version=1.0
COPY source/$version/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
