package main
import (
	"fmt"
	"time"
	"net"
	"encoding/gob"
)

func main(){
	fmt.Println("Cliente")
	var idProcess int
	var count int
	go Cliente(&idProcess, &count)
	var i int
	fmt.Scanln(&i)
	GetProcess(&idProcess)
}

func Process(idProcess *int, count *int) {

	for{
		*count++
		fmt.Println("ID: ",*idProcess, " tiempo: ", *count)
		time.Sleep(time.Second /2)
	}
}

func Cliente(idProcess *int, count *int){

	client, error := net.Dial("tcp", ":9999")
	if error != nil {
		fmt.Println("Error: ",error)
		return
	}
	newClient := true
	error = gob.NewEncoder(client).Encode(newClient)
	if error != nil {
		fmt.Println("Error: ",error)

	}
	error = gob.NewDecoder(client).Decode(idProcess)
	error2 := gob.NewDecoder(client).Decode(count)

	if error != nil || error2 != nil {
		fmt.Println("Error: ",error)
		fmt.Println("Error: ",error2)
		return
	}else {
		go Process(idProcess, count)
	}
	client.Close()
}


func GetProcess(idProcess *int){
	client, error := net.Dial("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		return
	}
	error = gob.NewEncoder(client).Encode(*idProcess)
	if error != nil {
		fmt.Println(error)
	}
}


