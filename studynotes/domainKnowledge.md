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