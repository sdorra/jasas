FROM scratch
MAINTAINER Sebastian Sdorra <s.sdorra@gmail.com>

COPY dist/jasas_linux_amd64 /jasas

EXPOSE 8000

ENTRYPOINT ["/jasas"]
CMD ["daemon"]
