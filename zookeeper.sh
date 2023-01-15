#!/usr/bin/env bash


KAFKA_PATH="kafka_2.13-3.3.1"

# enter the kafka downloaded package
cd $KAFKA_PATH || exit


# start the ZooKeeper service
bin/zookeeper-server-start.sh config/zookeeper.properties

# start the Kafka broker service
bin/kafka-server-start.sh config/server.properties