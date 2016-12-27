package models

type S3Wrapper struct{
  Region string
  AccessKey string
  Secretkey string
  BucketName string
  PathInBucket string
  KeyName string
}


func InitS3Wrapper(region,accessKey,secretKey,bucketName,pathInBucket,keyName string)(*S3Wrapper){

 return &S3Wrapper{Region:region,
                    AccessKey:accessKey,
                    Secretkey:secretKey,
                    BucketName:bucketName,
                    PathInBucket:pathInBucket,
                    KeyName:keyName}

}
