package tcputil

import (
	"time"
	"net"
	"fmt"
)

//阻塞接收数据
func GetData(conn net.Conn) string{
      dataBytes:=make([]byte, 1024)
      //接收数据
      n,err:=conn.Read(dataBytes)
      //向上层抛出异常
      if(err!=nil){
      	 panic(err)
      }
      //返回字符串数据
      return string(dataBytes[:n])
}
//发送数据
func SendData(conn net.Conn,sendStr string){
	 dataBytes:=[]byte(sendStr)
	 _,err:=conn.Write(dataBytes)
	 //向上层抛出异常
	 if(err!=nil){
	   panic(err)
	 }
}
//获取当前的时间字符串
func GetDateTimeStr() string{
	//获取时间戳
	timestamp :=time.Now().Unix()
	//转换为time.Time类型
	tm:=time.Unix(timestamp,0)
	//格式化为字符串
	str:=tm.Format("2006-01-02 03:04:05 PM")
	return str
}

//暂停屏幕
func ScreenPause(){
	var temp int
	fmt.Scanf("%d",&temp)
}
//异常处理工作
func Recover(conn net.Conn){
	//conn.Close()
    if r:=recover();r!=nil{
    		fmt.Println("客户端异常已关闭:",r)
    	}
}
/**************************************************/
//以下的几个方法是java中采用的方法，
//不过java读写是依靠数据流,明确了缓冲区概念，所以不会出现数组越界的情况，
//Go中测试失败
/*************************************************/
//int32转换为字节数组
func Int32ToByteArray(value int32) []byte{
	bytes := make([]byte, 4)
	bytes[0]=(byte)((value>>24)&0xFF);
	bytes[1]=(byte)((value>>16)&0xFF);
	bytes[2]=(byte)((value>>8)&0xFF);
	bytes[3]=(byte)(value&0xFF);
	return bytes
}
//字节数组转换为int32
func ByteArratToInt32(bytes []byte) int32{
	var value int32
	value=0
	for i:=0;i<4;i++ {
		value+=int32((bytes[i]&byte(0xFF))<<byte(8*(3-i)));
	}
	return value
}
//阻塞接收数据
func JavaGetData(conn net.Conn) string{
      //接收数据长度
      var lenBytes []byte
      _,err:=conn.Read(lenBytes)
      if(err!=nil){
     	 panic(err)
      }
      dataLen:=ByteArratToInt32(lenBytes)
      dataBytes:=make([]byte, dataLen)
      //接收数据正文
      _,err2:=conn.Read(dataBytes)
      if(err2!=nil){
      	 panic(err2)
      }
      //返回字符串数据
      return string(dataBytes)
}
//发送数据
func JavaSendData(conn net.Conn,sendStr string){
	 dataBytes:=[]byte(sendStr)
	 sendLen:=int32(len(dataBytes))
	 lenBytes:=Int32ToByteArray(sendLen)
	 //发送数据长度
	 _,err1:=conn.Write(lenBytes)
	 if(err1!=nil){
	   panic(err1)
	 }
	 //发送数据正文
	 _,err2:=conn.Write(dataBytes)
	 if(err2!=nil){
	   panic(err2)
	 }
}