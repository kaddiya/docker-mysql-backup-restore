package main

import (
	"github.com/kaddiya/mysql-backup-restore/dump"
	"github.com/kaddiya/mysql-backup-restore/models"
	"github.com/kaddiya/mysql-backup-restore/restore"
	"github.com/kaddiya/mysql-backup-restore/s3"
	"os"
	"time"
)

const (
	LATEST_DUMP_NAME = "latest-backup.sql"
	BACKUP_MODE      = "BACKUP"
	RESTORE_MODE     = "RESTORE"
)

func main() {

	if os.Getenv("dumper_mode") == "" {
		panic("please supply the mode to run.BACKUP/RESTORE")
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
		panic("Bucket name is not supplied")
	}

	if os.Getenv("dumper_s3_region") == "" {
		panic("Region is not supplied")
	}
	if os.Getenv("path_in_bucket") == "" {
		panic("path of bucket is not supplied")
	}

	client := models.InitMysqlClient(os.Getenv("dumper_db_host"),
		os.Getenv("dumper_db_port"),
		os.Getenv("dumper_db_user"),
		os.Getenv("dumper_db_password"),
		os.Getenv("dumper_db_name"))


	switch os.Getenv("dumper_mode") {

	case BACKUP_MODE:
		t := time.Now()
		latestDbBackupFileName := fmt.Sprintf("%s-%s", os.Getenv("dumper_db_name"), LATEST_DUMP_NAME)
		archiveFilename := fmt.Sprintf("%s-%d-%s-%d-%d:%d.sql", os.Getenv("dumper_db_name"), t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())

		args := models.GetCmdLineArgsFor(client)
		//execute it
		_, outputBuf := dumper.MysqlDump(args)

		s3WrappperForLatest := models.InitS3Wrapper(os.Getenv("dumper_s3_region"), os.Getenv("s3_access_key"), os.Getenv("s3_secret_key"), os.Getenv("s3_bucket_name"), os.Getenv("path_in_bucket"), fmt.Sprintf("latest/%s", latestDbBackupFileName))
		s3WrappperForBackup := models.InitS3Wrapper(os.Getenv("dumper_s3_region"), os.Getenv("s3_access_key"), os.Getenv("s3_secret_key"), os.Getenv("s3_bucket_name"), os.Getenv("path_in_bucket"), fmt.Sprintf("archived/%s", archiveFilename))
		s3.UploadFileToS3(outputBuf.Bytes(), s3WrappperForLatest)
		s3.UploadFileToS3(outputBuf.Bytes(), s3WrappperForBackup)

	case RESTORE_MODE:
		latestDbBackupFileName := fmt.Sprintf("%s-%s", os.Getenv("dumper_db_name"), LATEST_DUMP_NAME)
		s3WrappperForLatest := models.InitS3Wrapper(os.Getenv("dumper_s3_region"),
			os.Getenv("s3_access_key"),
			os.Getenv("s3_secret_key"),
			os.Getenv("s3_bucket_name"),
			os.Getenv("path_in_bucket"),
			fmt.Sprintf("latest/%s", latestDbBackupFileName))
		content, _ := s3.GetFileFromS3(s3WrappperForLatest)
		args := models.GetCmdLineArgsFor(client)
		restore.RestoreFromFile(content, args)
		fmt.Printf("restored the DB to the state as defined in the %s backup file\n", latestDbBackupFileName)
	default:
		panic("incorrect mode supplied")

	}

}
