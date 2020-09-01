package lib

import (
	"fmt"

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
	//var output ecs.ListServicesOutput

	input.Cluster = aws.String(id.ECSCluster)
	//input := &ecs.ListServicesInput{}

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
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	tgs := id.GetTargetGroupARN(result.ServiceArns)
	// for _, va := range result.ServiceArns {

	// 	fmt.Println(*va)
	// }

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
	}

	// Pretty-print the response data.
	//fmt.Println(awsutil.StringValue(resp))

	//fmt.Println(resp.Services)

	for _, va := range resp.Services {
		//fmt.Println(*va.ServiceName)
		//fmt.Println(*&va.LoadBalancers)
		if *&va.LoadBalancers != nil {
			//tgs.ServiceName = *va.ServiceName
			//tgs.ServiceName = *&va.LoadBalancers.TargetGroupArn
			for _, va := range va.LoadBalancers {

				//fmt.Println(*va.TargetGroupArn)
				tgs.TargetGroup = append(tgs.TargetGroup, *va.TargetGroupArn)

			}

			//[]*ecs.TargetGroupArn
		}
	}

	//fmt.Println(tgs.TargetGroup)

	return &tgs

}
