# Go Kafka Example

### How To
* install librdkafka => https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
* go get confluent-go-kafka  => https://github.com/confluentinc/confluent-kafka-go
* or just run glide install

## Install Kafka
* download source kafka in : https://kafka.apache.org/downloads
* extract downloaded file
* change directory to extracted folder


## Start kafka services
##### start zookeper
```
$ ./bin/zookeeper-server-start.sh config/zookeeper.properties
```
##### start server
```
$ ./bin/kafka-server-start.sh config/server.properties
```
##### create topic
```
$ ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
```

## Producer
##### start producer
```
$ ./bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
```

## Consumer
##### start consumer
```
$ ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test
```

## Run Go Kafka example


#### Producer
##### simple producer with one message inline
```
$ go run simpleProducer.go
```
##### simple producer with one message with cli argument
```
$ go run sendMessage.go <topic> <message>

// example

$ go run sendMessage.go test hallo-test
```

##### simple producer with bulk message
```
$ go run kafkaProducer.go
```
*) press ctrl+c to SIGINT



#### Consumer

##### simple consumer with one message and from earliest message
```
$ go run simpleConsumer.go

$ go run kafkaConsumer.go
```
### simple consumer with offset message
```
$ go run offsetConsumer.go
```

## tutorial
https://hardyantz.github.io/go-kafka/

