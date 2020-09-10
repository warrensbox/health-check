package lib

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type TargetGroups struct {
	TargetGroup []string
}

//GetServices :  Get list of services in the cluster
func (id *Constructor) GetServices() (*TargetGroups, error) {

	svc := ecs.New(id.Session)

	var input ecs.ListServicesInput

	input.Cluster = aws.String(id.ECSCluster)

	result, err := svc.ListServices(&input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
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
		return nil, err
	}

	tgs := id.GetTargetGroupARN(result.ServiceArns)

	return tgs, nil
}

func (id *Constructor) GetTargetGroupARN(va []*string) *TargetGroups {

	svc := ecs.New(id.Session)

	var params ecs.DescribeServicesInput
	var tgs TargetGroups

	params.Cluster = aws.String(id.ECSCluster)
	params.Services = va

	resp, err := svc.DescribeServices(&params)

	if err != nil {
		// A service error occurred.
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, va := range resp.Services {
		if *&va.LoadBalancers != nil {
			for _, va := range va.LoadBalancers {
				tgs.TargetGroup = append(tgs.TargetGroup, *va.TargetGroupArn)
			}
		}
	}

	return &tgs
}
