package main

import (
	"github.com/kaddiya/docker-mysql-backup-restore/dump"
	"io/ioutil"
	"os"
	"fmt"
	"time"
	"github.com/kaddiya/docker-mysql-backup-restore/fileutils"
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

	var latestSqlDumpBasePath = fmt.Sprintf("%s/latest",os.Getenv("dump_path"))
	var archivedSqlDumpBasePath = fmt.Sprintf("%s/archived",os.Getenv("dump_path"))
	fmt.Println(latestSqlDumpBasePath)
	err1 := fileutils.CreateDirectoryIfNotExists(latestSqlDumpBasePath,0777)

	if err1 != nil{
		fmt.Println("error in checking the existence of the latest dump directory")
	}
	err2 := fileutils.CreateDirectoryIfNotExists(archivedSqlDumpBasePath,0777)
	if err2 != nil{
		fmt.Println("error in checking the existence of the archived dump directory")
	}

	t := time.Now()
  var archivedDumpFileName = fileutils.GetFullyQualifiedPathOfFile(fmt.Sprintf("%s/archived",os.Getenv("dump_path")),fmt.Sprintf("%d-%s-%d-%d:%d.sql",t.Day(),t.Month(),t.Year(),t.Hour(),t.Minute()))

	filePath := fileutils.GetFullyQualifiedPathOfFile(latestSqlDumpBasePath,"backup.sql")
	errorFilePath := fileutils.GetFullyQualifiedPathOfFile(latestSqlDumpBasePath,"error.log")

	//execute it
	errorBuf, outputBuf := dumper.MysqlDump()

	//write the latest
	ferr := ioutil.WriteFile(filePath, outputBuf.Bytes(), 0644)
	if ferr != nil {
		panic(ferr)
	}
	//write the logs
	ferr1 := ioutil.WriteFile(errorFilePath, errorBuf.Bytes(), 0644)
	if ferr1 != nil {
		panic(ferr1)
	}

	//write it to the archived folder as well
	arErr := ioutil.WriteFile(archivedDumpFileName,outputBuf.Bytes(),0644)
	if arErr !=nil {
		panic(arErr)
	}

}
