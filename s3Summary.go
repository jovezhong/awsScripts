package main

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/s3"
import "fmt"
import "strconv"

func main(){
    config := aws.NewConfig().WithRegion("ap-southeast-1")
    svc := s3.New(session.New(config))

    var params *s3.ListBucketsInput
    resp, err := svc.ListBuckets(params)

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    
    for _,b :=range resp.Buckets{
        listParams := &s3.ListObjectsV2Input{Bucket: b.Name}
        for{
            respLs, errLs := svc.ListObjectsV2(listParams)
            if errLs != nil{
                fmt.Println(errLs.Error())
                return
            }
            for _,f :=range respLs.Contents{
                fmt.Printf("bucket=%s key=\"%s\" size=%d lastModified=%s\n",*b.Name,*f.Key,*f.Size,strconv.FormatInt(f.LastModified.Unix(),10))
            }
            if *respLs.IsTruncated{
                listParams.ContinuationToken=respLs.NextContinuationToken
            }else{
                break
            }
        }
        return
    }
}
