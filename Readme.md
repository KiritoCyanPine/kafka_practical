# Kafka Tutorial and Setup Guide

This repository contains a simple Kafka producer and consumer implementation in Go, along with a Docker Compose setup for running Kafka in KRaft mode. The producer simulates IoT device metrics, while the consumer processes these metrics and exposes Prometheus metrics for monitoring.

## Journey Log :

// create a table for the journey log

| Date       | Activity                                                | Link                                |
| ---------- | ------------------------------------------------------- | ----------------------------------- |
| 2026-10-20 | Set up Kafka cluster using Docker Compose in KRaft mode | [Log-1](docs/Journey/Instance-1.md) |

### Kafka cli commands for setup, troubleshooting and metadata

Below commands will be ran in the terminal where Kafka is installed and running.

#### Creating a topic :

The below command creates a topic named "metrics" with 6 partitions and a replication factor of 1.

```bash
kafka-topics --create --topic metrics --partitions 6 --replication-factor 1 --bootstrap-server localhost:9092
```

#### Describing a topic :

The below command describes the topic named "metrics" and provides details about its partitions, replication factor, and other configurations.

```bash
kafka-topics --describe --topic metrics --bootstrap-server localhost:9092
```

#### Listing all topics :

The below command lists all the topics available in the Kafka cluster.

```bash
kafka-topics --list --bootstrap-server localhost:9092
```

#### Reading messages from a topic :

The below command starts a console consumer that reads messages from the "metrics" topic, starting from the beginning of partition 1.

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic metrics --partition 1 --from-beginning
```
