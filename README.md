[![Build Status](https://travis-ci.org/warrensbox/health-check.svg?branch=master)](https://travis-ci.org/warrensbox/health-check)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/health-check)](https://goreportcard.com/report/github.com/warrensbox/health-check)

# Target Group Health Checker [WIP: Experimental - Official release 09/20/20]

The `health-check` command line tool. Concurrently checks for any healthy target groups attached to a load balancer. 

## Installation

`health-check` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/).


```ruby
brew install warrensbox/tap/health-check
```

### Linux

Installation for other linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/health-check/release/install.sh | bash
```

### Docker
...

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/health-check/releases)

## How to use:


## How it works

1. This command line tool only queries all target group that is attached to a load balancer.
2. When a ecs cluster is provided, it concurrently checks for the health status for all target groups.
3. If a target groups shows at least 1 healthy task, it will return the check while other target groups health checks are concurrently going on. This way, instead of using a loop to check the health status for one target group after another. We can minimize the check times. The total wait time for the results would be the number of attempts times the delay time(in seconds).
4. The program *will not* exit with 1 unless you pass the `-e` flag for any unhealthy targets.



