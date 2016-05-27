package main

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/s3"
import "fmt"

func main(){
    config := aws.NewConfig().WithRegion("us-west-2")
    svc := s3.New(session.New(config))

    var params *s3.ListBucketsInput
    resp, err := svc.ListBuckets(params)

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    
    for i,b :=range resp.Buckets{
        fmt.Println(i)
        fmt.Println(b)
    }

    // Pretty-print the response data.
    fmt.Println(resp.Buckets)    
}
