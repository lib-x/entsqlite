# entsqlite
Allow use `modernc.org/sqlite` in ent
just add one line in your import
```go
import _ "github.com/lib-x/entsqlite"
```
and then using ent as normal
```go
client, err := ent.Open("sqlite3", "file:./data.db")
```
