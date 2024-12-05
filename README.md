# entsqlite
Allow use `modernc.org/sqlite` in ent with concurrent support.

## Installation
```bash
go get github.com/lib-x/entsqlite
```

## Usage
Just add one line in your import:
```go
import _ "github.com/lib-x/entsqlite"
```

And then using ent as normal
file mode
```go
client, err := ent.Open("sqlite3", "file:./data.db?cache=shared&_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(10000)")
```
```go
// Basic memory database
client, err := ent.Open("sqlite3", "file::memory:?cache=shared")
// Temporary database that's deleted when connection closes
client, err := ent.Open("sqlite3", "file:?mode=memory&cache=shared")
// Named memory database that can be shared between connections
client, err := ent.Open("sqlite3", "file:memdb1?mode=memory&cache=shared")
```

## Connection Parameters

### Basic Parameters
- `file:./data.db` - Database file path
- `cache=shared` - Enable shared cache mode for better concurrency

### Pragma Parameters
Using `_pragma=name(value)` format to set SQLite PRAGMA:

1. `foreign_keys(1)`
   - Enables foreign key constraints
   - Values: 0 (off), 1 (on)
   - Recommended: 1

2. `journal_mode(WAL)`
   - Sets the journal mode
   - Values: DELETE, TRUNCATE, PERSIST, MEMORY, WAL, OFF
   - Recommended: WAL (Write-Ahead Logging) for better concurrency

3. `synchronous(NORMAL)`
   - Controls how SQLite writes to disk
   - Values:
     - OFF: Fastest but least safe
     - NORMAL: Good balance of safety and speed
     - FULL: Slowest but safest
   - Recommended: NORMAL for most cases

4. `busy_timeout(10000)`
   - Sets how long to wait when the database is locked
   - Value in milliseconds
   - Recommended: 5000-10000 (5-10 seconds)

### Additional Parameters
file mode
```go
// Full example with all recommended parameters
dsn := "file:./data.db?" +
    "cache=shared&" +                         // Enable shared cache
    "_pragma=foreign_keys(1)&" +             // Enable foreign keys
    "_pragma=journal_mode(WAL)&" +           // Use WAL mode
    "_pragma=synchronous(NORMAL)&" +         // Normal synchronization
    "_pragma=busy_timeout(10000)&" +         // 10 second timeout
    "_pragma=temp_store(MEMORY)&" +          // Store temp tables in memory
    "_pragma=mmap_size(30000000000)&" +      // 30GB mmap size
    "_pragma=cache_size(-2000)"              // 2MB cache size
```
memory mode
```go
// Optimized memory database configuration
dsn := "file::memory:?" +
    "cache=shared&" +
    "mode=memory&" +
    "_pragma=journal_mode(MEMORY)&" +
    "_pragma=synchronous(OFF)&" +
    "_pragma=foreign_keys(1)&" +
    "_pragma=temp_store(MEMORY)&" +
    "_pragma=cache_size(-2000)"

```
### Performance Tuning Parameters
1. `temp_store(MEMORY)`
   - Store temporary tables in memory
   - Values: DEFAULT, FILE, MEMORY
   - Improves performance for complex queries

2. `mmap_size(30000000000)`
   - Memory-mapped I/O size in bytes
   - Larger values can improve performance
   - Recommended: Adjust based on available RAM

3. `cache_size(-2000)`
   - Database cache size in KB (negative values)
   - -2000 means 2MB cache
   - Adjust based on available memory

### Connection Parameters Comparison

| Parameter | File Mode | Memory Mode | Description |
|-----------|-----------|-------------|-------------|
| cache | shared | shared | Enable shared cache |
| mode | rwc | memory | Database mode |
| _journal_mode | WAL | MEMORY | Journal mode |
| _synchronous | NORMAL | OFF | Sync mode |
| _temp_store | MEMORY | MEMORY | Temp storage |
| _mmap_size | 30GB | N/A | Memory mapping |
| _cache_size | -2000 | -2000 | Cache size in KB |


## Concurrency Support
The WAL mode enables multiple readers and a single writer to operate concurrently. To optimize for concurrent operations:

```go
// Configure client for concurrent access
client, err := ent.Open("sqlite3", 
    "file:./data.db?"+
    "cache=shared&"+                     // Enable shared cache
    "_pragma=journal_mode(WAL)&"+        // Enable WAL mode
    "_pragma=busy_timeout(10000)&"+      // Set busy timeout
    "_pragma=synchronous(NORMAL)",       // Set synchronous mode
)
```

## Error Handling
```go
if err != nil {
    log.Fatalf("failed opening connection to sqlite: %v", err)
}
defer client.Close()
```

## Best Practices
1. Always use WAL mode for concurrent access
2. Set appropriate busy_timeout for your use case
3. Enable foreign_keys for data integrity
4. Use shared cache for better concurrency
5. Adjust cache_size based on your data size
6. Consider using mmap for large databases
7. Monitor database performance and adjust parameters accordingly

## Limitations
1. SQLite is not designed for high concurrency
2. Single writer at a time
3. Limited by filesystem performance
4. Not suitable for high-traffic web applications

## References
- [SQLite Documentation](https://www.sqlite.org/docs.html)
- [SQLite PRAGMA Statements](https://www.sqlite.org/pragma.html)
- [SQLite WAL Mode](https://www.sqlite.org/wal.html)

