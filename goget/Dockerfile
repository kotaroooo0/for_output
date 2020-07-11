# FROM golang:latest

# [+] Building 139.3s (6/6) FINISHED
# RUN go get github.com/ChimeraCoder/anaconda \
#     github.com/PuerkitoBio/goquery \
#     github.com/alicebob/miniredis \
#     github.com/bamzi/jobrunner \
#     github.com/gin-gonic/gin \
#     github.com/go-redis/redis

# [+] Building 161.8s (11/11) FINISHED
# RUN go get github.com/ChimeraCoder/anaconda
# RUN go get github.com/PuerkitoBio/goquery
# RUN go get github.com/alicebob/miniredis
# RUN go get github.com/bamzi/jobrunner
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/go-redis/redis

# [+] Building 87.2s (17/17) FINISHED
# FROM golang:latest as builder1
# RUN go get github.com/ChimeraCoder/anaconda

# FROM golang:latest as builder2
# RUN go get github.com/PuerkitoBio/goquery

# FROM golang:latest as builder3
# RUN go get github.com/alicebob/miniredis

# FROM golang:latest as builder4
# RUN go get github.com/bamzi/jobrunner

# FROM golang:latest as builder5
# RUN go get github.com/gin-gonic/gin

# FROM golang:latest as builder6
# RUN go get github.com/go-redis/redis

# FROM scratch
# COPY --from=builder1 /go/bin /bin
# COPY --from=builder2 /go/bin /bin
# COPY --from=builder3 /go/bin /bin
# COPY --from=builder4 /go/bin /bin
# COPY --from=builder5 /go/bin /bin
# COPY --from=builder6 /go/bin /bin
