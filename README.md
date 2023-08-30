# cassandra-with-go

In `cassandra-with-go/` directory

### **Step 1: Set up Cassandra**

1.1 **Install Cassandra**:
If you're using macOS, the easiest way is with Homebrew:

```bash
brew install cassandra
```

1.2 **Start Cassandra**:

```bash
cassandra -f
```

1.3 **Cassandra CQL Shell**:
You can interact with Cassandra using the CQL (Cassandra Query Language) shell. To start it, just type:

```bash
cqlsh
```

### **Step 2: Create a Sample Database**

Using the CQL shell, let's create a simple keyspace and table:

```
cqlCopy code
CREATE KEYSPACE demo WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE demo;

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT,
    age INT
);

```

### **Step 3: Set Up Go and Cassandra Go Driver**

3.1 **Install Go Driver**:
You'll need the Go driver for Cassandra. To get it, run:

```bash
go get github.com/gocql/gocql
```

---

### **Step 4: Run Go Program**

```bash
go run main.go
```
If everything goes as planned, the output will indicate that a user "Alice" with age 30 has been added and then retrieved from the Cassandra users table.