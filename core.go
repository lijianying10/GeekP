package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

var connected []websocket.Conn

// Echo response what input
func Echo(ws *websocket.Conn) {
	var err error

	times := 0

	for {

		if times == 0 {
			connected = append(connected, *ws)
		}

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err.Error())
			break
		}
		key := []byte("5e8487e6")

		fmt.Println("Received back from client: " + reply)
		ddd, _ := base64.StdEncoding.DecodeString(reply)
		destext, _ := DesDecrypt(ddd, key)
		fmt.Println("获取解密结果：", string(destext))
		// cmd := exec.Command("geekp" + reply)
		// var out bytes.Buffer
		// cmd.Stdout = &out
		// err := cmd.Run()
		// outs := ""
		// if err != nil {
		// 	outs = err.Error()
		// } else {
		// 	outs = out.String()
		// }

		outs, _ := DesEncrypt(destext, key)
		dist := make([]byte, 2048) //开辟存储空间
		base64.StdEncoding.Encode(dist, outs)
		fmt.Println("加密送出:", string(outs))

		for index := 0; index < len(connected); index++ {
			if err = websocket.Message.Send(&connected[index], string(dist)); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
		times++
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
