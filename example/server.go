package main

import (
	"github.com/sulthonzh/transmit"
	"github.com/sulthonzh/transmit/debug"
	"log"
	// "time"
)

var tp = transmit.New()

func main() {
	// 开启Teleport错误日志调试
	debug.Debug = false
	// debug.Debug = true
	tp.SetUID("abc").SetAPI(transmit.API{
		"报到": new(报到),

		// 短链接不可以直接转发请求
		"短链接报到": new(短链接报到),
	}).Server(":20125")
	// time.Sleep(30e9)
	// tp.Close()
	select {}
}

type 报到 struct{}

func (*报到) Process(receive *transmit.NetData) *transmit.NetData {
	log.Printf("报到：%v", receive.Body)
	// 直接回复
	// return transmit.ReturnData("服务器："+receive.From+"客户端已经报到！", "报到")
	// 转发形式一
	return transmit.ReturnData("服务器："+receive.From+"客户端已经报到！", "报到", "C3")
}

type 短链接报到 struct{}

func (*短链接报到) Process(receive *transmit.NetData) *transmit.NetData {
	log.Printf("报到：%v", receive.Body)
	// 请求或转发形式二
	tp.Request("服务器："+receive.From+"客户端已经报到！", "报到", "", "C3")
	return nil
}
