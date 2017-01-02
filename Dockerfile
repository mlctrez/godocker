FROM scratch

ADD bin/godocker godocker

EXPOSE 8080

CMD ["/godocker"]