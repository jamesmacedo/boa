package src

import (
    "fmt"  
    "os"
    "os/exec"
    "regexp"
    // "strings"
)

func gitExec(args  ...string) string {
    out, err := exec.Command(GetConfigs().Bin,args...).Output()
    if err != nil {
        fmt.Println("Failed while trying to execute the command...")
        os.Exit(1)
    } 
    return string(out)
}

func getLines(syntax string,args []string) []string{ 
    r, _ := regexp.Compile(syntax)
    files := r.FindAllString(string(gitExec(args...)), -1)
    return files
}


func findCommand(commands []Command, command string) *Command{ 
    for _,c := range commands{
        if c.Command == command {
            return &c
        } 
    }
    return nil
}

// find a way to acess the list.Item variable from inside this file
func Run(command string){ 
    cmd := findCommand(GetConfigs().Commands,command)
    if cmd == nil {
        fmt.Println("Command not supported yet :/")  
        os.Exit(1)
    }

    //find modified files by regex
    // files := getLines(,[]string{"status"})
    //
    // if len(files) == 0 {
    //     fmt.Println("Nothing found :(")
    //     os.Exit(0)
    // }
    //
    // for _, file := range files {
    //     // files = append(files,{name:strings.Split(string(file),"   ")[1],path:})
    //     f := strings.Split(string(file),"   ")[1]
    //     fmt.Println(f)
    //     // items = append(items, Item{strings.Split(f,"/")[len(strings.Split(f,"/"))-1],f})
    // }
    //
    // Render(command['title'],items)
}
