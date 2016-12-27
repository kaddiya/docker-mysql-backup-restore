package restore


import(
  "os/exec"
  "fmt"
  "os"
  "io/ioutil"
  "bytes"
)

func RestoreFromFile(pathToFile string){

  b, err := ioutil.ReadFile(pathToFile)
  r := bytes.NewReader(b)

  if err !=nil {
    panic(err)
  }

  args := make([]string, 0, 16)
  args = append(args, fmt.Sprintf("-h%s",os.Getenv("dumper_db_host")))
  args = append(args, fmt.Sprintf("-P%s",os.Getenv("dumper_db_port")))
  args = append(args, fmt.Sprintf("-u%s",os.Getenv("dumper_db_user")))
  args = append(args, fmt.Sprintf("-p%s",os.Getenv("dumper_db_password")))
  args = append(args, fmt.Sprintf("%s",os.Getenv("dumper_db_name")))

  mysqlCmd := exec.Command("mysql",args...)
  mysqlCmd.Stdin = r
  mysqlCmd.Stderr = os.Stderr
  mysqlCmd.Stdout =  os.Stdout

  err2 :=  mysqlCmd.Run()

  if err2 !=nil{
    panic(err2)
  }


}
