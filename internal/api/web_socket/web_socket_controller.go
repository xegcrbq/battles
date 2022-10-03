package web_socket

import (
	"battles/internal/utils/registry"
	"encoding/json"
	"github.com/antoniodipinto/ikisocket"
)

type WSController struct {
	reg *registry.Registry
}

func NewWSController() *WSController {
	c := &WSController{reg: registry.Get()}
	c.Init()
	return c
}

func (c *WSController) Init() {
	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		if ep.SocketAttributes["PublicAddress"] != nil {
			c.reg.Log.Debugf("Connected: %v", ep.SocketAttributes["PublicAddress"].(string))
		}
	})

	ikisocket.On(ikisocket.EventMessage, c.wsReq)

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		if ep.SocketAttributes["PublicAddress"] != nil {
			c.reg.Log.Debugf("Disconnected: %v", ep.SocketAttributes["PublicAddress"].(string))
			//c.socketService.DeleteSocket(ep.SocketAttributes["username"].(string), ep.SocketUUID)
		}
	})
}

func (c *WSController) wsReq(ep *ikisocket.EventPayload) {

	var wsReq WSReq
	err := json.Unmarshal(ep.Data, &wsReq)
	if err != nil {
		c.reg.Log.Warnf(`Incorrect socket req, Err: %v`, err)
		return
	}
	c.reg.Log.Debugf(`wsReq get data: "%v`, wsReq)
}
func (c *WSController) SocketReaderCreate(kws *ikisocket.Websocket) {
	data, tkn, _ := c.reg.Tknz.ParseDataClaims(kws.Params("public_address_token"))
	if !tkn.Valid {
		kws.Close()
		return
	}
	kws.SetAttribute("PublicAddress", data.Data)
	return
}
