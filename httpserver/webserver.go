package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func QuickSort(data []int, left, right int) {
	if left >= right {
		return
	}

	tmpVal := data[left]
	mid := left
	// 划分数组，确定tmpVal的位置
	for i := left + 1; i <= right; i++ {
		// 从小到大排序，改为>则从大到小
		if data[i] < tmpVal {
			data[mid], data[i] = data[i], data[mid+1]
			mid++
		}
	}

	data[mid] = tmpVal
	QuickSort(data, left, mid-1)
	QuickSort(data, mid+1, right)
}


func HelloHandle(w http.ResponseWriter,r *http.Request) {
	//方法，URL结构体，协议
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto) //写入到 w 的是输出到客户端的数据
	//头部信息
	for k,v := range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
	//首部Host
	fmt.Fprintf(w,"Host = %q\n",r.Host)
	//请求地址
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)

	//解析body参数，会将参数写入Form字段和PostForm字段当中
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	//参数查询的数据
	for k,v := range r.Form {
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}

	//-----------------------------------------------
	//读取客户端Post传过来的数据
	body,_ := ioutil.ReadAll(r.Body)
	var arr []int
	//将字符串反解析为切片
	json.Unmarshal([]byte(body),&arr)
	//排序
	QuickSort(arr,0,len(arr)-1)
	//将切片解析为字符串,并输出
	str,_ := json.Marshal(arr)
	w.Write(str)
	fmt.Println()
}

// 接口URL地址：http://localhost:8080
// POST或GET
//POST时:Body->raw->json 参数可填：[2,3,1,4,5]

func main() {
	http.HandleFunc("/HelloHandle",HelloHandle)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServer:",err.Error())
	}
}

