<h1 align="center">
    Install
</h1>

<br />

Make sure to have java installed on your VM.

```shell
apt-get install default-jdk -y
```

## User and Download

First we create a new user named ```kafka``` to download
kafka source codes from kafka archive.

```shell
adduser kafka
adduser kafka sudo

su - kafka
wget https://dlcdn.apache.org/kafka/2.7.2/kafka-2.7.2-src.tgz

tar -xvzf kafka-2.7.2-src.tgz
mv kafka-2.7.2-src kafka

exit
```

## Kafka and Zookeeper

Now we are going to setup kafka and zookeeper.

```shell
nano /etc/systemd/system/zookeeper.service
```

```
[Unit]
Requires=network.target remote-fs.target
After=network.target remote-fs.target

[Service]
Type=simple
User=kafka
ExecStart=/home/kafka/kafka/bin/zookeeper-server-start.sh /home/kafka/kafka/config/zookeeper.properties
ExecStop=/home/kafka/kafka/bin/zookeeper-server-stop.sh
Restart=on-abnormal

[Install]
WantedBy=multi-user.target
```

```shell
nano /etc/systemd/system/kafka.service
```

```
[Unit]
Requires=zookeeper.service
After=zookeeper.service

[Service]
Type=simple
User=kafka
ExecStart=/bin/sh -c '/home/kafka/kafka/bin/kafka-server-start.sh /home/kafka/kafka/config/server.properties > /home/kafka/kafka/kafka.log 2>&1'
ExecStop=/home/kafka/kafka/bin/kafka-server-stop.sh
Restart=on-abnormal

[Install]
WantedBy=multi-user.target
```

Now let's create systemd units for both.

```shell
systemctl daemon-reload
systemctl enable --now kafka
systemctl status kafka zookeeper
```

## Dashboard with CMAK

CMAK is an open source tool for managing and monitoring Kafka services developed by Yahoo.
Let's connect it to our kafka cluster.

```shell
apt-get install git -y
git clone https://github.com/yahoo/CMAK.git
```

Edit the config file:

```shell
nano ~/CMAK/conf/application.conf
```

```
kafka-manager.zkhosts="kafka-manager-zookeeper:2181"
kafka-manager.zkhosts=${?ZK_HOSTS}
cmak.zkhosts="localhost:2181"
cmak.zkhosts=${?ZK_HOSTS}
```

```shell
cd ~/CMAK
./sbt clean dist

cd ~/CMAK/target/universal
unzip cmak-3.0.0.5.zip

cd cmak-3.0.0.5
bin/cmak
```

At this point, CMAK is started and listening on port 9000.