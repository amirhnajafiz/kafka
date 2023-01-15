#!/usr/bin/env bash


KAFKA_VERSION="3.3.1"
OUTPUT="kafka_2.13.tgz"


# download chosen version of kafka
curl https://dlcdn.apache.org/kafka/"$KAFKA_VERSION"/kafka_2.13-"$KAFKA_VERSION".tgz --output $OUTPUT

# extract kafka downloaded files
tar -xzf $OUTPUT