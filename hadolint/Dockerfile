FROM golang as qiitaexporter-builder
RUN CGO_ENABLED=0 go get github.com/tenntenn/qiitaexporter

FROM golang as blogsync-builder
RUN CGO_ENABLED=0 go get github.com/x-motemen/blogsync

FROM alpine

WORKDIR /Documents
COPY blogsync.template /Documents
COPY setup.sh /Documents
RUN chmod +x setup.sh

RUN mkdir -p ~/.config/blogsync
COPY config.yaml.tmp /root/.config/blogsync

RUN apk --no-cache add libintl && \
    apk --no-cache add --virtual .gettext gettext && \
    cp /usr/bin/envsubst /usr/local/bin/envsubst && \
    apk del .gettext

COPY --from=qiitaexporter-builder /go/bin/qiitaexporter /bin/qiitaexporter
COPY --from=blogsync-builder /go/bin/blogsync /bin/blogsync

ENTRYPOINT ./setup.sh
