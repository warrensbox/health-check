package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lib "github.com/warrensbox/health-checker/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	versionFlag *bool
	helpFlag    *bool
	awsRegion   *string
	ecsCluster  *string
)

var wg = sync.WaitGroup{}

type List struct {
	ServiceName string `json:"arn"`
	ARN         string `json:"arn"`
	Task        int    `json:"task"`
}

type AtomicInt struct {
	mu sync.Mutex // A lock than can be held by one goroutine at a time.
	n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add() {
	a.mu.Lock() // Wait for the lock to be free and then take it.
	a.n++
	fmt.Println("hew", a.n)
	a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func init() {

	const (
		versionFlagDesc = "Displays the version of tg-health-checker"
		awsRegionDesc   = "Provide AWS Region. Default region - us-east-1"
		ecsClusterDesc  = "ECS cluster name"
	)

	awsRegion = kingpin.Flag("region", awsRegionDesc).Short('r').String()
	ecsCluster = kingpin.Flag("ecs-cluster", ecsClusterDesc).Short('c').String()

}
func main() {

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()

	config := &aws.Config{Region: aws.String(*awsRegion)}

	session := session.Must(session.NewSession(config))

	construct := &lib.Constructor{*ecsCluster, session}

	profile := lib.NewConstructor(construct)

	fmt.Println(profile.ECSCluster)

	tgs, err := profile.GetServices()

	if err != nil {
		fmt.Println(err)
		//move to lib
	}

	numOfTargets := len(tgs.TargetGroup)

	fmt.Println("numOfTargets", numOfTargets)

	var n AtomicInt //initialize mutex

	num := 10
	ch := make(chan *List, 10)

	wg.Add(10)
	for i := 0; i < num; i++ {
		go getHealthCheck("test", &n, ch)
	}

	//wait until all target groups has been checked within n tries and after waiting for t times
	go func(ch chan<- *List) {
		defer close(ch)
		wg.Wait()
	}(ch)

	fmt.Println("val", n.Value())

	for i := range ch {
		fmt.Println(i)
	}

	select {
	case <-ch:
		if n.Value() == 10 {
			fmt.Println("All Succesfull")
		} else {
			fmt.Println("Not all services were succesfull")
		}
	case <-time.After(5 * time.Second):
		fmt.Println("TIMED OUT")
	}

}

func getHealthCheck(test string, n *AtomicInt, ch chan<- *List) {
	defer wg.Done()
	var listing List
	listing.ARN = "ABC"
	listing.Task = 1

	fmt.Println("getHealthCheck")
	time.Sleep(3 * time.Second)
	n.Add()
	ch <- &listing
}
