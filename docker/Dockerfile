# Stick goreleaser output in Distroless container
FROM gcr.io/distroless/static-debian12 AS final-stage
WORKDIR /
COPY dnh ./dnh

EXPOSE 9732

CMD ["/dnh", "server", "--port", "9732"]
