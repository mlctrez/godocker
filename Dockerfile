FROM scratch

ADD bin/godocker /godocker
ADD bin/ca-bundle-amazonlinux.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

CMD ["/godocker"]