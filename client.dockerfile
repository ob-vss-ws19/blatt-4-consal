FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o cli/client cli/client.go

FROM iron/go
COPY --from=builder /app/cli/client /app/cli
EXPOSE 52000-53000
ENTRYPOINT [ "/app/cli" ]
