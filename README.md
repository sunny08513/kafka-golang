# kafka-golang

Apache Kafka is an open-source distributed event streaming platform capable of handling trillions of events a day. Originally developed by LinkedIn and later open-sourced, Kafka is written in Scala and Java and is part of the Apache Software Foundation.

### Key Concepts and Terminology in Kafka

1. **Broker**:
   - A Kafka broker is a server that stores message data in Kafka.
   - A Kafka cluster is made up of multiple brokers to ensure fault tolerance and high availability.
   - Brokers receive messages from producers, store them, and serve them to consumers.

2. **Topic**:
   - Topics are categories or feed names to which messages are published.
   - Topics in Kafka are always multi-subscriber; that is, a topic can have zero, one, or many consumers that subscribe to the data written to it.
   - Each topic is split into partitions.

3. **Partition**:
   - A topic can have multiple partitions, which are logs ordered by time.
   - Partitions allow for scalability and parallel processing.
   - Each partition is an ordered, immutable sequence of records that is continually appended to—a structured commit log.
   - Each record in a partition is assigned a unique sequential id number called the offset.

4. **Producer**:
   - Producers are processes that publish (write) messages to one or more Kafka topics.
   - Producers can choose which partition to send messages to, enabling better load balancing and parallelism.

5. **Consumer**:
   - Consumers are processes that subscribe to topics and process the feed of published messages.
   - Each consumer is part of a consumer group.
   - A consumer group can have multiple consumers, which helps with parallel processing.

6. **Consumer Group**:
   - A consumer group is a group of consumers sharing a common group identifier.
   - Each consumer in the group processes messages from different partitions, and Kafka ensures that each partition is consumed by only one consumer in a consumer group.

7. **Offset**:
   - The offset is a unique identifier for each record within a partition.
   - It is used to track the position of the consumer in the log.
   - Consumers use offsets to track their position in each partition.

8. **Replication**:
   - Kafka replicates topic partitions across multiple brokers to ensure reliability and fault tolerance.
   - Each partition can have multiple replicas (a leader and followers).
   - The leader handles all reads and writes for the partition while followers replicate the data.

9. **ZooKeeper**:
   - ZooKeeper is used by Kafka to manage and coordinate the Kafka brokers.
   - It is responsible for leader election for partitions, maintaining configuration information, and managing service discovery.

10. **Log Compaction**:
    - Log compaction ensures that Kafka retains at least the last known value for each key, which is useful for systems that need to restore their state after a crash or restart.
    - It deletes old records, retaining only the most recent update for each key within the log.

11. **Streams API**:
    - Kafka Streams is a client library for building applications and microservices, where the input and output data are stored in Kafka clusters.
    - It allows for real-time processing of data.

12. **Connect API**:
    - Kafka Connect is a tool for connecting external systems (databases, key-value stores, search indexes, file systems) with Kafka.
    - It comes with a large ecosystem of connectors.

### How Kafka Works?

- **Producers** send data to Kafka topics.
- Each topic has multiple **partitions**, which allow Kafka to parallelize the handling of data.
- Kafka **brokers** store these partitions on disk and manage the replication of data across the cluster for fault tolerance.
- **Consumers** subscribe to topics and read data from partitions. Kafka ensures that each consumer within a **consumer group** gets a fair share of partitions.
- The **offset** helps consumers track which messages they have already processed.
- **ZooKeeper** coordinates the cluster, keeps track of the brokers, topics, and partitions, and handles leader elections.

### Why Use Kafka?

- **Scalability**: Kafka's partitioning allows for massive scalability.
- **Fault Tolerance**: Data replication across brokers ensures no data loss.
- **High Throughput**: Kafka can handle high throughput of messages due to its efficient storage mechanism.
- **Durability**: Messages are persisted on disk, ensuring durability.
- **Real-Time Processing**: Kafka Streams enables real-time data processing.
- **Integration**: Kafka Connect simplifies integration with other systems.

Kafka is widely used for building real-time streaming data pipelines and applications that adapt to the data streams. It is a foundational technology for modern data architectures.

# USE CASE
In a typical e-commerce platform, the order service can get confirmation through a combination of synchronous and asynchronous communication mechanisms. When an order is placed, it publishes the order details to Kafka, and other services (such as inventory, payment, and shipping) consume and process these details. To provide confirmation back to the order service, we can implement a feedback loop using Kafka topics.

Workflow for Order Confirmation
1. Order Placement:

The order service receives a new order request.
It publishes the order details to a Kafka topic, say order-events.

2. Order Processing:

Various services consume the order-events topic:
Inventory Service: Checks the availability of the items.
Payment Service: Processes the payment.
Shipping Service: Prepares the shipping details.

3. Publishing Processing Results:

Each service publishes the result of its processing to a specific Kafka topic:
Inventory Service: Publishes to inventory-status.
Payment Service: Publishes to payment-status.
Shipping Service: Publishes to shipping-status.

4. Order Confirmation Service:

A dedicated service (or the original order service) consumes the inventory-status, payment-status, and shipping-status topics.
It aggregates the results and determines the overall status of the order.
Once all necessary steps are completed and successful, it publishes an order-confirmation event to a Kafka topic, say order-confirmations.

5. Order Service Gets Confirmation:

The order service consumes the order-confirmations topic.
When it receives a confirmation event for the specific order, it updates the order status in its database and notifies the customer.

# Step By step guide to install kafka

1. **Install Kafka on docker**
 docker-compose-kafka.yml file defines two services: zookeeper and kafka. The Kafka service is configured to expose ports 9092 and 9093.

Open a terminal in the directory where the docker-compose.yml file is located and run the following command to start the Kafka and Zookeeper containers:

 ```
docker-compose -f docker-compose-kafka.yaml up -d
 ```

 This command will download the required Docker images and start the Kafka and Zookeeper services in detached mode (-d).

2.  **Create a Kafka Topic**

Create a Kafka topic using the following command:
```
docker exec -it <kafka-container-id> /opt/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic my-topic
```
3. **Produce and Consume Messages**
Produce messages:
```
docker exec -it <kafka-container-id> /opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic my-topic
```
Consume messages:

```
docker exec -it <kafka-container-id> /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic my-topic --from-beginning
```

4. **Stop and remove containers**

```
docker-compose  -f docker-compose-kafka.yaml down
```


 