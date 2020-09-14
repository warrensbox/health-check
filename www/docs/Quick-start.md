## How to use
### Pass your ecs-cluster name
<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-all-good.gif" alt="drawing"  height="300"/>

1. You must always provide the `ecs-cluster` on the command-line
2. Optionally, you can provide the `delay` option to delay in-between checks
3. Optionally, you can also provide the `attempts` option for the number of attempts for the health check


### Use the `-e` option
<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-all-bad-1.gif" alt="drawing"  height="300"/>

1. You you provide the `-e` flag, the program  will exit with `exit code 1` if any of the target is unhealthy
2. This is useful for continuous delivery - Jenkins, CircleCI and others  

<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-all-bad-0.gif" alt="drawing"  height="300"/>

1. Likewise, if you dont't provide the `e` flag, the program  will simply exit with `exit code 0` if any of the target is unhealthy

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