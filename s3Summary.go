package main

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/s3"
import "fmt"
import "strconv"

func main(){
    const usEast="us-east-1"
    config := aws.NewConfig().WithRegion(usEast)//ues any region to list buckets
    svc := s3.New(session.New(config))

    var params *s3.ListBucketsInput
    resp, err := svc.ListBuckets(params)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    
    for _,b :=range resp.Buckets{
        locationOutput,_:=svc.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: b.Name})

        bucketRegion:=locationOutput.LocationConstraint
        if bucketRegion!=nil{
            config.WithRegion(*bucketRegion)
        }else{
            config.WithRegion(usEast)//for 'US Standard', no such location in metadata
        }
        svc := s3.New(session.New(config)) //reset S3 sevice to use the new region
        listParams := &s3.ListObjectsV2Input{Bucket: b.Name}
        for{
            respLs, _ := svc.ListObjectsV2(listParams)
            for _,f :=range respLs.Contents{
                fmt.Printf("bucket=%s region=%s key=\"%s\" size=%d lastModified=%s\n",*b.Name,*bucketRegion,*f.Key,*f.Size,strconv.FormatInt(f.LastModified.Unix(),10))
            }
            if *respLs.IsTruncated{
                listParams.ContinuationToken=respLs.NextContinuationToken
            }else{
                break
            }
        }
    }
}
