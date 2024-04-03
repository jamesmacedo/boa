package src

import (
    "gopkg.in/yaml.v3"
    "fmt"
    "os"
    "bufio"
    "sync"
)

type Command struct {
   Command string `yaml:"command"`
   Confirm bool `yaml: confirm`
   Pattern string `yaml:"pattern"`
} 

type Config struct {
    Bin string `yaml:"bin"`
    Commands []Command
}

var lock = &sync.Mutex{}

var configs *Config

/* 
    Reads the config file, the file in question is '~/boa.yaml', located in user folder 
*/
func GetConfigs() *Config{
    
    if configs != nil {
        return configs
    }

    lock.Lock()
    defer lock.Unlock()

    file, err := os.OpenFile("/home/nemo/boa.yaml",os.O_RDWR,0644)

    if err != nil{
        fmt.Println("Configuration file not found :(")
        fmt.Printf("err: %v",err)
        os.Exit(1)
    }   

    stat, err := file.Stat()

    if err != nil {
        fmt.Printf("err: %v",err)
        os.Exit(1)
    }

    fbuffer := make([]byte,stat.Size()) 

    if _, err := bufio.NewReader(file).Read(fbuffer); err != nil{
        fmt.Println(err)
    }

    file.Close()

    errr := yaml.Unmarshal(fbuffer, &configs)
    if  errr != nil {
        fmt.Printf("error: %v", errr)
    }

    return configs
}
