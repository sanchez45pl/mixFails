package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"os"
	"strconv"
	"math/rand"
	"time"
	"strings"
)

func main() {
	path := os.Args[1]
	last := path[len(path)-1:]
	if last != "/" {
		path += "/"
	}

    files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

	if (os.Args[2] == "-randomise" || os.Args[2] == "-r") {
		for _, file := range files {
			random := rand.New(rand.NewSource(time.Now().UnixNano()))
			randomNumber := random.Intn(len(files))
			err :=  os.Rename( path+file.Name(), path+ strconv.Itoa(randomNumber) + "_" + file.Name()) 		
			  if err != nil {
			  fmt.Println(err)
			  return
			 }

			 if len(os.Args) == 4 {
				if (os.Args[3] == "-show" || os.Args[3] == "-s") {
					fmt.Println(strconv.Itoa(randomNumber) + "_" + file.Name())			
				}
			}
			time.Sleep(50 * time.Millisecond)
		}
	}
	if (os.Args[2] == "-unrandomise" || os.Args[2] == "-u") {
		for _, file := range files {
			oldName := strings.Split(file.Name(), "_")[1]
			err :=  os.Rename( path+file.Name(), path + oldName) 		
			  if err != nil {
			  fmt.Println(err)
			  return
			 }
			 if len(os.Args) == 4 {
				if (os.Args[3] == "-show" || os.Args[3] == "-s") {
					fmt.Println(oldName)
					time.Sleep(50 * time.Millisecond)			
				}
			}
		}
	}
    

	

}