package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	arr := []int {3,2,4,5,1}
	data,_ := json.Marshal(arr)
	res,err := http.Post("http://localhost:8080/HelloHandle","application/json;charset=utf-8",bytes.NewBuffer([]byte(data)))
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	content,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(content))
	//str := (*string)(unsafe.Pointer(&content))   //优化内存
	//fmt.Println(*str)
}


