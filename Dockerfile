FROM scratch

ADD bin/godocker /godocker
ADD bin/ca-bundle.crt /etc/pki/tls/certs/ca-bundle.crt
ADD static /static

EXPOSE 8080

CMD ["/godocker"]