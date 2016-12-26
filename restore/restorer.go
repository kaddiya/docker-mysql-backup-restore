package restore


import(
  "os/exec"
  "fmt"
  "os"
)
func RestoreFromFile(pathToFile string){

  args := make([]string, 0, 16)
  args = append(args, fmt.Sprintf("--host=%s",os.Getenv("dumper_db_host")))
  args = append(args, fmt.Sprintf("--port=%s",os.Getenv("dumper_db_port")))
  args = append(args, fmt.Sprintf("--user=%s",os.Getenv("dumper_db_user")))
  args = append(args, fmt.Sprintf("--password=",os.Getenv("dumper_db_password")))
  args = append(args, fmt.Sprintf("--user=%s",os.Getenv("dumper_db_user")))
  args = append(args, fmt.Sprintf("%s < %s",os.Getenv("dumper_db_name"),pathToFile))

  cmd := exec.Command("mysql", args...)
  cmd.Stderr = os.Stderr
  cmd.Stdout = os.Stdout
  cmd.Run()

}
