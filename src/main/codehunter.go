// Author: Arturo Aviles Castellanos
// Creation Date: 18 Nov 2016

package main

import (
    "fmt"
    "os"
    "flag"
    "path/filepath"
    "strings"
    "encoding/json"
    "io/ioutil"
    //"reflect"
)


// Global Variables
var numberOfFiles int = 0
var numberOfDirectories int = 0

var fileTypeMap = make(map[string]int)
var fileTypeMapPercentage = make(map[string]float64)


func main(){
    fmt.Println("Hello World!")

    // Go to Users Home Directory
    //var homeDirectory string = os.Getenv("HOMEPATH") //Remember to uncomment os
    //fmt.Println(homeDirectory)

    var testDirectory string = "C:\\Users\\100636976\\Desktop\\test"
    fmt.Println(testDirectory)
    //var testDirectory string = "C:\\Users\\100636976\\Desktop\\test"



    flag.Parse()
    err := filepath.Walk(testDirectory, visit)
    fmt.Printf("filepath.Walk() returned %v\n", err)

    fmt.Println("\nNumber of files:", numberOfFiles)
    fmt.Println("Number of directories:", numberOfDirectories)
    fmt.Println(fileTypeMap)

    // Get Percentages
    var percentage float64 = 0
    for key, value := range fileTypeMap {
        percentage = (float64(value) * float64(100)) / float64(numberOfFiles)
        fileTypeMapPercentage[key] = percentage
    }

    fmt.Println(fileTypeMapPercentage)

    // Map to Json
    jsonParse, _ := json.Marshal(fileTypeMapPercentage)

    var variableMake string = "var dictionary =" + string(jsonParse) + ";"
    fmt.Println(variableMake)

    var jsDictionary []byte = []byte(variableMake)

    // Json to File
    err = ioutil.WriteFile("data.json", jsDictionary, 0644) // To save as plain json change jsDictionary to jsonParse
    if err != nil {
        panic(err)
    }

    fmt.Println("Open index.html")
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
