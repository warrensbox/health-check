## How it works

<img style="text-allign:center" src="https://kepler-images.s3.us-east-2.amazonaws.com/warrensbox/health-check/health-check-diagram.png" alt="drawing"  height="300"/>


1. This command line tool only queries all target group that is attached to a load balancer.
2. Given an ecs cluster is provided, it concurrently checks for the health status for all target groups in that cluster.
3. If a target groups shows *at least* 1 healthy task, it will return the check while other target groups health checks are concurrently going on.  
4. This way, instead of using a loop to check each target group health status - one after another, we can minimize the time by checking all the target groups' health concurrently. The total wait time for the results would be the `number of attempts` X `delay time`*(in seconds)*.
5. The program **will not** exit with *error code 1* unless you pass the `-e` flag for any unhealthy targets.
