package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	versionFlag *bool
	helpFlag    *bool
	awsRegion   *string
	ecsCluster  *string
	timeout     *int
	attempts    *int
	delay       *int
)

func init() {

	const (
		versionFlagDesc = "Displays the version of tg-health-checker"
		awsRegionDesc   = "Provide AWS Region. Default region - us-east-1"
		ecsClusterDesc  = "ECS cluster name"
		timeoutDesc     = "Timeout if target groups cannot be found. Default is 300 seconds"
		attemptsDesc    = "Number of attempts to query healthcheck. Default is 5 seconds"
		delayDesc       = "Delay in between healthcheck. Default is 10 seconds"
	)

	awsRegion = kingpin.Flag("region", awsRegionDesc).Short('r').String()
	ecsCluster = kingpin.Flag("ecs-cluster", ecsClusterDesc).Short('c').String()
	timeout = kingpin.Flag("timeout", timeoutDesc).Short('t').Int()
	attempts = kingpin.Flag("attempts", attemptsDesc).Short('a').Int()
	delay = kingpin.Flag("delay", delayDesc).Short('d').Int()

}
func main() {

	// kingpin.CommandLine.Interspersed(false)
	// kingpin.Parse()

	// config := &aws.Config{Region: aws.String(*awsRegion)}

	// session := session.Must(session.NewSession(config))

	// construct := &lib.Constructor{*ecsCluster, *timeout, *attempts, *delay, session}

	// profile := lib.NewConstructor(construct)

	// tgs, err := profile.GetServices()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// profile.GetHealthCheck(tgs)

	// bar := progressbar.Default(100)
	// for i := 0; i < 100; i++ {
	// 	bar.Add(1)
	// 	time.Sleep(40 * time.Millisecond)
	// }

}
