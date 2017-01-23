package dumper

import(
  "os/exec"
  "bytes"
)

func MysqlDump(args []string) (stdErrBuf,stdOutBuf bytes.Buffer){

  var dumpOutBuf, dumperrBuf bytes.Buffer
   cmd := exec.Command("mysqldump", args...)
   cmd.Stderr = &dumperrBuf
   cmd.Stdout = &dumpOutBuf
   cmd.Run()

   return dumperrBuf,dumpOutBuf
}
