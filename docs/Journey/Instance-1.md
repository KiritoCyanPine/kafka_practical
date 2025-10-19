# Journey Log

Learnt about Kafka cli commands for setup, troubleshooting and metadata.

- Created an actual Kafka cluster using Docker Compose in KRaft mode.

Used the below docker-compose.yml file to create the Kafka cluster.

```yaml
# Docker Compose file for Kafka in KRaft mode

# Creating a named volume to persist Kafka data
volumes:
  kafka-data:

# Defining the Kafka service
services:
  kafka:
    image: confluentinc/cp-kafka:7.8.3
    container_name: kafka
    ports:
      - '9092:9092'
    environment:
      KAFKA_KRAFT_MODE: 'true' # Enable KRaft mode
      CLUSTER_ID: '1L6g7nGhU-eAKfL--X25wo' # You can generate a UUID for a kafka cluster
      KAFKA_NODE_ID: 1 # Unique ID for the broker

      KAFKA_PROCESS_ROLES: broker,controller # A Kafka node can take 2 roles: broker Or  controller Or both (only in Kraft mode)

      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@localhost:9093 # defines which nodes will participate on the cluster control decisions.

      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1 # there is only one copy of the metadata in KRaft mode

      # open Kafka listeners for communication; one for data Traffic and one for controller communication
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9093

      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092 # Advertised listener for clients to connect to Kafka broker
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER # specify which listener will be used for controller communication
      KAFKA_LOG_DIRS: /tmp/kraft-combined-logs # log storage directory
    volumes:
      - kafka-data:/var/lib/kafka/data # persist kafka data
```

- Learning about kafka-topics command to create, describe and list topics in Kafka cluster.
- Ran a producer and consumer application to produce and consume messages to/from the Kafka cluster.
