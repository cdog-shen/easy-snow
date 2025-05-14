# easy-snow

## Overview

A Simple SnowFlake ID Generator

## Features

- Generates unique snowflake IDs.
- Supports both UUID-based and non-UUID-based NodeIDs.
- High performance with bitwise operations for ID construction.
- Includes performance tests for ID generation.

## Installation

To install the package, run:

```bash
go get github.com/cdog-shen/easy-snow
```

## Usage

Here's a basic example of how to use the Easy Snowflake ID Generator:

```go
package main

import (
    "fmt"
    "github.com/cdog-shen/easy-snow"
)

func main() {
    // Initialize the snowflake generator with a node ID
    err := easysnow.SetNodeID("your-node-UUID", true)
    if err != nil {
        fmt.Println("Error setting NodeID:", err)
        return
    }

    // Generate a unique snowflake ID
    snowflakeID, err := easysnow.GenerateSnowflakeID()
    if err != nil {
        fmt.Println("Error generating snowflake ID:", err)
        return
    }

    fmt.Println("Generated Snowflake ID:", snowflakeID)
}
```

## Testing

The package includes performance tests for ID generation. To run the tests, use:

```bash
go test -v ./...
```

Result:

```txt
?       github.com/cdog-shen/easy-snow  [no test files]
?       github.com/cdog-shen/easy-snow/internal [no test files]
=== RUN   TestGenerateSnowflakeWithUUIDPerformance
    performance_test.go:31: Generated 500000 IDs in 55.1966ms
--- PASS: TestGenerateSnowflakeWithUUIDPerformance (0.06s)
=== RUN   TestGenerateSnowflakeIDPerformance
    performance_test.go:54: Generated 500000 IDs in 53.7617ms
--- PASS: TestGenerateSnowflakeIDPerformance (0.05s)
PASS
ok      github.com/cdog-shen/easy-snow/test     0.520s
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any bugs or feature requests.
