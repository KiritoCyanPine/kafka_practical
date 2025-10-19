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
`kafka-topics --create --topic metrics --partitions 6 --replication-factor 1 --bootstrap-server localhost:9092`

#### Describing a topic :

The below command describes the topic named "metrics" and provides details about its partitions, replication factor, and other configurations.
`kafka-topics --describe --topic metrics --bootstrap-server localhost:9092`

#### Listing all topics :

The below command lists all the topics available in the Kafka cluster.
`kafka-topics --list --bootstrap-server localhost:9092`
