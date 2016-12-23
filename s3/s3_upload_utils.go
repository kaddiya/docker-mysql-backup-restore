package s3

import(
  "fmt"
  "bytes"
  "os"
  "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/aws/awsutil"
 "github.com/aws/aws-sdk-go/aws/credentials"
 "github.com/aws/aws-sdk-go/service/s3"
 "github.com/aws/aws-sdk-go/aws/session"
)

func UploadFileToS3(content []byte){

    aws_access_key_id := os.Getenv("s3_access_key")
    aws_secret_access_key := os.Getenv("s3_secret_key")
    token := ""
    creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

    _, err := creds.Get()
    if err != nil {
      fmt.Printf("bad credentials: %s", err)
    }
    cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)

    svc := s3.New(session.New(), cfg)

    fileType := ".sql"
    path := "/dumps/latest/latest.sql"
    fileBytes := bytes.NewReader(content)
    params := &s3.PutObjectInput{
      Bucket: aws.String(os.Getenv("s3_bucket_name")),
      Key: aws.String(path),
      Body: fileBytes,
      ContentType: aws.String(fileType),
  }

  resp, err := svc.PutObject(params)
  if err != nil {
    fmt.Printf("bad response: %s", err)
  }
  fmt.Printf("response %s", awsutil.StringValue(resp))
}
