package main

import (
	"fmt"
)


type Kesha struct{
	name string
	surname string
}



func Hello(data struct{})struct{}{
	return data


}
func main(){
	kesha := Kesha{"Kesha","Saparow"}
	fmt.Println(kesha)
	fmt.Println(Hello(Kesha{"Sapaorw","dds"}))
}
