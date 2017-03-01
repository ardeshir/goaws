package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/aws/session"
)


func main(){

    sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})

    if err != nil {
     fmt.Println(err)
    }

    svc := s3.New(sess)

    if svc != nil {
      fmt.Println("got new s3 svc session\n")
    }

    fmt.Println("Bye!")
}
