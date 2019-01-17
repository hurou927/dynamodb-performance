package main

import (
  // "fmt"
  "github.com/aws/aws-sdk-go/aws"
  // "github.com/aws/aws-sdk-go-v2/aws/defaults"
  "github.com/aws/aws-sdk-go-v2/aws/endpoints"
  "github.com/aws/aws-sdk-go-v2/aws/external"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb"
  // "github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
  // "github.com/aws/aws-sdk-go-v2/aws/awserr"
  // . "./sub"
  "./sub"
  "strconv"
)
type Record struct {
  id   string
  // URLs   []string
}
func main()  {
  ts := timeStamp.NewTimeStamp(20)

  ts.StampWithTag("GetSession");
  cfg, err := external.LoadDefaultAWSConfig()
  if err != nil {
    panic("unable to load SDK config, " + err.Error())
  }

  cfg.Region = endpoints.ApNortheast1RegionID

  svc := dynamodb.New(cfg)

  for i := 0; i < 16; i++ {
    ts.StampWithTag("GetItem");
    _, _  = svc.GetItemRequest(&dynamodb.GetItemInput{
      Key: map[string]dynamodb.AttributeValue{
        "id": {
          S: aws.String(strconv.Itoa(i)),
        },
      },
      TableName: aws.String("performance-test-dynamodb"),
    }).Send()
  }
  ts.Stamp()
  ts.Print()
  // fmt.Println(result)

}