package main
import(
    "fmt"
    "net"
    "tcputil"
)
var URL string = "localhost:18893"
//客户连接处理函数
func handleConn(conn net.Conn) {
    defer tcputil.Recover(conn)
    for {
   	 //接收数据
   	  getStr:=tcputil.GetData(conn)
   	  fmt.Println(conn.RemoteAddr(),":",getStr)
      //发送时间数据
      timeStr:=tcputil.GetDateTimeStr()
      tcputil.SendData(conn,timeStr)
    }
}
//TCP服务端
func main() {
    serverSocket, err := net.Listen("tcp", URL)
    //异常处理和屏幕暂停
    defer func(listener net.Listener){
    	if r:=recover();r!=nil{
    		fmt.Println("服务器异常，已关闭",r)
    	}
    	tcputil.ScreenPause()
    }(serverSocket)
    if err != nil {
        fmt.Println("服务器打开错误:",err)
        return
    }
    fmt.Println("服务器打开成功:",URL)
    for {
    	fmt.Println("等待连接.....")
        con, err := serverSocket.Accept()
        if err != nil {
            fmt.Println("接收连接错误:", err)
            break
        }
        fmt.Println("接收到新连接：",con.RemoteAddr())
        //开始一个新协程来处理客户连接
        go handleConn(con)
    }
}

