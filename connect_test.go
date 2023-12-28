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

func TestConnectHttp_err(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	err := ConnectHttp_err()
	if err != nil {
		t.Error("error: ", err)
	}

	//err := Connect_err()
	//defer CloseConnection()
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if Conn == nil {
	//	t.Error("Conn = nil")
	//}
}

func TestGetToken_err(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	Otvet, err := GetToken_err()
	if err != nil {
		t.Error("error: ", err)
	}

	if Otvet == "" {
		t.Error("Otvet = \"\"")
	}
}
