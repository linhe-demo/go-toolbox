package websocket

func Run() *Hub {
	hub := NewHub()
	go hub.Run()
	return hub
}
