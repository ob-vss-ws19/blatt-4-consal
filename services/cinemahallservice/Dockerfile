FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/cinemahallservice/main services/cinemahallservice/main.go

FROM iron/go
COPY --from=builder /app/services/cinemahallservice/main /app/services/cinemahallservice
EXPOSE 3002
ENTRYPOINT [ "/app/cinemahallservice" ]
