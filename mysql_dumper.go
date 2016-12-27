package main

import (
	"fmt"
	"github.com/kaddiya/docker-mysql-backup-restore/dump"
	"github.com/kaddiya/docker-mysql-backup-restore/restore"
	"github.com/kaddiya/docker-mysql-backup-restore/fileutils"
	"github.com/kaddiya/docker-mysql-backup-restore/s3"
	"os"
	"time"
	"github.com/kaddiya/docker-mysql-backup-restore/models"
)

const (
	LATEST_DUMP_NAME = "latest-backup.sql"
	BACKUP_MODE = "BACKUP"
	RESTORE_MODE = "RESTORE"
)

func main() {

	if os.Getenv("dumper_mode") ==""{
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

	var latestSqlDumpBasePath = fmt.Sprintf("%s/latest", "/backups")
	var archivedSqlDumpBasePath = fmt.Sprintf("%s/archived", "/backups")

	err1 := fileutils.CreateDirectoryIfNotExists(latestSqlDumpBasePath, 0777)

	if err1 != nil {
		fmt.Println("error in checking the existence of the latest dump directory")
	}

	err2 := fileutils.CreateDirectoryIfNotExists(archivedSqlDumpBasePath, 0777)
	if err2 != nil {
		fmt.Println("error in checking the existence of the archived dump directory")
	}

	switch os.Getenv("dumper_mode") {

		case BACKUP_MODE:
			t := time.Now()
			latestDbBackupFileName := fmt.Sprintf("%s-%s", os.Getenv("dumper_db_name"), LATEST_DUMP_NAME)
			archiveFilename := fmt.Sprintf("%s-%d-%s-%d-%d:%d.sql",os.Getenv("dumper_db_name"), t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
			archivedDumpFileName := fileutils.GetFullyQualifiedPathOfFile(archivedSqlDumpBasePath, archiveFilename)
			latestDumpFilePath := fileutils.GetFullyQualifiedPathOfFile(latestSqlDumpBasePath, latestDbBackupFileName)
			errorFilePath := fileutils.GetFullyQualifiedPathOfFile(latestSqlDumpBasePath, "error.log")

			args:= models.GetCmdLineArgsFor(client)
			//execute it
			errorBuf, outputBuf := dumper.MysqlDump(args)


			//write the latest
			fileutils.WriteToFile(latestDumpFilePath, outputBuf.Bytes())
			//write error log
			fileutils.WriteToFile(errorFilePath, errorBuf.Bytes())
			//write the archive itself
			fileutils.WriteToFile(archivedDumpFileName, outputBuf.Bytes())

			s3.UploadFileToS3(outputBuf.Bytes(), "/db-backups/latest", latestDbBackupFileName)
			s3.UploadFileToS3(outputBuf.Bytes(), "/db-backups/archived", archiveFilename)

 		case RESTORE_MODE:
			latestDbBackupFileName := fmt.Sprintf("%s-%s", os.Getenv("dumper_db_name"), LATEST_DUMP_NAME)
			content,_:= s3.GetFileFromS3("/db-backups/latest", latestDbBackupFileName)
			args:= models.GetCmdLineArgsFor(client)
			restore.RestoreFromFile(content,args)
			fmt.Printf("restored the DB to the state as defined in the %s backup file",latestDbBackupFileName)
		default:
			panic("incorrect mode supplied")

	}

}
