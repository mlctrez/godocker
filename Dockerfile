FROM scratch

#ADD bin/godocker godocker
ADD bin/* /

EXPOSE 8080

CMD ["/godocker"]