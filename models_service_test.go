package gigachat

import (
	"github.com/ManyakRus/gigachat/api/protocol"
	"github.com/ManyakRus/starter/config_main"
	"google.golang.org/grpc"
	"testing"
)

func TestListModels(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	err := Connect_err()
	defer CloseConnection()
	if err != nil {
		t.Error(err)
	}

	ListModelsRequest := &protocol.ListModelsRequest{}
	opts := []grpc.CallOption{}
	Otvet, err := ListModels(ListModelsRequest, opts...)
	if err != nil {
		t.Error(err)
	}
	if Otvet == nil {
		t.Error("Otvet = nil")
	}
}
