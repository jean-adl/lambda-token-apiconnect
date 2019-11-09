# Lambda Token API Connect

Lambda token api connect is a golang lambda for getting token from api connect.

## Installation

From source code you must run the follow command for create a executable program
```bash
GOOS=linux go build -o main
```

## Libraries and SDKs
The official Amazon Web Service SDK
```go
package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
        return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
        lambda.Start(HandleRequest)
}
```

## Autor
Jean Carlos Rodríguez 

## License
Aval Digital Labs