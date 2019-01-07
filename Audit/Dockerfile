FROM alpine
RUN mkdir -p app/{static,run}
COPY config.json /app/run/config.json
ADD audit /app/run/audit

EXPOSE 8093

ENTRYPOINT ["/app/run/audit"]