[![Build Status](https://travis-ci.org/warrensbox/health-check.svg?branch=master)](https://travis-ci.org/warrensbox/health-check)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/health-check)](https://goreportcard.com/report/github.com/warrensbox/health-check)

# Target Group Health Checker [WIP: Experimental - Official release 09/20/20]

The `health-check` command line tool concurrently checks all target groups's health status. Only checks for target groups that are attached to a load balancer.  

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
```sh
docker run --rm \
  -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
  -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
  -e AWS_SESSION_TOKEN=${AWS_SESSION_TOKEN} \
  -e AWS_REGION=${AWS_REGION} \
  -e AWS_DEFAULT_REGION=${AWS_REGION} \
  health-check \
  -c esp-devops
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/health-check/releases)

## How to use:


## How it works

1. This command line tool only queries all target group that is attached to a load balancer.
2. Given an ecs cluster is provided, it concurrently checks for the health status for all target groups in that cluster.
3. If a target groups shows *at least* 1 healthy task, it will return the check while other target groups health checks are concurrently going on.  
4. This way, instead of using a loop to check each target group health status - one after another, we can minimize the time by checking all the target groups' health concurrently. The total wait time for the results would be the `number of attempts` X `delay time`*(in seconds)*.
5. The program **will not** exit with *error code 1* unless you pass the `-e` flag for any unhealthy targets.



