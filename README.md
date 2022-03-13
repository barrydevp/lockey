# lockey [![Go Reference](https://pkg.go.dev/badge/github.com/barrdevp/lockey)](https://pkg.go.dev/github.com/barrydevp/lockey)

Golang mutex lock by key(string). Easy to use, friendly API, use when you want to lock something by key.

Easy and Fast, support RWMutex.

## Installation

```
go get github.com/barrydevp/lockey

```

## Usage

```go
// Import lockey into your code and refer it as `lockey`.
import "github.com/barrydevp/lockey"

```

### Simple Lock/Unlock and RLock/RUnlock
```go
// Create lockey instance
locker := lockey.NewRWLockKey()

// Lock/Unlock usage

// Lock
locker.Lock("your key") 

// ...

// Unlock
locker.Unlock("your key")

// RLock/RUnlock usage

// Lock
locker.RLock("your key") 

// ...

// Unlock
locker.RUnlock("your key")

```


## Testing and Benmarking

```
=== RUN   TestKeyMutex
--- PASS: TestKeyMutex (0.02s)
=== RUN   TestLockeyWithLock
--- PASS: TestLockeyWithLock (2.00s)
=== RUN   TestNewLockeyWithoutLock
--- PASS: TestNewLockeyWithoutLock (1.00s)
goos: linux
goarch: amd64
pkg: github.com/barrydevp/lockey
cpu: Intel(R) Core(TM) i3-4010U CPU @ 1.70GHz
BenchmarkLockey
BenchmarkLockey-4           2088            526444 ns/op
PASS
ok      github.com/barrydevp/lockey     4.207s
```

