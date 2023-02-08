#!/usr/bin/env bash

cd /home/kafka/kafka || exit
./gradlew jar -PscalaVersion=2.13.3

chown -R kafka:kafka /home/kafka/kafka