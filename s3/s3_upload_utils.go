package s3

import(
  "fmt"
  "bytes"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/aws/credentials"
 "github.com/aws/aws-sdk-go/service/s3"
 "github.com/aws/aws-sdk-go/aws/session"
 "io/ioutil"
 "github.com/kaddiya/mysql-backup-restore/models"
)

func UploadFileToS3(content []byte,wrapper *models.S3Wrapper){


    creds := credentials.NewStaticCredentials(wrapper.AccessKey, wrapper.Secretkey, "")

    _, err := creds.Get()
    if err != nil {
      fmt.Printf("bad credentials: %s", err)
    }

    cfg := aws.NewConfig().WithRegion(wrapper.Region).WithCredentials(creds)

  svc := s3.New(session.New(), cfg)

  fileBytes := bytes.NewReader(content)
    params := &s3.PutObjectInput{
      Bucket: aws.String(wrapper.BucketName),
      Key: aws.String(getFilePathInS3(wrapper)),
      Body: fileBytes,
  }

  _, err1 := svc.PutObject(params)

  if err1 != nil {
    panic(err1)
  }
  fmt.Printf("File succesfully uploaded to bucket %s.File is found at %s with name %s\n",wrapper.BucketName,wrapper.PathInBucket,wrapper.KeyName)
}

func GetFileFromS3(wrapper *models.S3Wrapper)(content []byte,err error){

  creds := credentials.NewStaticCredentials(wrapper.AccessKey, wrapper.Secretkey,"")

  _, errz := creds.Get()
  if errz != nil {
    fmt.Printf("bad credentials: %s", errz)
  }

  cfg := aws.NewConfig().WithRegion(wrapper.Region).WithCredentials(creds)

  svc := s3.New(session.New(), cfg)

  params := &s3.GetObjectInput{
      Bucket:aws.String(wrapper.BucketName),
      Key:aws.String(getFilePathInS3(wrapper)),
  }
  if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    panic(err)
}

  resp, err2 := svc.GetObject(params)
  if err2 !=nil {
    panic(err2)
  }

  return ioutil.ReadAll(resp.Body)

}

func getFilePathInS3(s3Wrapper *models.S3Wrapper) string{
  return fmt.Sprintf("%s/%s",s3Wrapper.PathInBucket,s3Wrapper.KeyName)
}
