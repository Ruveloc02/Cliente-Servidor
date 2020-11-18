package main
import (
	"fmt"
	"net"
	"encoding/gob"
	"time"
)

func main(){
	fmt.Println("Servidor")
	count := 0
	processes := []int{1,2,3,4,5}
	go Processes(&processes, &count)
	go Server(&processes, &count)
	var i int
	fmt.Scanln(&i)
}

func Processes(processes *[]int, count *int) {
	for{
		*count++
		for _, i := range *processes {
			fmt.Println("ID: ",i, " tiempo: ", *count)
		}
		fmt.Println("")
		time.Sleep(time.Second /2)
	}
}

func Client(client net.Conn, processes *[]int, count *int){
	var process int
	error := gob.NewDecoder(client).Decode(&process)
	if error != nil {
		fmt.Println("Error: ",error)
	}else {
		*processes = append(*processes, process)
		return
	}
	newProcesses := *processes
	idProcess :=  newProcesses[0]
	*processes = append(newProcesses[:0], newProcesses[0+1:]... )
	error = gob.NewEncoder(client).Encode(idProcess)
	if error != nil {
		fmt.Println("Error: ",error)
	}
	error = gob.NewEncoder(client).Encode(count)
	if error != nil {
		fmt.Println("Error: ",error)
	}
}

func Server(processes *[]int, count *int) {
	serv, error := net.Listen("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		return
	}
	for{
		client, error := serv.Accept()
		if error != nil{
			fmt.Println(error)
			continue
		}
		go Client(client, processes, count)
	}
}




