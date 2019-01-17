package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/aws/awserr"
  // . "./sub"
  "./sub"
)

func main()  {
  ts := timeStamp.NewTimeStamp(10)

  ts.StampWithTag("GetSession");
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("ap-northeast-1")},
  )

  ts.StampWithTag("InitSession");
  svc := dynamodb.New(sess)

  ts.StampWithTag("Ready...");
  input := &dynamodb.GetItemInput{
    Key: map[string]*dynamodb.AttributeValue{
      "id": {
        S: aws.String("0"),
      },
    },
    TableName: aws.String("performance-test-dynamodb"),
  }
  ts.StampWithTag("GetItem");
  result, err := svc.GetItem(input)
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