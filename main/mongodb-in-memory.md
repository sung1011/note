# MongoDB In-Memory

    纯粹的内存数据库 | > 3.2

## usage

```js
    # shell
    mongod --storageEngine inMemory --dbpath

    # config
    storage:
        engine: inMemory
        dbPath: <path>
```

## concurrency

    doc级别的并发

## memroy use

    0.5 * ( RAM - 1024mb )

> shell --inMemorySizeGB

> conf storage.inMemory.engineConfig.inMemorySizeGB 

## durability

    无持久化

## transaction

    无

## deploy

### replset

    in-memory*2 + wiredtiger*1

    wiredtiger 
        hidden: true
        priority: 0

### sharding

    in-memory*2 + wiredtiger*1

    wiredtiger 
        hidden: true
        priority: 0

    分别标记tag进行管理 如:inmem / persisted