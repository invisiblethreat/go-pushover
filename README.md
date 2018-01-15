# Pushover API Go Package

[![Documentation](https://godoc.org/github.com/invisiblethreat/go-pushover/pushover?status.svg)](https://godoc.org/github.com/invisiblethreat/go-pushover)[![Build Status](https://travis-ci.org/invisiblethreat/go-pushover.svg?branch=master)](https://travis-ci.org/invisiblethreat/go-pushover)

A Golang package for sending notifications via https://api.pushover.net.

This library is a mash-up of https://github.com/bdenning/go-pushover and
https://github.com/gregdel/pushover. Neither did exactly what I wanted, in the
way that I wanted.

## Library Example

You can use the pushover package within your Golang applications as follows:

```shell
// Set your pushover API keys. These keys are fake.
export PUSHOVER_TOKEN="KzGDOReQKggMaC0QOYAMyEEuZJnyUi"
export PUSHOVER_USER="e9e1495ec45826de5983cd1abc8031"
```

```go
package main
import (
  "fmt"

  "github.com/invisiblethreat/go-pushover/pushover"

)
// from environmental variables. Example of setting is above
config := pushover.GetConfigEnv()
// or form a YAML config file. pushover.yaml.sample is the template
config := pushover.GetConfigFile(file)

msg := pushover.NewMessageConfig(config)
res, err := m.Push("message")

if err != nil {
    fmt.Errorf("Error sending message: %s\n", err.Error())
}
```

## Dependencies

The dependency set is small- `glide` can be used to quickly resolve
all of the needed repositories with `glide install`. More information about
`glide` can be found [here](https://github.com/Masterminds/glide).

## Command Line Tool

A binary is provided by:

* running `go install github.com/invisiblethreat/go-pushover`, which installs in `$GOPATH/bin/`
* running `go build`, which builds in the root directory.

```shell

Usage of ./go-pushover:
  -f, --file string      YAML config file location (default "pushover.yaml")
  -m, --message string   Message to send. Omit if piping from STDIN
  -t, --title string     Title for message. If empty, default name for the token will be used

  ```

Then messages can be sent by piping output to the pushover command.

```shell

export PUSHOVER_TOKEN="KzGDORePKggMaC0QOYAMyEEuZJnyUi" # fake
export PUSHOVER_USER="e9e1495ec75826de5983cd1abc8031"  # fake
echo "Foo is in a critical state" | go-pushover
# or
echo "Foo is in a critical state" | go-pushover --file ../../pushover.yaml
# or
go-pushover --file ../../pushover.yaml --message "Foo is in a critical state"


```
