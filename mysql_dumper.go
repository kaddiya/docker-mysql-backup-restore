package main

import(
  "os/exec"
  "os"  
)

func main(){

    args := make([]string, 0, 16)
    args = append(args, "--host=127.0.0.1")
    args = append(args, "--port=3306")
    args = append(args, "--user=deploy")
    args = append(args, "--password=deploy")
    args = append(args, "proof")


    cmd := exec.Command("/usr/local/bin/mysqldump", args...)
    cmd.Stderr = os.Stderr
	  cmd.Stdout = os.Stdout
    cmd.Run()
}
