### vertical partitoning
need to learn about this

### horizontal partitioning

### how is data integrity insured in case of failure for sql database
- persisted on disk
- redo/undo log in queue

### how do we improve performance of sql db architecture
- index the data
- sharding
- master slave relationship
- partitioning processes or sharding

### what makes no sql more scalable than a sql database
- don't have foreign keys or concerns about atomicity (needing to update rows on sepearate tables or servers)
- SQL: ACID (all or nothing, no altering structure, transactions don't contend with one another, permanent)
  - acid comes at performance cost, but it is highly reliable
- NoSQL: BASE (support partial failures, values may change consistency is eventual, )

### advantages of SQL database
- acidity
- store relational data for complex information
- efficient and quick, long established standards
- powerful means to combine separate dta to express complex relationships

### primary v foreign key
- primary keys are unique, foreign keys refer to another table

### normalized v denormalized
- denormalized data tends to be more optimized for time than space
- normalized databases involve multiple tables

### inner v outer join
- union v intersection

### db optimizations
- in memory dbs (SAP HANA, TimesTen, Aerospike, Redis)
- optimized search engines (elasticsearch, apache Solr, CrateDB)
- realtime subscribing (rethinkDB, OrientDB)
- Realtime sync (firebase, pubnub, pusher)

### in memroy dbs
- relies on main memory storage
- faster than disk, non invasive and small
- cost more, lack of expertise, smaller size

### graphql
- query language for apis, backed by facebook
- define data you want
- replace multiple rest requests
- advantages over REST
  - easy to consume or produce
  - types
- queries retrieve data, mutations modify them
- resources: graphqlhub.com \ howtographql.com 

### HTML5
- canvas
- tags, support for video
- supports offline data storage (local and session storage)
- data attributes
- div is block and span is inline

### CSS3
- border-radius, box-shadow and text-shadow
- transform / animate
- flexbox
- CSS grid! 

### HTTP

### SOAP
Simple Object Access Protocol 

### RESTful APIs
representational state transfer (REST) technology
RFC 2616 
all calls are stateless; nothing can be retained by the RESTful service between executions

### PUT v PATCH

### 

## Redis v memcache
Redis tends to be superior, but there are certain use cases where memcache may be better
Main differences are:
- Serialization - Memcached is optimal if you are simply storing html strings, it is slightly faster with this but it does not have the ability to store native data, it serializes instead. Redis on the other hand can store serialized data, but the associated meta data may take slightly longer for very simple use cases.
- Scaling -- Memcached is slightly better here because it is multithreaded so it quickly scales but loses data potentially. Redis is mostly single-threaded but can be expanded with clustering, just slightly mroe complex to set up.
- Data eviction - Redis is better because you can control how it evicts data, whereas Memcached is mainly LRU

## Redux
First Principle: All state management is concentrated in a single place
Second Principle: The state is read only, can only be changed through dispatching actions
Third Principle: It is essential to use pure functions when making changes to the state

Reducers take the old state and an action to return a new state, which is a new object, not a mutated version of the original. However, it can use the old references if they don't change, that helps make it fast.

createStore needs to be called with the reducer that controls it
store has a getState() and a dispatch(action) methods, as well as subscribe(callback) in order to update the UI

### Reselect
Reselect is special because whereas react does shallow comparisons when updating state (and will therefore make a new array when all the inner elements of an array are the same), the reselect will look at each of the values as it selects the state and if there are no changes it will use memoization to avoid computing and just use the old state.

the createSelector joins together the different states, and combines them into a single selector at the end.

### Processes and Threads

Processes
  - Resource Ownership
    - Includes a virtual address space to hold process image
    - Prevents unwanted interference

  - Scheduling/Execution
    - Has an execution path that may be interleaved with other processes
    - Has an execution state (running, ready, etc)
    - Scheduled and dispatched by the OS

  - Unit of dispatching is a thread or lightweight process
  - unit of resource ownership is a process or task
  - multithreading is the ability of an OS to support multiple concurrent pahts of execution within a signle process

  - Single threaded means single thread per process
  - Java runtime environment is a one process multiple thread system

  Threads
    - have an execution state
    - save thread context, have per thread static storage
    - have access to shared memory of its process
    - Thread states are: running, ready, blocked
    - Thread operations are: spawn, block, unblock, finish

  ### Graph Algorithm Notes

  - BFS search can gurantee shortest path if weighted graph
  - If unweighted, we can still find shortest path using BFS and a priority queue as opposed to a FIFO queue. 
  - The priority of each node in the queue is its distance from teh start node. 
  - This is Dijkstra's algorithm, only works with positive edges.
  
  - Queue's can be implemented with two stacks. 
  - Priority queue's can be implemented with a binary heap
  - DFS usually takes less memory than BFS, because you immediately jump into the call stack as oppposed to storing up a queue of order.
  - A* is a modification of Dijkstra's algorithm which uses a heuristic added to the distance from the start node as the priority in the priroity queue.
  - Floyd-Warshall algorithm is good for shortest paths on a graph with negative edge weights.
  - Dijkstra's algorithm can find the shortest path if some edges have weight 0.
  

### ArrayList resizing

- Implementation depends, but generally has amortized constant time
- Probably doubles size and copies over.

### Hash Tables
- Insert is O(1), Update is O(1), Retrieve is O(1);

### BST
- Insertion, Search, and Delete are all O(logn)
- Two common balancing mechanisms are AVL and Red Black Trees

### Red Black Trees
- Paths from the root to a leaf can't differ by more than a factor of two.
- Good enough to ensure height is only a constant multiple of log(n)
- All node are red and black
- Red's can't connect to other reds
- Same numberr of black nodes on every path from root to leaf.
- Max height is alternating red and black, min height is when you have a bunch of black nodes in a row.

### BTrees
- Memory Tradeoffs:
  - Huge between speed and size in memory
  - SRAM is fast but expensive
  - Great for on disk storage because of huge branching factor
  - Minimizes the number of blocks to read on lookup
  - Common in databases, file systems, in mem data structures


### Memory Hierarchy
1. Register -> 256B - 8KB -> .25 - 1 ns
2. L1 Cache -> 16KB - 64KB -> 1ns - 5ns
3. L2 Cache -> 1MB - 4MB -> 5ns - 25ns
4. Main Memory -> 4GB - 256GB -> 25ns - 100ns
5. Hard Disk -> 1TB+ -> 3 - 10ms
6. Network/Cloud -> Lots -> 10 - 2000ms

### AVL Tree
  - Two subtrees have heights that differ by no more than one

### MinHeap
  - Essentially a complete binary tree
  - Complete means a binary tree that all levels are fully filled except for the last
  - Full tree means that each node has either 0 or 2 child nodes
  - Methods include insert, extractMin. 
  - Insert O(logn), place at bottom right, bubble up
  - Extract Min, swap top and bottom right, remove bottom right, bubble down top
  - Asymptotically identical to BST, but consume less memory and faster in practice

  ### Randomized Data Structures
  - Involve tradeoffs to make partial judgments mroe quickly
  - Bloom filter (stashes buckets, can quickly determine if something does not exist, but being sure it does is harder)
  - Count-min sketches and Hyper Log Logs are similar i think
  - Treap
    - Form of BST maintains a dynamic set of ordered keys. 

### Normalization
  - Primary Key
    - Column with unique value for each row
    - Each table should have one
  - Foreign Key
    - Define relationships between tables, match primary key
  - Composite Key
    - Key made of more than one column
    - Prevent a table from using same combo twice
    - Sounds like a key on a join table
  - 1:1 Relationship
    - can be useful for optimizations, but typically all one to one relationshps are on the same table
  - 1:N (One to many)
    - Links multiple rows in a child table to a single row in a parent, with a foreign key.
  - N:M (Many to many)
    - Creates a join table containing foreign keys to both parents
  - Up to 3NF is normalized
  - 1NF (First Normal Form)
    - Tables must not contain repeating groups of data
    - Don't have multiple identical columns or serialized info.
    - Because we have an unknown number of phone numbers, we extract it to another table
  - Second Normal Form (2NF)
    - No field should only be partially dependent on any candidate key in the table.
    - In other words, we don't want to worry about making sure multiple columns maintain consistency
    - Must move these fields out of the table into own tables
    - Creating FKs to link the new tables to the main tables
  - Third Normal Form (3NF)
    - Columns should depend only upon the primary key of the table
    - Tying countries to people limits countries.
  - Normalization is a design goal
    - Can impact performance and resource conservation
    - 1NF is essential, but others are case by case
    - Generally 3NF is worthwhile as a minimum. 
  - Denormalization can speed up reads
    - Enforce data integrity by using check constraints

### Database Indexes
  - Indexes speed up reads for the column they are on
  - Make writes slower and increase memory requirements
  - Often implemented with B trees
  - Composite indexes 

### SQL Explain
  - Shows the query plan for any query
  - Query plan is a tree of plan nodes
  - Scan nodes at the bottom and return raw rows from table
  - Different types of scan nodes
  - EXPLAIN has one line for each node in the plan tree
  - Planner tries to minimize total execution cost for the plan
  - Planner shows its plans, somtimes a sequential scan with a filter
  - Other times uses a Bitmap Heap Scan, or Bitmap index scan

### Postgres Query Planning
  - Understanding it helps you to optimize your database with indexes to improve performance
  - Explain Generic, Analyze (actually runs), and Verbose (tay away)

### NoSQL 
  - Good for unstructured schemaless data
  - Favors denormalization because of no joins
  - Performance versus data loss tradeoff
  - Very easy to scale NoSQL by adding nodes *

### ORM is Object Relational Mapping
  - Uses object oriented paradigm to query a database. 
  - Makes it easier to use, but can drag performance
  - Can protect you from SQL injection attacks

### HTTP
  - Application-level protocol
  - Stateless
  - TCP/IP Delivers HTML, images, query results
  -  and media independent
  - Content type using MIME type must be specified
  - Web caching is a feature of HTTP to minimize network traffic
  - Seems like a CDN
  - Can be stored in Browswer cache, Intermediary proxy, reverse cache
  - Websockets
    - Uses same TCP/IP connection but bidirectional
  - TCP v UDP (Transmission control v User Data)
  - TCP , UDP is more of a broadcast, less secure
  - TCP is connection, UDP is connectionless
  - TCP is handshake, UDP is spamming
  - TCP is high reliability, UDP is fast and efficient

### Mutexes and Sempahores

Semaphores are created by Dijkstra
  - binary semaphore or counting semaphore
  - controls access between threads to a shared resource
  - problems include accidental release, deadlock, priority inversion, semaphore as a signal
  - Mutex: Different from semaphore because of principle of ownership
  - only the task that locks a mutex can unlock it.
  - Mutex is safer than semaphore, but circular deadlock and non cooperation still possible.

### Locks
  - Advisory locks, each thread cooperates by acquiring the lock before accessing the corresponding data.
  - A binary semaphore is a lock, 
  - Atomic interactions such as test and set, fetch and add, compare and swap

### Monitors 
  - A synchronization that allows threads to have both mutal exclusion and the ability to wait block for a certain condition to become true.
  - Monitors consist of a mutex lock and condition variables

### Readers-Writer Lock
  - Allows concurrent access for read only operations