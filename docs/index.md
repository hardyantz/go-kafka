
this post is for beginner who want to learn basic kafka with golang. it is far from advance. if you think you are more advance than this post. please share some thought. thanks.

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

source :
* tutorialspoint.com
* kafka.apache.org


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
type something in the new line in the consumer screen
```
> test one
> test two
```

open the consumers terminal and the results will be
```
test one
test two
```

try to type something in your producer terminal, it will show in your consumer terminal

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

cheers,
hope you like it.