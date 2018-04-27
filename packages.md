## Important Go Interface Packages

### Reader
* Equivalent of a read stream in other languages.

### Writer
* Equivalent of a write stream in other languages.

### ReadWriter
* Encapsulates the read-write interface.

### Closer
* Used in closing tcp connections.

### ReadWriteClose
* Encapsulates the read-write-close interfaces.

### Scan
* A way to retrieve data from the stdin


## HTTP Package

### Example request methods
* http.Get(path)
* http.Post(path, contentType, buffer)
* http.PostForm(path, url.Values{})

### Exmaple server methods
* http.Handle(path, handler)
* http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request))
* http.ListenAndServe(tcp address, nil)

## Example networking methods
* Dian(network, address) connects to network address
* Conn () describes a network connection
```go
    package main

    import (
        "io"
        "log"
        "net"
    )

    func main() {
        // Listen on TCP port 2000 on all interfaces
        l, err := net.Listen("tcp", ":2000")
        if err != nil {
            log.Fatal(err)
        }
        defer l.Close()
        for {
            // Wait for a connection
            conn, err := l.Accept()
            if err != nil {
                log.Fatal(err)
            }
            // Handle the connection in a new goroutine
            // The loop the returns to accepting, so that
            // multiple connections may be served concurrently
            go func(c net.Conn) {
                // Echo all incoming data
                io.Copy(c, c)
                // Shut down the connection
                c.Close()
            }(conn) // passing in conn to guarantee an accepted connection gets used
        }
    }
```