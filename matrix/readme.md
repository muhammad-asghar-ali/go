# Matrix - Computer Systems Management

## Overview

The **Matrix - Computer Systems Management** is a project that demonstrates the implementation of gRPC for managing computer system components. The system simulates the management and monitoring of hardware components such as keyboards, laptops, memory, processors, screens, and storage devices. It leverages **Protocol Buffers** for defining data structures and **gRPC** for communication between services.

This project models real-world interactions with computer systems using multiple components, which are defined as Protocol Buffers in the `proto/` directory.

### Key Features:

- **gRPC-based communication**: All inter-service communication is done using gRPC, enabling high-performance and scalable communication.
- **Protocol Buffers (protobuf)**: All data exchanged between services is serialized and deserialized using Protocol Buffers, ensuring efficient and compact data transmission.
- **Real-life use case**: This system models a real-world scenario of computer management with various services interacting with each other, making it suitable for testing and exploring the potential of gRPC in a production-like environment.

### Project Structure:

```
matrix/
├── cmd/              # Main server entry point
├── internal/
│   ├── fns/          # Utility functions (random data generators)
│   ├── pb/           # Generated protobuf code
│   │   └── pbconnect/ # Connect RPC handlers
│   ├── serializer/   # Data serialization (JSON, file)
│   ├── services/     # gRPC service implementations
│   └── tmp/          # Experimental code
└── proto/            # Protocol Buffer definitions
    ├── keyboard.proto
    ├── laptop.proto
    ├── laptop_service.proto
    ├── memory.proto
    ├── processor.proto
    ├── screen.proto
    └── storage.proto
```
