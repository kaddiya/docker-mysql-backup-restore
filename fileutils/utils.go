package fileutils

import(
  "os"
  "bytes"

)
func CreateDirectoryIfNotExists(path string, mode os.FileMode)(err error){

  if _, err := os.Stat(path); os.IsNotExist(err) {
     os.Mkdir(path, mode)
  }
  return err
}

func GetFullyQualifiedPathOfFile(basePath,fileName string)string {

  var fullyQualifiedFilePathBuffer bytes.Buffer
  fullyQualifiedFilePathBuffer.WriteString(basePath)
  fullyQualifiedFilePathBuffer.WriteString("/")
  fullyQualifiedFilePathBuffer.WriteString(fileName)
  fullyQualifiedFilePath := fullyQualifiedFilePathBuffer.String()
  return fullyQualifiedFilePath
}
