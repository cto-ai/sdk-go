# sdk-go

The Ops Platform SDK for Go

# Usage

## Example sdk usage

Documentation and example usage can be found in their respective source code.

```go
package main

import (
	"fmt"
	"github.com/cto-ai/sdk-go"
	"log"
	"time"
)

func main() {
	// instantiate new CTO.ai clients
	u := ctoai.NewUx()
    p := ctoai.NewPrompt()

	// printing text to interface
	err := u.Print("starting e2e test of Go SDK")
	if err != nil {
		log.Fatal(err)
	}

	// prompt user to confirm
	// captures user stdin and converts to a bool
	output, err := p.Confirm("confirmation", "confirm?", ctoai.OptConfirmFlag("C"), ctoai.OptConfirmDefault(true))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You confirmed with a %v\n", output)

	// displays a progress bar to the interface
	err = u.ProgressBarStart(5, 1, "Finishing...")
	if err != nil {
		log.Fatal(err)
	}

	// do some work
	time.Sleep(2 * time.Second)

	err = u.ProgressBarAdvance(4)
	if err != nil {
		log.Fatal(err)
	}

	err = u.ProgressBarStop("Done!")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Running inside an op

This uses the example code above.

```docker
############################
# Build container
############################
FROM golang:1.13.6 AS build

WORKDIR /ops

ADD . .

RUN go build -ldflags="-s -w" -o main

############################
# Final container
############################
FROM registry.cto.ai/official_images/base:latest

COPY --from=build /ops/main /ops/main
```

Corresponding ops.yml run command:
```yaml
run: /ops/main
```
