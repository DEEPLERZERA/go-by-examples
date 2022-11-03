package main

import ( // import the os package
    "fmt"
    "os"
)

func main() { // main function

    f := createFile("/tmp/defer.txt") // create a file
    defer closeFile(f) 			  // defer the closing of the file
    writeFile(f) 				  // write to the file
}

func createFile(p string) *os.File { // create a file
    fmt.Println("creating") // print creating
    f, err := os.Create(p)  // create the file
    if err != nil { 	   // if there is an error
        panic(err) 		   // panic
    }
    return f 			   // return the file
}

func writeFile(f *os.File) { // write to the file
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) { // close the file
    fmt.Println("closing") // print closing
    err := f.Close() 	   // close the file

    if err != nil { 	   // if there is an error
        fmt.Fprintf(os.Stderr, "error: %v\n", err) // print the error
        os.Exit(1) 							   // exit with a status of 1
    }
}