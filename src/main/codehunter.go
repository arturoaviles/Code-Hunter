// Author: Arturo Aviles Castellanos
// Creation Date: 18 Nov 2016

package main

import (
    "fmt"
    //"path/filepath"

    "io/ioutil"
)

func main(){
    fmt.Println("Hello World!")

    files, _ := ioutil.ReadDir("./")
    for _, f := range files {
            fmt.Println(f.Name())
    }

}
