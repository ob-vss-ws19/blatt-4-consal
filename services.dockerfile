FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/services Services/main.go

FROM iron/go
COPY --from=builder /app/services/services /app/services
EXPOSE 3000
EXPOSE 3001
EXPOSE 3002
EXPOSE 3003
EXPOSE 3004
ENTRYPOINT [ "/app/services" ]
