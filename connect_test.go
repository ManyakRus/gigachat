package gigachat

import (
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestConnect(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	err := Connect_err()
	defer CloseConnection()
	if err != nil {
		t.Error(err)
	}

	if Conn == nil {
		t.Error("Conn = nil")
	}

}
