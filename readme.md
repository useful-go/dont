# Don't
There are plenty of libraries already doing stuff, but we sometimes don't want to avoid doing something.

## Installing

```
go get github.com/useful-go/dont
```

## Test Coverage
100%

## Usage

If you don't want to do something:
```go
import "github.com/useful-go/dont"
import "log"

func main() {
	// this will never execute
	err := dont.Do(func() {
		log.Println("This")
	})
	// check if it really didn't execute
	if err != nil {
		log.Fatal("Yayx, it failed", err)
	}
}
```

## Use Cases

**Not Todo List**
- list of things you never want to do

```go
import "github.com/useful-go/dont"

func main() {
	
}
```