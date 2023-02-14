package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Reader interface {
	Read(channel chan []interface{})
}

type ReaderPlugin struct {
	Base
	Reader
}

// GetConn get websocket connect to Push data to writer
func (r *ReaderPlugin) GetConn() (*websocket.Conn, error) {
	cli := websocket.Dialer{}
	websocketConn, _, err := cli.Dial(
		fmt.Sprintf("ws://127.0.0.1:%d/", r.Port),
		http.Header{},
	)
	return websocketConn, err
}

// Push data to writer and get result
func (r *ReaderPlugin) Put(websocketConn *websocket.Conn, channel chan []interface{}) error {
	go r.Read(channel)
	for {
		select {
		case obj := <- channel:
			if len(obj) == 0 {
				return nil
			}
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
		default:
			continue
		}
	}
}
