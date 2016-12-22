package main

import (
	"github.com/kaddiya/docker-mysql-backup-restore/dump"
	"io/ioutil"
	"os"
	//"fmt"
	"bytes"
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
	//create the file path for the dump
	var outputFilePathNameBuffer bytes.Buffer
	outputFilePathNameBuffer.WriteString(os.Getenv("dump_path"))
	outputFilePathNameBuffer.WriteString("/")
	outputFilePathNameBuffer.WriteString("latestbackup.sql")

	//create an error log for the dump
	var errorFilePathNameBuffer bytes.Buffer
	errorFilePathNameBuffer.WriteString(os.Getenv("dump_path"))
	errorFilePathNameBuffer.WriteString("/")
	errorFilePathNameBuffer.WriteString("error.log")

	//get the fully qualified path names
	filePath := outputFilePathNameBuffer.String()
	errorFilePath := errorFilePathNameBuffer.String()

	//execute it
	errorBuf, outputBuf := dumper.MysqlDump()

	//write it
	ferr := ioutil.WriteFile(filePath, outputBuf.Bytes(), 0644)
	if ferr != nil {
		panic(ferr)
	}

	ferr1 := ioutil.WriteFile(errorFilePath, errorBuf.Bytes(), 0644)
	if ferr1 != nil {
		panic(ferr1)
	}

}
