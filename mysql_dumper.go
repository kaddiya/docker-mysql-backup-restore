package main

import(
  "os"
  "github.com/kaddiya/docker-mysql-backup-restore/dump"
)

func main(){
  //args := make([]string, 0, 2)
    if(os.Getenv("dumper_db_host") == "" ){
      panic("the database host is not supplied")
    }

    if(os.Getenv("dumper_db_port") == "" ){
      panic("the database port is not supplied")
    }

    if(os.Getenv("dumper_db_user") == "" ){
      panic("the database user is not supplied")
    }
    if(os.Getenv("dumper_db_password") == "" ){
        panic("the database password is not supplied")
    }
    if(os.Getenv("dumper_db_name") == "" ){
       panic("the database name is not supplied")
    }

    dumper.MysqlDump()


}
