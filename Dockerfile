FROM scratch
MAINTAINER Sebastian Sdorra <s.sdorra@gmail.com>

COPY dist/jasas /jasas
COPY index.html /index.html
COPY dist/bundle.js /dist/bundle.js
COPY dist/bundle.js.map /dist/bundle.js.map

EXPOSE 8000

ENTRYPOINT ["/jasas"]
CMD ["daemon"]
