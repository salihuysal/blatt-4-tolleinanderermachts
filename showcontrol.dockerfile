FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o showcontrol/main showcontrol/main.go showcontrol/showcontrol.go

FROM iron/go
COPY --from=builder /app/showcontrol/main /app/showcontrol
EXPOSE 50000-70000
ENTRYPOINT ["/app/showcontrol"]
