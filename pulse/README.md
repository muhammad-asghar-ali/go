# Pulse: Real-Time Notification System with Go and Kafka

## Overview

**Pulse** is a real-time notification system built using Go and Kafka. This system includes a **producer** to send notifications and a **consumer** to process them.

## Features

- **Producer**: Sends notifications to Kafka topics.
- **Consumer**: Listens to Kafka topics and processes notifications in real-time.
- Written in Go, leveraging Kafka for reliable message queuing and delivery.

---

## Quick Start

### Prerequisites

1. Install [Go](https://golang.org/dl/) (1.20+ recommended).
2. Install [Kafka](https://kafka.apache.org/downloads).
3. Ensure `zookeeper` and `kafka-server` are running:
   ```bash
   bin/zookeeper-server-start.sh config/zookeeper.properties
   bin/kafka-server-start.sh config/server.properties
   ```

### Run the Application

1. Start the Kafka Consumer:
   `go run cmd/consumer/main.go`

2. Start the Kafka Producer:
   `go run cmd/producer/main.go`

## Testing

### Sending a Notification

Send a POST request to the producer to send a notification:

```
curl -X POST http://localhost:8080/send \
-d "fromID=4&toID=1&message=Lena liked your post: 'My weekend getaway!'"
```

**Response:**

```
{"message":"Notification sent successfully!"}
```

### Retrieving Notifications

Fetch notifications for a specific user (in this case, toID=1) using the consumer:

````
curl http://localhost:8081/notifications/1
```
````

**Response:**

```
{
  "notifications": [
    {
      "from": {"id": 4, "name": "Lena"},
      "to": {"id": 1, "name": "John"},
      "message": "Lena liked your post: 'My weekend getaway!'"
    }
  ]
}
```

### Future Improvements

1. Add detailed logging for monitoring.
2. Implement message retries for failed deliveries.
3. Extend support for multiple notification types.
4. Use environment varibales
