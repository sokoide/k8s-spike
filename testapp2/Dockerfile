FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir -p /app
RUN mkdir -p /logs;chmod 777 /logs
COPY ./testapp2 /app
WORKDIR /app
CMD ["/app/testapp2"]

