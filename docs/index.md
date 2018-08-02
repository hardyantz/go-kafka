
this post is for beginner who want to learn basic kafka with golang. it is far from advance. if you think you are more advance than this. please share some thought. thanks.

# Basic Kafka

### What is kafka?
Apache Kafka was originated at LinkedIn and later became an open sourced Apache project in 2011, then First-class Apache project in 2012. Kafka is written in Scala and Java
Kafka is distributed streaming platform, means that kafka has capabilities to :
* Publish and subscribe to streams of records, similar to a message queue or enterprise messaging system.
* Store streams of records in a fault-tolerant durable way.
* Process streams of records as they occur.

Apache Kafka is publish-subscribe based fault tolerant messaging system. It is fast, scalable and distributed by design.
Kafka tends to work very well as a replacement for a more traditional message broker
In comparison to other messaging systems, Kafka has better throughput, built-in partitioning, replication and inherent fault-tolerance, which makes it a good fit for large-scale message processing applications.

### what is messaging system?
A Messaging System is responsible for transferring data from one application to another, so the applications can focus on data, but not worry about how to share it

main terminologies of kafka such as topics, brokers, producers and consumers. The following diagram illustrates the main terminologies and the table describes the diagram components in detail.

![](https://www.tutorialspoint.com/apache_kafka/images/fundamentals.jpg)

### Topics
A stream of messages belonging to a particular category is called a topic. Data is stored in topics.

Topics are split into partitions. For each topic, Kafka keeps a mini-mum of one partition. Each such partition contains messages in an immutable ordered sequence. A partition is implemented as a set of segment files of equal sizes.

### Brokers
* Brokers are simple system responsible for maintaining the pub-lished data. Each broker may have zero or more partitions per topic. Assume, if there are N partitions in a topic and N number of brokers, each broker will have one partition.

* Assume if there are N partitions in a topic and more than N brokers (n + m), the first N broker will have one partition and the next M broker will not have any partition for that particular topic.

* Assume if there are N partitions in a topic and less than N brokers (n-m), each broker will have one or more partition sharing among them. This scenario is not recommended due to unequal load distri-bution among the broker.

### Producers
Producers are the publisher of messages to one or more Kafka topics. Producers send data to Kafka brokers. Every time a producer pub-lishes a message to a broker, the broker simply appends the message to the last segment file. Actually, the message will be appended to a partition. Producer can also send messages to a partition of their choice.

### Consumers
Consumers read data from brokers. Consumers subscribes to one or more topics and consume published messages by pulling data from the brokers

![](https://www.tutorialspoint.com/apache_kafka/images/cluster_architecture.jpg)


# Install Kafka

this page is a step by step to install kafka for messaging system

### verify java
make sure you already install java.

### download kafka
* download kafka source (https://www.apache.org/dyn/closer.cgi?path=/kafka/2.0.0/kafka_2.11-2.0.0.tgz)
* extract source & change directory to kafka directory
```
$ tar -xzf kafka_2.11-2.0.0.tgz
$ cd kafka_2.11-2.0.0
```

### start zookepeer server
open new terminal and run / type
```
$ ./bin/zookeeper-server-start.sh config/zookeeper.properties
```

### start kafka server
open different new terminal, go to kafka folder and run / type
```
$ ./bin/kafka-server-start.sh config/server.properties
```

## Test the producer & consumer

### create topic
open different new terminal, go to kafka folder and run / type
```
$ ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
```
### run consumer
open different new terminal, go to kafka folder and run / type
```
$ ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test
```
### run producers
open different new terminal, go to kafka folder and run / type
```
$ ./bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
```
type something in the new line in the producer screen
```
> test one
> test two
```

open the consumers terminal and the results will be
```
test one
test two
```

try to type something in your producer terminal, it will shown in your consumer terminal

### run consumer message from beginning
lets try to retrieve all message from beginning in the topic `test`
open different new terminal, go to kafka folder and runt / type
the syntax is the same in the consumer but added `--from-beginning`
```
$ ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
```

### run consumer with specific offset message
you can get message in specified offset just add `--offset <offset>`
in case we want to retrieve message from offset 1 (in kafka, offset start from 0)
```
$ ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --offset 1 --partition 0
```

# Kafka With Golang

in this repository you can learn how to use basic usage of kafka with golang.
assume we use default configuration of kafka. 

how to :
* install librdkafka => https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
* clone this repository => https://github.com/hardyantz/go-kafka
* run glide install

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

# Kafka Next Step

## Retention log setup

Log retention setup is The minimum age of a log file to be eligible for deletion due to age,
means that kafka will save message for period of hours. default is 168 hours or 7 days,
to change this setup edit `config/server.properties` and change the value of
```
log.retention.hours=168
```

## Partition
in kafka, you can set numbers of partition you want to setup.

How to :

#### edit config file
edit config file in `config/server.properties` and change the value of
```
num.partitions=3
```
you can set the default number of log partitions per topic. More partitions allow greater
parallelism for consumption, but this will also result in more files across the brokers.


#### Setup topic partition
previously we already setup topic `test`. now lets setup topic `test` into 2 partitions.
change directory to kafka folder, and run
```
$ ./bin/kafka-topics.sh --zookeeper localhost:2181 --alter --topic test --partitions 2
```
lets test the status of the partition by typing
```
$ ./bin/kafka-topics.sh --describe --zookeeper localhost:2181 --topic test
```

you will see the result that topic test is set partition to 2 partition
```
Topic:test	    PartitionCount:2	ReplicationFactor:1	Configs:
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	Topic: test	Partition: 1	Leader: 0	Replicas: 0	Isr: 0

```

if you want to test how the message receive in broker.
test with run the consumer with `partition 0` & `partition 1` .
if partition running well. you will see that the consumer will receieve message alternately in every partition.

example :
consumer partition 0
```
$ ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --partition 0
one
three
```

consumer partition 1
```
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --partition 1
two
four
```

producer
```
$ ./bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
>one
>two
>three
>four
```

#### setup replication
to setup replication, you have to create new config for every replication.

##### create config file

###### replication 1
example :
create or copy from the existing config file. let say we create `server-1.properties` in config folder.
and set `broker.id=1` and `port=9093`
open `server-1.properties` and edit

```
broker.id=1  #must be unique for every replication
listeners=PLAINTEXT://:9093  #port to listen 9093
```

###### replication 2
example :
create or copy from the existing config file. let say we create `server-2.properties` in config folder.
and set `broker.id=2` and `port=9094`
open `server-2.properties` and edit

```
broker.id=2  #must be unique for every replication
listeners=PLAINTEXT://:9094  #port to listen 9094
```

###### start server

replication 1
```
$ ./bin/kafka-server-start.sh config/server-1.properties
```

replication 2
```
$ ./bin/kafka-server-start.sh config/server-2.properties
```

###### create topic
create topic with 2 replication & 2 partitions

```
$ ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 2 --partitions 2 --topic testreplica
```

###### check status
check status replication & partition for topic `testreplica`
```
$ ./bin/kafka-topics.sh --describe --zookeeper localhost:2181 --topic testreplica
```
the result will be

```
Topic:testreplica	PartitionCount:2	ReplicationFactor:2	Configs:
	Topic: testreplica	Partition: 0	Leader: 1	Replicas: 1,2	Isr: 1,2
	Topic: testreplica	Partition: 1	Leader: 2	Replicas: 2,1	Isr: 2,1
```



cheers,

hope you enjoy it


source :
* tutorialspoint.com
* kafka.apache.org
