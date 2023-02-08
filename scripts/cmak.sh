#!/usr/bin/env bash

apt-get install git -y
git clone https://github.com/yahoo/CMAK.git

nano ~/CMAK/conf/application.conf

#  kafka-manager.zkhosts="kafka-manager-zookeeper:2181"
#  kafka-manager.zkhosts=${?ZK_HOSTS}
#  cmak.zkhosts="localhost:2181"
#  cmak.zkhosts=${?ZK_HOSTS}

cd ~/CMAK || exit
./sbt clean dist

cd ~/CMAK/target/universal || exit
unzip cmak-3.0.0.5.zip

cd cmak-3.0.0.5 || exit
bin/cmak

# localhost:9000