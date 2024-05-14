FROM registry.access.redhat.com/ubi9/go-toolset:1.19.13-4.1697647145

COPY . .
RUN go mod download

RUN go build -buildvcs=false -o ./main

ENV PORT 8081
EXPOSE 8081

CMD [ "./main -concurrent 100 -interval 10 https://api.sandbox-m2.ll9k.p1.openshiftapps.com:6443" ]
