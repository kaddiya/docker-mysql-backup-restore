package models

import(
  "fmt"
)

type MysqlClient struct{
  Host string
  Port string
  Username string
  Password string
  DatabaseName string
}


func InitMysqlClient(host,port,username,password,databaseName string)(*MysqlClient){

 return &MysqlClient{Host:host,
                    Port:port,
                    Username:username,
                    Password:password,
                    DatabaseName:databaseName}

}

func GetCmdLineArgsFor(c *MysqlClient)([]string){
  args := make([]string, 0, 16)
  args = append(args, fmt.Sprintf("-h%s",c.Host))
  args = append(args, fmt.Sprintf("-P%s",c.Port))
  args = append(args, fmt.Sprintf("-u%s",c.Username))
  args = append(args, fmt.Sprintf("-p%s",c.Password))
  args = append(args, fmt.Sprintf("%s",c.DatabaseName))
  return args
}
