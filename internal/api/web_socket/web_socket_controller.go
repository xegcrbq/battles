package web_socket

import (
	"battles/internal/users/user_models"
	"battles/internal/utils/registry"
	"encoding/json"
	"github.com/antoniodipinto/ikisocket"
	"strconv"
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
		if ep.SocketAttributes["PublicAddress"] != nil && ep.SocketAttributes["UserId"] != nil {
			c.reg.Log.Debugf("Connected: %v", ep.SocketAttributes["PublicAddress"].(string))
		}
	})

	ikisocket.On(ikisocket.EventMessage, c.wsReq)

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		if ep.SocketAttributes["PublicAddress"] != nil && ep.SocketAttributes["UserId"] != nil {
			c.reg.Log.Debugf("Disconnected: %v", ep.SocketAttributes["PublicAddress"].(string))
			//c.socketService.DeleteSocket(ep.SocketAttributes["username"].(string), ep.SocketUUID)
		}
	})
}

func (c *WSController) wsReq(ep *ikisocket.EventPayload) {
	// !!! По хорошему это нужно в сервис кидать
	if ep.SocketAttributes["PublicAddress"] == nil || ep.SocketAttributes["UserId"] == nil {
		return
	}
	var wsReq WSReq
	err := json.Unmarshal(ep.Data, &wsReq)
	if err != nil {
		c.reg.Log.Warnf(`Incorrect socket req, Err: %v`, err)
		return
	}
	c.reg.Log.Debugf(`wsReq get data: "%v`, wsReq)
	execWSReq(&wsReq, ep)

}
func (c *WSController) SocketReaderCreate(kws *ikisocket.Websocket) {
	data, tkn, _ := c.reg.Tknz.ParseDataClaims(kws.Params("public_address_token"))
	if !tkn.Valid {
		kws.Close()
		return
	}
	answ := registry.Get().Repo.Exec(user_models.QueryUserReadByUserPublicAddress{PublicAddress: data.Data})
	if answ.Err != nil || answ.User == nil {
		kws.Close()
		return
	}
	kws.SetAttribute("PublicAddress", data.Data)
	kws.SetAttribute("UserId", strconv.Itoa(int(answ.User.UserId)))
	return
}
