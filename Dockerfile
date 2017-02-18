FROM scratch

ADD bin/godocker /godocker
ADD bin/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

CMD ["/godocker"]