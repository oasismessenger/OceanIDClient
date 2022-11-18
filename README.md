# OceanIDClient

`golang client for OceanID`

### ⚠ Please _fork_ the repository and modify the `proto_type` in the `schemes`

### quick start

#### add it to `go.mod` dependencies 

```shell
go get github.com/oasismessenger/OceanIDClient
```

#### add some code

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/oasismessenger/OceanIDClient"
)

func main() {
	const addr = "127.0.0.1:11451"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := IDClient.QuickDialWithContext(ctx, addr)
	if err != nil {
		log.Fatalln("dial fatal", err)
	}
	defer client.Close()
	
	id, err := client.GetID()
	if err != nil {
        log.Fatalln(err)
	}
	
	client.Recycle(id)
}
```
