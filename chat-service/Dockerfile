FROM golang:latest
RUN mkdir /app
COPY main /app/service/
COPY certmagic /root/.local/share/
WORKDIR /app/service
CMD ["./main"]