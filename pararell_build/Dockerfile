FROM alpine as test1
RUN sleep 1

FROM alpine as test2
RUN sleep 2

FROM alpine as test3
RUN sleep 3

FROM alpine as test4
RUN sleep 4

FROM alpine
RUN sleep 5

COPY --from=test1 /bin /bin
COPY --from=test2 /bin /bin
COPY --from=test3 /bin /bin
COPY --from=test4 /bin /bin

RUN sleep 6
