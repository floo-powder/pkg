package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Writer Operator interface class for writing data.
type Writer interface {
	Write([]interface{}) BaseResponse
}

type WriterPlugin struct {
	Port uint   `json:"uint"`
	Mode string `json:"mode"`
	Writer
}

func (w *WriterPlugin) GetApplication() *gin.Engine {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	gin.SetMode(w.Mode)
	app := gin.Default()
	app.GET(BaseURL, func(ctx *gin.Context) {
		ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}
		defer ws.Close()
		for {
			var value Payload
			err = ws.ReadJSON(&value)
			if err != nil {
				ws.WriteJSON(BaseResponse{Code: http.StatusBadRequest, Msg: err.Error()})
				continue
			}
			res := w.Write(value.Value)
			ws.WriteJSON(res)
		}
	})
	return app
}
