#!/usr/bin/env bash

nano /etc/systemd/system/zookeeper.service

#  [Unit]
#  Requires=network.target remote-fs.target
#  After=network.target remote-fs.target
#
#  [Service]
#  Type=simple
#  User=kafka
#  ExecStart=/home/kafka/kafka/bin/zookeeper-server-start.sh /home/kafka/kafka/config/zookeeper.properties
#  ExecStop=/home/kafka/kafka/bin/zookeeper-server-stop.sh
#  Restart=on-abnormal
#
#  [Install]
#  WantedBy=multi-user.target

nano /etc/systemd/system/kafka.service

#  [Unit]
#  Requires=zookeeper.service
#  After=zookeeper.service
#
#  [Service]
#  Type=simple
#  User=kafka
#  ExecStart=/bin/sh -c '/home/kafka/kafka/bin/kafka-server-start.sh /home/kafka/kafka/config/server.properties > /home/kafka/kafka/kafka.log 2>&1'
#  ExecStop=/home/kafka/kafka/bin/kafka-server-stop.sh
#  Restart=on-abnormal
#
#  [Install]
#  WantedBy=multi-user.target

systemctl daemon-reload

systemctl enable --now kafka

systemctl status kafka zookeeper