package gigachat

import (
	"github.com/ManyakRus/gigachat/api/protocol"
	"github.com/ManyakRus/starter/config_main"
	"google.golang.org/grpc"
	"testing"
)

func TestChat(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	err := Connect_err()
	defer CloseConnection()
	if err != nil {
		t.Error(err)
	}

	Message1 := &protocol.Message{}
	Message1.Content = "Hello"
	Messages := make([]*protocol.Message, 0)
	Messages = append(Messages, Message1)

	ChatRequest := &protocol.ChatRequest{}
	ChatRequest.Messages = Messages
	opts := []grpc.CallOption{}
	Otvet, err := Chat(ChatRequest, opts...)
	if err != nil {
		t.Error(err)
	}
	if Otvet == nil {
		t.Error("Otvet = nil")
	}
}
