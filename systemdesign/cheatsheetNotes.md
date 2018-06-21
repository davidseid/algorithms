## Cheat Sheet Notes

### Approach

1. Scope
  - Use cases (who and how)
  - Constraints (traffic and data handling)
  - RPS, Request Types, data written per second, data read per second
  - Special requirements (read v write)

2. Architecture Design
  - Basic Components Sketch
    - Application Service Layer 
    - Different Services required
    - Data storage layer
    - Load balancer, partition, database cluster (master/slave), caching system

3. Component Design 
  - Component and specific APIs for each
  - OOD for functionalities
    - map features to modules
    - singletons, inheritance
  - Database schema design

4. Bottlenecks
  - User requeests? 
  - Huge data? 
  - Memory caching? 

5. Scaling Design
  - Vertical: CPU RAM
  - Horizontal: Machines
  - Caching: 
    - Application caching: integration into app code
    - Database caching: free, tuning
    - In memory cache: memcached or Redis
    - Precalculating results, pregenerate expensive indexes
    - Frequently used data
  - Load balancing: 
    - Smart client
    - Hardware ($$$ but reliable)
    - Software (hybrid)
  - Database replication
    - related to normalization, removing data ambiguity
    - distributed data
  - Database partitioning
    - Decomposing relational data by row or column
  - Map-Reduce 
    - added layer for intensive operations in reasonable time
    - Hadoop, Hive, HBase
  - Platform Layer
    - Microservices, allows for easy flexibility with infastructure

### Key Topics

1. Concurrency
  - threads, locks, starvation? 
  - threads are smaller processes that share the same memory space as their process, allow for concurrent work, must be managed to make sure they do not attempt to access the same resource simultaneously
  - locks are a mechanism for managing this, such as semaphores and mutexes
  - starvation is when a process can't access its resource to work, caused by mutex, resource leaks, or DDOS attacks

2. Networking
  - IPC, TCP/IP
    - IPC is the means by which OS dictate processes managing shared data
    - Transmission control protocol, internet protocol
      - byte stream
      - TCP v UDP (TCP is reliable, UDP is fast)
    - client/server relationship, synchronization
  - throughput v latency

3. Abstraction
  - OS
  - Filesystem
  - Database
  - Modern OS caching levels

4. Real-World Performance
  - Speed of things
  - RAM
  - Disk
  - SSD
  - Network

5. Estimation
  - back of envelope calculations for memory storage, throughput requirements, databbase size, number of users, etc.

6. Availability and Reliability
  - How can things fail?
  - How to cope with network failures?
  - Durability?

### Web App System Design Requirements


