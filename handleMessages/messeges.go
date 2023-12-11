package handlemessages

import (
	"chat/Type"
	"fmt"
)

func Message() {
	for {
		msg := <-Type.Broadcast

		for client := range Type.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(Type.Clients, client)
			}
		}
	}
}
