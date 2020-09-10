package lib

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
)

//Constructor : struct for session
type Constructor struct {
	ECSCluster string
	Timeout    int
	Attempts   int
	Delay      int
	Session    *session.Session
}

//NewConstructor :validate object
func NewConstructor(attr *Constructor) *Constructor {

	if attr.ECSCluster == "" {
		fmt.Println("You must provide a cluster")
		os.Exit(1)
	}
	if attr.Timeout == 0 {
		attr.Timeout = 300
	}
	if attr.Attempts == 0 {
		attr.Attempts = 5
	}
	if attr.Delay == 0 {
		attr.Delay = 10
	}
	return attr
}
