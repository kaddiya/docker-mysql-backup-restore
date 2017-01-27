package dumper

import(
  "os/exec"
  "bytes"
)

func MysqlDump(args []string) (stdErrBuf,stdOutBuf bytes.Buffer){

  var dumpOutBuf, dumperrBuf bytes.Buffer

   dumpArgs := make([]string,0,len(args)+1)
   //ensure it dumps the routines too
   dumpArgs = append(dumpArgs,"--routines")
   dumpArgs = append(dumpArgs,args...)
   cmd := exec.Command("mysqldump", dumpArgs...)

   cmd.Stderr = &dumperrBuf
   cmd.Stdout = &dumpOutBuf
   cmd.Run()

   return dumperrBuf,dumpOutBuf
}
