FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/showservice/main services/showservice/main.go

FROM iron/go
COPY --from=builder /app/services/showservice/main /app/services/showservice
EXPOSE 52000-53000
ENTRYPOINT [ "/app/services/showservice" ]
