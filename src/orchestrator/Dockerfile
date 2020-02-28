FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY orchestrator /app/orchestrator
COPY pod.json /app/pod.json

CMD ["./orchestrator"]