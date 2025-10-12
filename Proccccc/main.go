package main

import (
  "fmt"
  "os/exec"
    "strings"
)

// глобальный слайс для хранения запущенных процессов
var runningProcesses []*exec.Cmd

func main() {
    fmt.Println("\nCommands\n\n1.start\n2.kill [name]\n3.exit\n4.list")
    
    input := ""
    
    // Бесконечный цикл для обработки команд пользователя
    for{
        fmt.Println("\nEnter commad...")
        fmt.Scan(&input) 
        
        // Обработка команды выхода из программы
        if strings.ToLower(input) == "exit"{
            killProgram(input)
            break 
        }
        
        // Обработка команды завершения процесса
        if strings.ToLower(input) == "kill"{
            fmt.Scan(&input) 
            killProgram(input)
            continue // Переход к следующей итерации цикла
        }
        
        // Обработка команды запуска процесса
        if strings.ToLower(input) == "start"{
            fmt.Println("Please, write in this format - name.exe")
            fmt.Scan(&input)
            
            cmd := exec.Command(input)
            cmd.Start() 
            
            // Добавление процесса в список отслеживаемых
            runningProcesses = append(runningProcesses, cmd)
        }
        
        // Обработка команды вывода списка процессов
        if strings.ToLower(input) == "list" {
            listProcesses() // Вызов функции для отображения процессов
        }
    }
}

// функция для принудительного завершения процесса по имени
func killProgram(name string) error {
  cmd := exec.Command("taskkill", "/IM", name, "/F")
  
  err := cmd.Run()
  
  // Обработка ошибки, если завершение не удалось
  if err != nil {
    return fmt.Errorf("taskkill failed: %v", err)
  }
  
  fmt.Printf("Process %s killed successfully\n", name)
  return nil
}

// функция для вывода списка запущенных процессов
func listProcesses() {
    // Заголовок списка процессов
    fmt.Println("Processes started from this program:")
    
    for i, cmd := range runningProcesses {
        if cmd.Process != nil {
            // Вывод порядкового номера и PID процесса
            fmt.Printf("%d. PID: %d\n", i+1, cmd.Process.Pid)
        }
    }
}
