package s3

import(
  "fmt"
  "bytes"
  "os"
  "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/aws/credentials"
 "github.com/aws/aws-sdk-go/service/s3"
 "github.com/aws/aws-sdk-go/aws/session"
)

func UploadFileToS3(content []byte,path,name string){

    aws_access_key_id := os.Getenv("s3_access_key")
    aws_secret_access_key := os.Getenv("s3_secret_key")
    token := ""
    creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

    _, err := creds.Get()
    if err != nil {
      fmt.Printf("bad credentials: %s", err)
    }

    cfg := aws.NewConfig().WithRegion(os.Getenv("dumper_s3_region")).WithCredentials(creds)

    svc := s3.New(session.New(), cfg)

  fileBytes := bytes.NewReader(content)
    params := &s3.PutObjectInput{
      Bucket: aws.String(os.Getenv("s3_bucket_name")),
      Key: aws.String(fmt.Sprintf("%s/%s",path,name)),
      Body: fileBytes,
  }

  _, err1 := svc.PutObject(params)

  if err1 != nil {
    fmt.Printf("bad response: %s", err1)
    return
  }
  fmt.Println(fmt.Sprintf("file %s succefully uploaded to %s",name,path))

}
