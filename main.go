package main

import (
    "os"
    "fmt"
    boa "github.com/jamesmacedo/boa/src"
)

func main(){
    
    if len(os.Args)<2{
        fmt.Println("Please provide an argument!") 
        os.Exit(1)
    }

    boa.Run(os.Args[1])
}
