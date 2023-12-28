package gigachat

import (
	"context"
	"github.com/ManyakRus/gigachat/api/protocol"
	"github.com/ManyakRus/starter/contextmain"
	"google.golang.org/grpc"
	"time"
)

// ListModels_ctx - получить список моделей
func ListModels_ctx(ctx context.Context, ListModelsRequest *protocol.ListModelsRequest, opts ...grpc.CallOption) (*protocol.ListModelsResponse, error) {
	var Otvet *protocol.ListModelsResponse
	var err error

	Otvet, err = ClientModels.ListModels(ctx, ListModelsRequest, opts...)

	return Otvet, err
}

// RetrieveModel_ctx - получить модель
func RetrieveModel_ctx(ctx context.Context, RetrieveModelRequest *protocol.RetrieveModelRequest, opts ...grpc.CallOption) (*protocol.RetrieveModelResponse, error) {
	var Otvet *protocol.RetrieveModelResponse
	var err error

	Otvet, err = ClientModels.RetrieveModel(ctx, RetrieveModelRequest, opts...)

	return Otvet, err
}

// ListModels - получить список моделей
func ListModels(ListModelsRequest *protocol.ListModelsRequest, opts ...grpc.CallOption) (*protocol.ListModelsResponse, error) {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	Otvet, err := ListModels_ctx(ctx, ListModelsRequest, opts...)

	return Otvet, err
}

// RetrieveModel - получить модель
func RetrieveModel(RetrieveModelRequest *protocol.RetrieveModelRequest, opts ...grpc.CallOption) (*protocol.RetrieveModelResponse, error) {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	Otvet, err := RetrieveModel_ctx(ctx, RetrieveModelRequest, opts...)

	return Otvet, err
}
