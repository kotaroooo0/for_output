FROM alpine:latest as stage1
LABEL stage=1
WORKDIR /workdir
RUN sleep 3 && touch stage1.txt


FROM alpine:latest as stage2
LABEL stage=
WORKDIR /workdir
RUN sleep 6 && touch stage2.txt

COPY --from=stage1 /workdir/stage1.txt /workdir/stage1.txt

FROM alpine:latest as stage3
LABEL stage=3
WORKDIR /workdir
COPY --from=stage1 /workdir/stage1.txt /workdir/stage1.txt
COPY --from=stage2 /workdir/stage2.txt /workdir/stage2.txt



