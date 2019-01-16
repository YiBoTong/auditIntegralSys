FROM alpine
RUN mkdir -p app/{static,run}
COPY config.json /app/run/config.json
ADD worker /app/run/worker

EXPOSE 8092

ENTRYPOINT ["/app/run/worker"]