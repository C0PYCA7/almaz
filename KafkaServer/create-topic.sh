#!/bin/bash

KAFKA_BIN_DIR=./bin
BROKER=localhost:9092

create_topic() {
  local topic="$1"
  local partitions="${2:-1}"
  local replication="${3:-1}"

  $KAFKA_BIN_DIR/kafka-topics.sh \
    --create \
    --if-not-exists \
    --bootstrap-server "$BROKER" \
    --replication-factor "$replication" \
    --partitions "$partitions" \
    --topic "$topic"
}

create_topic dbTopic 3 1
create_topic reportTopic 3 1