FROM alpine
RUN mkdir -p app/{static,run}
COPY config.json /app/run/config.json
ADD worker /app/run/worker

EXPOSE 8090

ENTRYPOINT ["/app/run/worker"]