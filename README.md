# teivah/multilock

multilock is a Go library allowing to store multiple `sync.Mutex` or `sync.RWMutex`.

The internal data structure, depending on the use case, has either a fixed or a variable length.
Accessing a variable structure means acquiring a shared lock first.

## Structures

* Fixed length structure of Mutex: `multilock.Fixed`
* Fixed length structure of RWMutex: `multilock.RWFixed`
* Variable length structure of Mutex: `multilock.Var`
* Variable length structure of RWMutex: `multilock.RWVar`

## Examples

```go
const multilockLength = 5
const customStructLength = 1000

// Initialize a fixed length multilock structure
mlock := multilock.NewFixed(multilockLength)

// Create custom structures
maps := make([]map[string]string, customStructLength)
for i := 0; i < 1000; i++ {
	maps[i] = make(map[string]string)
}

// Retrieve a lock for a given maps index
mutex := mlock.Get(maps[42])
mutex.Lock()
defer mutex.Unlock()
``` 