#!/usr/bin/env bash

adduser kafka

adduser kafka sudo

su - kafka
wget https://dlcdn.apache.org/kafka/2.7.2/kafka-2.7.2-src.tgz

tar -xvzf kafka-2.7.2-src.tgz
mv kafka-2.7.2-src kafka

exit