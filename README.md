# hcloud: A Go library for the Hetzner Cloud API

[![GitHub Actions status](https://github.com/hetznercloud/hcloud-go/workflows/Continuous%20Integration/badge.svg)](https://github.com/hetznercloud/hcloud-go/actions)
[![GoDoc](https://godoc.org/github.com/hetznercloud/hcloud-go/hcloud?status.svg)](https://godoc.org/github.com/hetznercloud/hcloud-go/hcloud)

Package rewardcloud is a library for the Reward Cloud API.

The library’s documentation is available at [GoDoc](https://godoc.org/github.com/hetznercloud/hcloud-go/hcloud),
the public API documentation is available at [docs.hetzner.cloud](https://docs.hetzner.cloud/).

## Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rewardenv/rewardcloud-go/rewardcloud"
)

func main() {
    client := rewardcloud.NewClient(rewardcloud.WithToken("token"))

    project, _, err := client.Project.GetByID(context.Background(), 1)
    if err != nil {
        log.Fatalf("error retrieving project: %s\n", err)
    }
    if project != nil {
        fmt.Printf("project 1 is called %q\n", project.Name)
    } else {
        fmt.Println("server 1 not found")
    }
}
```

## License

This work is licensed under the MIT license. See LICENSE file for details.
