```go
func Get(url string) (resp *Response, err error)
```

```go
type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	Header Header
	Body io.ReadCloser
	ContentLength int64
	TransferEncoding []string
	Close bool
	Uncompressed bool
	Trailer Header
	Request *Request
	TLS *tls.ConnectionState
}
```

```go
type ReadCloser interface {
	Reader
	Closer
}
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
```

---

```go
func ReadAll(r Reader) ([]byte, error)
```

---

```go
decoder := json.NewDecoder(res.Body)
var decodedData []User
err = decoder.Decode(&decodedData)
```

```go
jsonData := []byte(`[
        {"Name": "Aragorn", "Age": 87, "Weapon": "Sword"},
        {"Name": "Legolas", "Age": 2931, "Weapon": "Bow"}
    ]`)

    var characters []Character
    if err := json.Unmarshal(jsonData, &characters); err != nil {
        log.Fatalf("Error unmarshalling JSON: %v", err)
    }
```

```go
spaceship := Spaceship{
        Name:      "Serenity",
        CrewCount: 5,
        Active:    true,
    }

    jsonData, err := json.Marshal(spaceship)
```

