FROM alpine
RUN mkdir -p app
COPY config.json /app/config.json
ADD system_setup /app/system_setup

EXPOSE 8090

ENTRYPOINT ["/app/system_setup"]