# ptr

`ptr` is a small Go library that provides utility functions for working with pointers, including setting fallback values
if a pointer is `nil`.

## Installation

To install the package, use `go get`:

```sh
go get github.com/GiGurra/ptr
```

## Usage

### OrElse

The `OrElse` function returns the value pointed to by `ptr` if it is not `nil`, otherwise it returns the `orElse` value.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/ptr"
)

func main() {
	var str *string
	fallback := "fallback"
	value := ptr.OrElse(str, fallback)
	fmt.Println(value) // Output: fallback
}
```

### OrElseF

The `OrElseF` function is similar to `OrElse`, but it takes a function that returns the fallback value. This can be
useful if the fallback value is expensive to compute.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/ptr"
)

func main() {
	var str *string
	value := ptr.OrElseF(str, func() string { return "fallback" })
	fmt.Println(value) // Output: fallback
}
```

### ToPtr

The `ToPtr` function converts a value to a pointer. Useful for inlining values and literals where a pointer is expected.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/ptr"
)

func main() {
	ptrValue := ptr.ToPtr("hello")
	fmt.Println(*ptrValue) // Output: hello
}
```

### Map

The 'Map' function applies a function to the value pointed to by `ptr` if it is not `nil`, otherwise it returns `nil`.

```go

package main

import (
	"fmt"
	"github.com/GiGurra/ptr"
)

func main() {
	var str *string
	value := ptr.Map(str, func(s string) string { return s + " world" })
	fmt.Println(value) // Output: <nil>
}
```

### FlatMap

The 'FlatMap' function applies a function to the value pointed to by `ptr` if it is not `nil`, otherwise it returns
`nil`.

```go

package main

import (
	"fmt"
	"github.com/GiGurra/ptr"
)

func main() {
	var str *string
	value := ptr.FlatMap(str, func(s string) *string { return ptr.ToPtr(s + " world") })
	fmt.Println(*value) // Output: <nil>
}
```

## Testing

To run the tests, use the `go test` command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

*This README was AI-generated and may require further editing for accuracy and completeness.*
