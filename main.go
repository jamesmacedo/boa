package main

import (
    "os"
    "fmt"
    parse "github.com/jamesmacedo/boa/src"
)

func main(){
    
    if len(os.Args)>2{
        fmt.Println("Please provide an argument!") 
        os.Exit(1)
    }
    parse.Define()
}
