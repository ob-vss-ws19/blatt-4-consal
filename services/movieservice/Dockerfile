FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o movieservice/main movieservice/main.go

FROM iron/go
COPY --from=builder /app/movieservice/main /app/movieservice
EXPOSE 3001
ENTRYPOINT [ "/app/movieservice" ]
