# cassandra test

## setup cassandra

```
docker-compose up -d

```

## create keyspace, table
```
docker exec -it cassandra cqlsh

CREATE KEYSPACE test WITH REPLICATION =
 {'class': 'SimpleStrategy', 'replication_factor': 2};

 USE test;

 CREATE TABLE test_table (
  id text,
  test_value text,
  PRIMARY KEY (id)
 );
```
## build

```
GOOS=linux ARCH=amd64 go build && docker build -t sample .
```

## run

```
docker run --net cassandra_default --rm sample
```
