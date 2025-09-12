package main

import (
  "fmt"
  "os/exec"
    "strings"
)
var runningProcesses []*exec.Cmd

func main() {
    fmt.Println("\nCommands\n\n1.start\n2.kill [name]\n3.exit\n4.list")
    input := ""
    for{
        fmt.Println("\nEnter commad...")
        fmt.Scan(&input)
        if strings.ToLower(input) == "exit"{
            killProgram(input)
            break
        }
        if strings.ToLower(input) == "kill"{
            fmt.Scan(&input)
            killProgram(input)
            continue
        }
        if strings.ToLower(input) == "start"{
		fmt.Println("Please, write in this format - name.exe")
        fmt.Scan(&input)
        cmd := exec.Command(input)
        cmd.Start()
		runningProcesses = append(runningProcesses, cmd)
        }
		if strings.ToLower(input) == "list" {
			listProcesses()
		}
    }
}

func killProgram(name string) error {
  cmd := exec.Command("taskkill", "/IM", name, "/F")
  err := cmd.Run()
  
  if err != nil {
    return fmt.Errorf("taskkill failed: %v", err)
  }
  fmt.Printf("Process %s killed successfully\n", name)
  return nil
}
func listProcesses() {
	fmt.Println("Processes started from this program:")
    for i, cmd := range runningProcesses {
        if cmd.Process != nil {
            fmt.Printf("%d. PID: %d\n", i+1, cmd.Process.Pid)
		}
	}
}