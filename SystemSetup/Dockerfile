FROM alpine
RUN mkdir -p app/{static,run}
COPY config.json /app/run/config.json
ADD system_setup /app/run/system_setup

EXPOSE 8090

ENTRYPOINT ["/app/run/system_setup"]