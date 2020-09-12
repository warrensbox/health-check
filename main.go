package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/warrensbox/health-checker/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	versionFlag *bool
	helpFlag    *bool
	ErrorFlag   *bool
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
		errorDesc       = "If provided, program will exit with exit status 1 if any target group is unhealthy"
	)

	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	ErrorFlag = kingpin.Flag("error", versionFlagDesc).Short('e').Bool()
	awsRegion = kingpin.Flag("region", awsRegionDesc).Short('r').String()
	ecsCluster = kingpin.Flag("ecs-cluster", ecsClusterDesc).Short('c').String()
	timeout = kingpin.Flag("timeout", timeoutDesc).Short('t').Int()
	attempts = kingpin.Flag("attempts", attemptsDesc).Short('a').Int()
	delay = kingpin.Flag("delay", delayDesc).Short('d').Int()

}
func main() {

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()

	config := &aws.Config{Region: aws.String(*awsRegion)}

	session := session.Must(session.NewSession(config))

	construct := &lib.Constructor{
		ECSCluster: *ecsCluster,
		Timeout:    *timeout,
		Attempts:   *attempts,
		Delay:      *delay,
		ErrorCode:  *ErrorFlag,
		Session:    session,
	}

	profile := lib.NewConstructor(construct)

	tgs, err := profile.GetServices()

	if err != nil {
		fmt.Println(err)
	}

	profile.GetHealthCheck(tgs)
}
