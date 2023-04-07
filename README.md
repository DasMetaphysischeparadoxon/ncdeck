# Nextcloud Deck API 

A Go package to interact with [Nextcloud Deck REST API](https://deck.readthedocs.io/en/latest/API/).

## Example

```go
package main

import (
	"fmt"

	"github.com/DasMetaphysischeparadoxon/ncdeck"
)

func main() {

	ncdeck := ncdeck.NewClient("myuser", "mypassword", "https://my.nextcloud.com")

	// TODO: more code examples

}
```

