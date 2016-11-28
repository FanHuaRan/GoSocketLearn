package main

import (
	"fmt"
	"net"
	"tcputil"
)

var URL string = "localhost:18894"

func main(){
	conn, err := net.Dial("tcp", URL)
	//异常处理和屏幕暂停
	defer func(conn net.Conn){
		tcputil.Recover(conn)
		tcputil.ScreenPause()
		}(conn)
	if err != nil {
		fmt.Println("与",URL,"连接错误：",err)
		return
	}else{
		fmt.Println("与",URL,"连接已建立")
		//新开协程发送数据
		go ConSendData(conn)
		//主协程中接收数据
		ConGetData(conn)
	}
	
}

//循环接收数据
func ConGetData(conn net.Conn) {
	for {
		getStr := tcputil.GetData(conn)
		fmt.Println("服务端已接收数据=======" + getStr)
	}
}
//发送数据
func ConSendData(conn net.Conn) {
	for {
		var scanStr string
		fmt.Scanf("%s\n", &scanStr)
		tcputil.SendData(conn, scanStr)
	}
}
