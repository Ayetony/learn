package ipc

import "encoding/json"

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{*c}
}

func (client *IpcClient) Call(method, params string) (*Response, error) {
	req := &Request{method, params}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 客户端获取 session（conn）chan ，执行写入数据
	client.conn <- string(b)
	str := <-client.conn // wait for the response
	var resp Response
	err = json.Unmarshal([]byte(str), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
