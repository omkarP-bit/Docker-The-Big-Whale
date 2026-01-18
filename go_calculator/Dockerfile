# Base image
FROM ubuntu AS build

RUN apt-get update && apt-get install -y golang-go

ENV G0111MODULE=off

COPY . .

RUN CGO__ENABLED=0 go build -o /app .

# stage2-main runtime image (go doesn't have runtime dependencies)
FROM scratch

COPY --from=build /app /app

ENTRYPOINT [ "/app" ]