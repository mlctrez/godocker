FROM scratch

ADD bin/godocker /godocker
ADD bin/ca-bundle.crt /etc/pki/tls/certs/ca-bundle.crt

EXPOSE 8080

CMD ["/godocker"]