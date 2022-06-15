package main

 import (
   "fmt"
   "io"
   "os"
   "bufio"
   "bytes"
 )

 func main () {
     file, err := os.Open("testdata.txt")

     if err != nil {
       panic(err.Error())
     }

     defer file.Close()

     reader := bufio.NewReader(file)
     buffer := bytes.NewBuffer(make([] byte,0))


     var chunk []byte
     var eol bool
     var strArray []string

     for {
       if chunk, eol, err = reader.ReadLine(); err != nil {
         break
       }
       buffer.Write(chunk)
       if !eol {
          strArray = append(strArray, buffer.String())
          buffer.Reset()
       }
     }

     if err == io.EOF {
        err = nil
     }

     fmt.Println(strArray)

     var x []byte

     for i:=0; i<len(strArray); i++{
         b := []byte(strArray[i])
         for j:=0; j<len(b); j++{
             x = append(x,b[j])
         }
     }

     fmt.Println(x)

     str := ""
     for i := 0; i < len(x); i++ {
         str += string(x[i])
     }

     fmt.Println(str)
 }
 