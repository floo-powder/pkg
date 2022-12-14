package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Reader interface {
	Read()
}

type ReaderPlugin struct {
	Port uint `json:"port"`
	Reader
}

func (r *ReaderPlugin) GetConn() (*websocket.Conn, error) {
	cli := websocket.Dialer{}
	websocketConn, _, err := cli.Dial(
		fmt.Sprintf("ws://127.0.0.1:%d/", r.Port),
		http.Header{},
	)
	return websocketConn, err
}

func (r *ReaderPlugin) Put(websocketConn *websocket.Conn, obj []interface{}) error {
	err := websocketConn.WriteJSON(Payload{Value: obj})
	if err != nil {
		return err
	}
	var res BaseResponse
	websocketConn.ReadJSON(&res)
	if res.Code == 200 {
		return nil
	}
	return errors.New(res.Msg)
}
