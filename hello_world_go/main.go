package main

import (
    "fmt"
    "syscall/js"
)

func main() {
    js.Global().Set("helloWorld", js.FuncOf(helloWorld))
    select {}
}

func helloWorld(this js.Value, args []js.Value) interface{} {
    fmt.Println("Hello, World!")
    return nil
}
