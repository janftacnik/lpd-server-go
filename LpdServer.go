/* LPD Server
*/
package main

import (
  "fmt"
  "net"
  "os"
  "strconv"
  "io"
  "io/ioutil"
)

func main() {
   var name string
   service := ":515"
   icount  := 0
   name1   := "prjob"
   tcpAddr, err := net.ResolveTCPAddr("tcp", service)
   checkError(err)

   listener, err := net.ListenTCP("tcp", tcpAddr)
   checkError(err)
   for {
        conn, err := listener.Accept()
        if err != nil {
           continue
        }
        // run as a goroutine
        icount = icount + 1
        name = name1 + strconv.Itoa(icount) 
        go handleClient(conn,name)
       }
    }

func handleClient(conn net.Conn, filename string) {
     // close connection on exit
     defer conn.Close()
          

     var buf [512]byte
     var code byte = 0x0     
     for {
           ack := []byte{code}
           
           // read up to 512 bytes
           n, err := conn.Read(buf[0:])
           if err != nil {
              return
           }
           
           // write the n bytes read
           _, err2 := conn.Write(ack)
           if err2 != nil {
              fmt.Fprintf(os.Stderr,"Fatal error: %s", err2.Error())
              return
           }
           
           // read up to 512 bytes
           n, err = conn.Read(buf[0:])
           if err != nil {
              return
           }
           
            // write the n bytes read
            _, err2 = conn.Write(ack)
            if err2 != nil {
               fmt.Fprintf(os.Stderr,"Fatal error: %s", err2.Error())
               return
            }
           
           // read up to 512 bytes
           n, err = conn.Read(buf[0:])
           if err != nil {
              return
           }
           
           errc := ioutil.WriteFile(filename+".cfg", buf[0:n],0666)
           if errc != nil {
              fmt.Fprintf(os.Stderr, "Fatal error: %s", errc.Error())
           }           
                      
           // write the n bytes read
           _, err2 = conn.Write(ack)
           if err2 != nil {
              fmt.Fprintf(os.Stderr,"Fatal error: %s", err2.Error())
              return
           }
           
           // read up to 512 bytes
           n, err   = conn.Read(buf[0:])
           if err  != nil {
              return
           }
           
           // write the n bytes read
           _, err2  = conn.Write(ack)
           if err2 != nil {
              fmt.Fprintf(os.Stderr,"Fatal error: %s", err2.Error())
              return
           }
                      
           // read up to 512 bytes
           n, err   = conn.Read(buf[0:])           
           if err != nil {
              return
           }
           
           err5 := ioutil.WriteFile(filename+".prn", buf[0:n],0666)
             if err5 != nil {
               fmt.Fprintf(os.Stderr, "Fatal error: %s", err5.Error())
             }           
           
            //OPEN FILE TO APPEND INTO
           file, err6 := os.OpenFile(filename+".prn", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
            if err6 != nil {
              fmt.Fprintf(os.Stderr, "Fatal error: %s", err6.Error())
              return
            }
           defer file.Close()

           
           // write the n bytes read
           _, err2  = conn.Write(ack)
           if err2 != nil {
              fmt.Fprintf(os.Stderr,"Fatal error: %s", err2.Error())
              return
           }
           
           // read up to 512 bytes
           
          _, err = io.Copy(file, conn) // **IT DOES NOT STOP!!**
            if err != nil {
               fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
               return
            }
                      
           // write the n bytes read
           _, err2  = conn.Write(ack)
           if err2 != nil {
              fmt.Fprintf(os.Stderr,"Fatal error: %s", err2.Error())
              return
           }
       fmt.Printf("Saved file %s \n", filename+".prn")              
     }
}

func checkError(err error) {
       if err != nil {
              fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
              os.Exit(1)
       }
}