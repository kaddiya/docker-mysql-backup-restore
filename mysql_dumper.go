package main

import (
	"fmt"
	"github.com/kaddiya/docker-mysql-backup-restore/dump"
	"github.com/kaddiya/docker-mysql-backup-restore/fileutils"
	"github.com/kaddiya/docker-mysql-backup-restore/s3"
	"os"
	"time"
)

func main() {

	if os.Getenv("dump_path") == "" {
		panic("the path for the backups")
	}
	if os.Getenv("dumper_db_host") == "" {
		panic("the database host is not supplied")
	}

	if os.Getenv("dumper_db_port") == "" {
		panic("the database port is not supplied")
	}

	if os.Getenv("dumper_db_user") == "" {
		panic("the database user is not supplied")
	}
	if os.Getenv("dumper_db_password") == "" {
		panic("the database password is not supplied")
	}
	if os.Getenv("dumper_db_name") == "" {
		panic("the database name is not supplied")
	}

	if os.Getenv("s3_access_key") == "" {
		panic("the s3 access key is not supplied")
	}

	if os.Getenv("s3_secret_key") == "" {
		panic("the s3 secret key is not supplied")
	}

	if os.Getenv("s3_bucket_name") == "" {
		panic("Please supply the bucket name")
	}

	var latestSqlDumpBasePath = fmt.Sprintf("%s/latest", os.Getenv("dump_path"))
	var archivedSqlDumpBasePath = fmt.Sprintf("%s/archived", os.Getenv("dump_path"))

	err1 := fileutils.CreateDirectoryIfNotExists(latestSqlDumpBasePath, 0777)

	if err1 != nil {
		fmt.Println("error in checking the existence of the latest dump directory")
	}

	err2 := fileutils.CreateDirectoryIfNotExists(archivedSqlDumpBasePath, 0777)
	if err2 != nil {
		fmt.Println("error in checking the existence of the archived dump directory")
	}

	t := time.Now()
	archiveFilename := fmt.Sprintf("%d-%s-%d-%d:%d.sql", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	archivedDumpFileName := fileutils.GetFullyQualifiedPathOfFile(archivedSqlDumpBasePath, archiveFilename)
	latestDumpFilePath := fileutils.GetFullyQualifiedPathOfFile(latestSqlDumpBasePath, "backup.sql")
	errorFilePath := fileutils.GetFullyQualifiedPathOfFile(latestSqlDumpBasePath, "error.log")

	//execute it
	errorBuf, outputBuf := dumper.MysqlDump()

	//write the latest
	fileutils.WriteToFile(latestDumpFilePath, outputBuf.Bytes())
	//write error log
	fileutils.WriteToFile(errorFilePath, errorBuf.Bytes())
	//write the archive itself
	fileutils.WriteToFile(archivedDumpFileName, outputBuf.Bytes())

	s3.UploadFileToS3(outputBuf.Bytes())

}
