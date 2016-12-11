// Author: Arturo Aviles Castellanos
// Creation Date: 18 Nov 2016

package main

import (
    "fmt"
    "os"
    "flag"
    "path/filepath"
    "strings"
    //"io/ioutil"
)

// Global Variables
var numberOfFiles int = 0
var numberOfDirectories int = 0

var fileTypeMap = make(map[string]int)


func main(){
    fmt.Println("Hello World!")

    // Go to Users Home Directory
    //var homeDiretory string = os.Getenv("HOMEPATH") //Remember to uncomment os
    var testDirectory string = "C:\\Users\\100636976\\Desktop\\test"
    fmt.Println(testDirectory)


    flag.Parse()
    err := filepath.Walk(testDirectory, visit)
    fmt.Printf("filepath.Walk() returned %v\n", err)

    fmt.Println("\nNumber of files:", numberOfFiles)
    fmt.Println("Number of directories:", numberOfDirectories)

    fmt.Println(fileTypeMap)
}

func visit(path string, f os.FileInfo, err error) error {
    //fmt.Printf("%s\n", path)
    getFileExtension(path)
    return nil
}

func getFileExtension(path string) error {
    var splitExtensionFromPath = strings.SplitAfterN(path, ".", 2)

    if len(splitExtensionFromPath) > 1 { // File
        fmt.Println(splitExtensionFromPath[1])
        var extension string = splitExtensionFromPath[1]
        addOrUpdateMap(extension)
        numberOfFiles += 1
    } else { // Directory
        fmt.Println("Directory: %s", path)
        numberOfDirectories += 1
    }

    return nil
}

func addOrUpdateMap(extension string) error {
    // Search key in map
    if val, ok := fileTypeMap[extension]; ok { // If found, increase plus 1
        fileTypeMap[extension] = val + 1
    } else { // Add to Map
        fileTypeMap[extension] = 1
    }

    return nil
}
