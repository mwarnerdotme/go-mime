# go-mime

## Usage

It's very simple! Include it as a dependency in your project and use any of the constants or functions in the `mime` package.

### Example

```go
package main

import (
    "fmt"

    "gitlab.com/mwarnerdotme/go-mime/mime"
)

func main() {
    fmt.Println(mime.APPLICATION_JSON)
    // "application/json"
}
```
