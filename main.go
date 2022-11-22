package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "github.com/lunux2008/xulu"
)

func readmetadata(property string) string {
  file, err := os.Open(strings.Split(os.Args[1:][0], ".")[0] + ".ssi-build")
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  var thing int
  var packagename string
  var errort = "no error"
	for scanner.Scan() {
    thing = thing + 1
		var line = strings.Split(scanner.Text(), "=")
    
    xulu.Use(line)
    if line[0] == "name" {
      if line[1] != "" {
        packagename = line[1]
        xulu.Use(packagename)
      } else {
        errort = fmt.Sprint("\033[31m", "toshscpt: err[build-syntax]: 'name' requires 1 argument but got 0 at line ", thing)
        xulu.Use(errort)
      }
    } else {
      errort = fmt.Sprint("\033[31m", "toshscpt: err[build-syntax]: invalid command at line ", thing)
      xulu.Use(errort)
    }
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
  var output string
  xulu.Use(output)
  if errort != "no error"{
    output = errort
  } else {
    if property == "packagename" {
      output = packagename
    }
  }
  return output
}
func scancode(code []string, line int) string {
  var output string
  xulu.Use(output)
  if code[0] == "printline"{  
    if code[1] != "" {
      if strings.HasPrefix(code[1], "%"){
        if code[1] == "%PackageName" { 
          output = readmetadata("packagename") 
        } else {
          output = fmt.Sprint("\033[31m", "toshscpt: err[syntax]: invalid metadata variable at line ", line)
        }
      } else {
        output = code[1]
      }
    } else {
      output = fmt.Sprint("\033[31m", "toshscpt: err[syntax]: 'printline' requires 1 argument but got 0 at line ", line)
    }
  }
  return output
}
func main() {
	file, err := os.Open(os.Args[1:][0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  var thing int
	for scanner.Scan() {
    thing = thing + 1
		fmt.Println(scancode(strings.Split(scanner.Text(), " , "), thing))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
// 
