package main

import (
  // "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  // "github.com/aws/aws-sdk-go/aws/awserr"
  "strconv"
  // . "./sub"
  "./sub"
)


func main()  {
  ts := timeStamp.NewTimeStamp(20)

  // ts.StampWithTag("GetSession");
  sess, _ := session.NewSession(&aws.Config{
    Region: aws.String("ap-northeast-1")},
  )

  // ts.StampWithTag("InitSession");
  svc := dynamodb.New(sess)


  for i := 0; i < 16; i++ {
    ts.StampWithTag("GetItem");
    _, _  = svc.GetItem(&dynamodb.GetItemInput{
      Key: map[string]*dynamodb.AttributeValue{
        "id": {
          S: aws.String(strconv.Itoa(i)),
        },
      },
      TableName: aws.String("performance-test-dynamodb"),
    })
  }
  ts.Stamp()
  ts.Print()
  // fmt.Println(result.Item["id"])

}