package dumper

import(
  "fmt"
  "os/exec"
  "bytes"
)

func MysqlDump(args []string) (stdErrBuf,stdOutBuf bytes.Buffer){

  var dumpOutBuf, dumperrBuf bytes.Buffer
  dumpArgs := make([]string, len(args)+1)
  copy(dumpArgs,args[:len(args)-1])
  dumpArgs = append(dumpArgs,"--add-drop-database")
  dumpArgs = append(dumpArgs,"--databases")
  dumpArgs = append(dumpArgs,args[len(args)-1])

  fmt.Println(dumpArgs)

  cmd := exec.Command("mysqldump", dumpArgs...)
  cmd.Stderr = &dumperrBuf
  cmd.Stdout = &dumpOutBuf
  cmd.Run()

  return dumperrBuf,dumpOutBuf
}
