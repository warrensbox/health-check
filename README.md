[![Build Status](https://travis-ci.org/warrensbox/health-check.svg?branch=master)](https://travis-ci.org/warrensbox/health-check)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/health-check)](https://goreportcard.com/report/github.com/warrensbox/health-check)
[![CircleCI](https://circleci.com/gh/warrensbox/health-check/tree/master.svg?style=shield&circle-token=c5d416ceb68675bb6602c58b084a2df2d51d7601)](https://circleci.com/gh/warrensbox/health-check)



# AWS Target Group Health Checker 

<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/logo.svg" alt="drawing" width="100" height="130"/>

The `health-check` command-line tool concurrently checks all target groups' health status (for target groups that are attached to a load balancer).  

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
docker pull warrensbox/health-check:latest
```
or  
```sh
docker pull docker.pkg.github.com/warrensbox/health-check/health-check:latest
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/health-check/releases)

## How to use
### Pass your ecs-cluster name
<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-all-good.gif" alt="drawing"  height="300"/>

1. You must always provide the `ecs-cluster` on the command-line
2. Optionally, you can provide the `delay` option to delay in-between checks
3. Optionally, you can also provide the `attempts` option for the number of attempts for the health check


### Use the `-e` option
<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-all-bad-1.gif" alt="drawing"  height="300"/>

1. You you provide the `e` flag, the program  will exit with `exit code 1` if any of the target is unhealthy
2. This is useful for continuous delivery - Jenkins, CircleCI and others  

<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-all-bad-0.gif" alt="drawing"  height="300"/>

1. Likewise, if you dont't provide the `-e` flag, the program  will simply exit with `exit code 0` if any of the target is unhealthy

### With Docker
```sh
docker run --rm \
  -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
  -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
  -e AWS_SESSION_TOKEN=${AWS_SESSION_TOKEN} \
  -e AWS_REGION=${AWS_REGION} \
  -e AWS_DEFAULT_REGION=${AWS_REGION} \
  health-check \
  --ecs-cluster esp-devops \ #cluster name 
  --attempts 50 \  #number of attempts
  --delay 2 #number of attempts
```

## How it works

<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-diagram.png" alt="drawing"  height="300"/>


1. This command line tool only queries all target group that is attached to a load balancer.
2. Given an ecs cluster is provided, it concurrently checks for the health status for all target groups in that cluster.
3. If a target groups shows *at least* 1 healthy task, it will return the check while other target groups health checks are concurrently going on.  
4. This way, instead of using a loop to check each target group health status - one after another, we can minimize the time by checking all the target groups' health concurrently. The total wait time for the results would be the `number of attempts` X `delay time`*(in seconds)*.
5. The program **will not** exit with *error code 1* unless you pass the `-e` flag for any unhealthy targets.


## Want new feature? Want to contibute?

Please open  *issues* here: [New Issue](https://github.com/warrensbox/health-check/issues)  

Or open a Pull Request.  


