// Package ipc - interprocesses communication
package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

// Connect - session 设计就是关于一个原生管道的属性变化
func (server *IpcServer) Connect() *chan string {
	// 闭包获取 session 实参，执行读取
	var session = make(chan string)
	for {
		go func(c chan string) {
			for {
				request := <-c
				if request == "CLOSE" {
					fmt.Println("Closing session proactively")
					close(c)
					return
				}
				var req Request
				err := json.Unmarshal([]byte(request), &req)
				if err != nil {
					fmt.Println("Error: invalid request format:", request)
					continue
				}
				response := server.Handle(req.Method, req.Params)
				b, err := json.Marshal(response)
				if err != nil {
					fmt.Println("Error: ", err, "-", request)
					continue
				}
				c <- string(b)
			}
		}(session)
		fmt.Println("session opened")
		return &session
	}
}

func (server *IpcServer) Handle(method, params string) *Response {
	var response *Response
	switch method {
	case "GET":
		response = &Response{Code: "SUCCESS", Body: "echo get:" + params}
	case "POST":
		response = &Response{Code: "SUCCESS", Body: "echo post:" + params}
	}
	return response
}

func (server *IpcServer) Name() string {
	return "EchoServer"
}
