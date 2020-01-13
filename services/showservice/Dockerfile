FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o showservice/main showservice/main.go

FROM iron/go
COPY --from=builder /app/showservice/main /app/showservice
EXPOSE 3003
ENTRYPOINT [ "/app/showservice" ]
