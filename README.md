# Go Programming: Introduction and Basics

## Table of Contents
- Introduction to Go
- Installation of Go
- Setting up Go Workspace
- Basic Concepts in Go
  -Types
  - Functions and Methods
  -Structs
  -Channels
  -Mutex
- Building RESTful APIs with Go
  - CRUD Operations
  - MongoDB Integration

## Introduction to Go

Go (Golang) is an open-source, statically typed, compiled programming language created by Google. It is designed for simplicity, efficiency, and reliability. Go is often used for building scalable web applications, cloud infrastructure, and microservices due to its concurrency support and fast execution.

### Key Features of Go
- Simplicity and ease of use
- Built-in support for concurrent programming (goroutines and channels)
- Static typing and garbage collection
- Fast compilation and execution
- Rich standard library

## Installation of Go

### Windows
1. Download the Go installer from the official Go website (https://golang.org/dl/).
2. Run the installer and follow the instructions.
3. Verify the installation:
   ```bash
   go version
   ```

### Linux/MacOS
1. Use the following command to install Go:
   ```bash
   sudo apt-get install golang
   ```
2. Verify the installation:
   ```bash
   go version
   ```

## Setting up Go Workspace

1. Create a directory for your Go projects:
   ```bash
   mkdir ~/go-workspace
   ```
2. Set the Go environment variables:
   ```bash
   export GOPATH=~/go-workspace
   export PATH=$PATH:/usr/local/go/bin
   ```

3. Create a `src` folder inside `go-workspace` for your Go source code:
   ```bash
   mkdir -p ~/go-workspace/src
   ```

4. Start your first Go program:
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```

## Basic Concepts in Go

### Types

Go has a variety of built-in types, including:

- **Basic types**: `int`, `float64`, `string`, `bool`
- **Composite types**: Arrays, Slices, Maps, Structs, and Channels
- **Pointer types**: You can use pointers to reference a memory address in Go, but Go has no pointer arithmetic.

```go
var a int = 42
var b float64 = 3.14
var c string = "Hello, Go"
var d bool = true
```

### Functions and Methods

Functions in Go are first-class citizens and can be passed around as values.

```go
func add(a int, b int) int {
    return a + b
}
```

Methods are functions with a receiver, which allows the function to be associated with types.

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Structs

Structs are user-defined types that represent a collection of fields.

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "John", Age: 25}
    fmt.Println(p)
}
```

### Channels

Channels are Go's way to handle concurrent programming by sending and receiving data between goroutines.

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

val := <-ch
fmt.Println(val)
```

### Mutex

A `Mutex` is used to provide synchronization between goroutines to avoid race conditions.

```go
var mu sync.Mutex

mu.Lock()
// Critical section
mu.Unlock()
```

## Building RESTful APIs with Go

To build RESTful APIs in Go, you can use the `net/http` package or a framework like **Gorilla Mux**.

### CRUD Operations

Let's use **Gorilla Mux** to build a simple API with Create, Read, Update, and Delete (CRUD) operations.

#### Installation
```bash
go get -u github.com/gorilla/mux
```

#### Example API
```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Movie struct {
    ID    string `json:"id"`
    Title string `json:"title"`
    Year  string `json:"year"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(movies)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/movies", getMovies).Methods("GET")
    http.ListenAndServe(":8000", router)
}
```

### MongoDB Integration

To integrate MongoDB with Go, use the `MongoDB Go Driver`.

#### Installation
```bash
go get go.mongodb.org/mongo-driver/mongo
```

#### Connecting to MongoDB
```go
import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    return client
}
```

#### CRUD Operations with MongoDB
You can now implement CRUD operations using the MongoDB Go driver by connecting your API to MongoDB and performing operations such as inserting, reading, updating, and deleting documents from collections.

## Additional Topics

### Concurrency in Go

Go’s concurrency model is based on goroutines and channels, allowing developers to run functions independently without complicated thread management.

#### Goroutines

A goroutine is a lightweight thread managed by Go.

```go
go func() {
    fmt.Println("Hello from Goroutine!")
}()
```

#### Channels

Channels are the primary way to communicate between goroutines.

```go
ch := make(chan int)

go func() {
    ch <- 10
}()

value := <-ch
fmt.Println(value)
```

### Mutexes and Synchronization

Go provides mutexes for safely accessing shared memory from multiple goroutines.

```go
var mu sync.Mutex

func updateCount() {
    mu.Lock()
    // Critical section
    mu.Unlock()
}
```

## Conclusion

Go is a powerful language that excels in web development, systems programming, and more, thanks to its concurrency support, efficient execution, and simplicity. With tools like **Gorilla Mux** and MongoDB integration, you can build scalable and high-performance RESTful APIs with ease.