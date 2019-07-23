# Elasticsearch Kibana Filebeat GoLang service

### Using technologies

- Golang 1.12
- Elasticsearch v7.2.0
- Kibana v7.2.0
- Filebeat v7.2.0
- Docker and docker-compose for running


### How to run

1. Clone project:
```
go get github.com/ypronin/elk_golang_compose
```

2. Open project folder:
```
cd ../github.com/ypronin/elk_golang_compose
```

3. Run services:
```
docker-compose up -d
```

### Answers
1. As "any requests for new additional data might take a looooong time to implement". Better to have a shipper(filebeat) for log files. In the request/response logs it's important to see fields e.g(txId, playerId, ipAddress, amount, gameName, currency, balance etc.) required to analyze data.
2. There is not need to spend a lot of time of poker team, mostly developers are logging all required information for debugging purpose
3. I have tried two options, to use hook and send log imiddialty to ES or use filebeat to parse log and send them to ES
4. In elasticsearch we are able to create a different alerts, monitoring, froud detections.