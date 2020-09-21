package lib

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/jedib0t/go-pretty/table"
	progressbar "github.com/schollz/progressbar/v3"
)

type List struct {
	ARN    string
	Status string
}
type AtomicInt struct {
	mu sync.Mutex // A lock than can be held by one goroutine at a time.
	n  int
}

var wg = sync.WaitGroup{}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add() {
	a.mu.Lock() // Wait for the lock to be free and then take it.
	a.n++
	//fmt.Println("mutex Var:", a.n)
	a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

type tmpStruct struct {
}

func (id *Constructor) GetHealthCheck(tgs *TargetGroups) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	numOfTargets := len(tgs.TargetGroup)

	fmt.Println("(2/4) Number of target groups:", numOfTargets)

	/*Set up progress bar
	**Lets users know that the progam is still querying the target groups
	 */
	lengthOfBar := numOfTargets * id.Attempts
	bar := progressbar.NewOptions(lengthOfBar,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetRenderBlankState(false),
		progressbar.OptionSetDescription("(3/4) [cyan][Checking][reset] Target group-health ..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	var n AtomicInt //initialize mutex

	ch := make(chan *List, numOfTargets)

	wg.Add(numOfTargets)
	for i := 0; i < numOfTargets; i++ {
		go id.GetHealthStatus(tgs.TargetGroup[i], &n, bar, ch)
	}

	//wait until all target groups has been checked within n tries and after waiting for t times
	go func(ch chan<- *List) {
		defer close(ch)
		wg.Wait()
		fmt.Println("")
	}(ch)

	t.AppendHeader(table.Row{"Target Group ARN", "Status"})

	for i := range ch {
		components, _ := ParseARN(i.ARN)
		t.AppendRow([]interface{}{components.Resource, i.Status})
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	select {
	case <-ch:
		if n.Value() == numOfTargets {
			fmt.Println("(4/4) All target groups are healthy")
			os.Exit(0)
		} else {
			fmt.Println("(4/4) Not all target groups are healthy. Please log in to your AWS console to verify")
			if id.ErrorCode {
				os.Exit(1)
			}
		}
	case <-time.After(5 * time.Second):
		fmt.Println("(4/4) TIMED OUT")
		os.Exit(1)
	}

}

//GetHealthStatus get health status
func (id *Constructor) GetHealthStatus(arn string, n *AtomicInt, bar *progressbar.ProgressBar, ch chan<- *List) {

	defer wg.Done()

	var input elbv2.DescribeTargetHealthInput
	var listing List

	svc := elbv2.New(id.Session)
	listing.ARN = arn
	listing.Status = "unhealthy"
	attempt := 0
	input.TargetGroupArn = &arn

	for attempt < id.Attempts {
		result, err := svc.DescribeTargetHealth(&input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case elbv2.ErrCodeInvalidTargetException:
					fmt.Println(elbv2.ErrCodeInvalidTargetException, aerr.Error())
				case elbv2.ErrCodeTargetGroupNotFoundException:
					fmt.Println(elbv2.ErrCodeTargetGroupNotFoundException, aerr.Error())
				case elbv2.ErrCodeHealthUnavailableException:
					fmt.Println(elbv2.ErrCodeHealthUnavailableException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
					os.Exit(1)
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
				os.Exit(1)
			}
			return
		}

		for _, vl := range result.TargetHealthDescriptions {

			if *vl.TargetHealth.State == "healthy" {
				listing.Status = "healthy"
				add := id.Attempts - attempt
				bar.Add(add)
				break
			}
		}

		if listing.Status == "healthy" {
			n.Add()
			break
		} else {
			time.Sleep(time.Duration(id.Delay) * time.Second)
			attempt++
			bar.Add(1)
		}

	}

	ch <- &listing
}