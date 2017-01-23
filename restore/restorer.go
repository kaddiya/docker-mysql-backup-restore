package restore


import(
  "os/exec"
  "os"
  "fmt"
  "bytes"
)

func RestoreFromFile(fileContent []byte,args []string){
  r := bytes.NewReader(fileContent)

  mysqlCmd := exec.Command("mysql",args...)
  mysqlCmd.Stdin = r
  mysqlCmd.Stderr = os.Stderr
  mysqlCmd.Stdout =  os.Stdout

  err :=  mysqlCmd.Run()

  if err !=nil{
    panic(err)
  }

}
