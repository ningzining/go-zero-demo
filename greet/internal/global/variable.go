package global

import "github.com/gorilla/websocket"

// 保存连接信息
var WsMap = make(map[string]*websocket.Conn)

// 保存code信息
var CodeMap = make(map[string]string)
