package main

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister reqeuests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:	make(chan []byte),
		register:	make(chan *Client),
		unregister:	make(chan *Client),
		clients:	make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register: // looks fine - login?
			h.clients[client] = true;
		case client := <-h.unregister: // looks fine - logout?
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client);
				close(client.send);
			}
		case message := <-h.broadcast:
			for client := range h.clients { // find out which ones should receive local message (same x, y, or z coordinate)
				select {
				// first determine if client is in range, otherwise delete from list?
				case client.send <- message;
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

