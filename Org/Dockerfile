FROM alpine
RUN mkdir -p app/{static,run}
COPY config.json /app/run/config.json
ADD org /app/org

EXPOSE 8091

ENTRYPOINT ["/app/run/org"]