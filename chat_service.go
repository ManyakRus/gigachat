package gigachat

import (
	"context"
	"github.com/ManyakRus/gigachat/api/protocol"
	"github.com/ManyakRus/starter/contextmain"
	"google.golang.org/grpc"
	"time"
)

// Chat_ctx - отправляет сообщение на сервер GigaChat
func Chat_ctx(ctx context.Context, ChatRequest *protocol.ChatRequest, opts ...grpc.CallOption) (*protocol.ChatResponse, error) {
	var Otvet *protocol.ChatResponse
	var err error

	Otvet, err = ClientChat.Chat(ctx, ChatRequest, opts...)

	return Otvet, err
}

// ChatStream_ctx - отправляет сообщение стрим на сервер GigaChat
func ChatStream_ctx(ctx context.Context, ChatRequest *protocol.ChatRequest, opts ...grpc.CallOption) (protocol.ChatService_ChatStreamClient, error) {
	var Otvet protocol.ChatService_ChatStreamClient
	var err error

	Otvet, err = ClientChat.ChatStream(ctx, ChatRequest, opts...)

	return Otvet, err
}

// Chat - отправляет сообщение на сервер GigaChat
func Chat(ChatRequest *protocol.ChatRequest, opts ...grpc.CallOption) (*protocol.ChatResponse, error) {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	Otvet, err := Chat_ctx(ctx, ChatRequest, opts...)

	return Otvet, err
}

// ChatStream - отправляет сообщение стрим на сервер GigaChat
func ChatStream(ChatRequest *protocol.ChatRequest, opts ...grpc.CallOption) (protocol.ChatService_ChatStreamClient, error) {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	Otvet, err := ChatStream_ctx(ctx, ChatRequest, opts...)

	return Otvet, err

}
