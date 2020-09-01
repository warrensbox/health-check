package lib

import "github.com/aws/aws-sdk-go/aws/session"

//Constructor : struct for session
type Constructor struct {
	ECSCluster string
	Session    *session.Session
}

//NewConstructor :validate object
func NewConstructor(attr *Constructor) *Constructor {

	if attr.ECSCluster == "" {

	}
	return attr
}
