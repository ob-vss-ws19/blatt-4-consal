FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o cinemahallservice/main cinemahallservice/main.go

FROM iron/go
COPY --from=builder /app/cinemahallservice/main /app/cinemahallservice
EXPOSE 3002
ENTRYPOINT [ "/app/cinemahallservice" ]
