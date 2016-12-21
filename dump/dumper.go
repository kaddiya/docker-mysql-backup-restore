package dumper

import(
  "fmt"
  "os/exec"
  "os"
  "bytes"
)

func MysqlDump(){

  mysqldumpPathcmdArgs:= make([]string, 0, 2)
  mysqldumpPathcmdArgs = append(mysqldumpPathcmdArgs,"mysqldump")

  var outbuf, errbuf bytes.Buffer
  mysqldumpPathcmd := exec.Command("which",mysqldumpPathcmdArgs...)
  mysqldumpPathcmd.Stderr = &errbuf
  mysqldumpPathcmd.Stdout = &outbuf
  mysqldumpPathcmd.Run()

  if &errbuf == nil{
    panic("could not find the installation for mysqldump")
  }

  args := make([]string, 0, 16)
  args = append(args, fmt.Sprintf("--host=%s",os.Getenv("dumper_db_host")))
  args = append(args, fmt.Sprintf("--port=%s",os.Getenv("dumper_db_port")))
  args = append(args, fmt.Sprintf("--user=%s",os.Getenv("dumper_db_user")))
  args = append(args, fmt.Sprintf("--password=%s",os.Getenv("dumper_db_password")))
  args = append(args, fmt.Sprintf("%s",os.Getenv("dumper_db_name")))

  fmt.Println(os.Getenv("dumper_db_name"))
  cmd := exec.Command(outbuf.String(), args...)
  cmd.Stderr = os.Stderr
  cmd.Stdout = os.Stdout
  cmd.Run()
  


}
