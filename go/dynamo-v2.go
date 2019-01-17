package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    // "github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	// . "./sub"
	"./sub"
)
type Record struct {
    id     string
    // URLs   []string
}
func main()  {
	ts := timeStamp.NewTimeStamp(10)

	ts.StampWithTag("GetSession");
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Set the AWS Region that the service clients should use
	cfg.Region = endpoints.ApNortheast1RegionID

	ts.StampWithTag("InitSession");
	svc := dynamodb.New(cfg)

	ts.StampWithTag("Ready...");
	// r := Record{
    //     id:   "66a36e77-fd00-3779-8097-17841f998f4d",
    // }
	// item, err := dynamodbattribute.MarshalMap(r)
	input := &dynamodb.GetItemInput{
    	Key: map[string]dynamodb.AttributeValue{
	        "id": {
	            S: aws.String("0"),
	        },
	    },
	    TableName: aws.String("performance-test-dynamodb"),
	}
	ts.StampWithTag("GetItem");
	result, err := svc.GetItemRequest(input).Send();
		if err != nil {
    		if aerr, ok := err.(awserr.Error); ok {
    		    switch aerr.Code() {
    		    case dynamodb.ErrCodeProvisionedThroughputExceededException:
    		        fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
    		    case dynamodb.ErrCodeResourceNotFoundException:
    		        fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
    		    // case dynamodb.ErrCodeRequestLimitExceeded:
    		        // fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
    		    case dynamodb.ErrCodeInternalServerError:
    		        fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
    		    default:
    		        fmt.Println(aerr.Error())
    	    }
    	} else {
    	    // Print the error, cast err to awserr.Error to get the Code and
    	    // Message from an error.
    	    fmt.Println(err.Error())
    	}
    	return
	}	
	ts.Stamp()
    ts.Print()
	fmt.Println(result)

}