package gigachat

import (
	"bytes"
	"errors"
	"fmt"
	protocol "github.com/ManyakRus/gigachat/api/protocol"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/examples/data"
	"io"
	"net/http"
)

// Conn - соединение grpc с GigaChat
var Conn *grpc.ClientConn

// Client - клиент для GigaChat
var ClientChat protocol.ChatServiceClient
var ClientModels protocol.ModelsServiceClient

// Connect - подключается и авторизуется на сервере GigaChat
func Connect() {
	err := Connect_err()
	if err != nil {
		log.Error("Connect() error:", err)
	} else {
		log.Info("Connect() OK")
	}
}

// Connect_err - подключается и авторизуется на сервере GigaChat, возвращает ошибку
func Connect_err() error {
	var err error

	token := &oauth2.Token{
		AccessToken: Settings.Token,
	}

	// Set up the credentials for the connection.
	dir := micro.ProgramDir_bin()
	FileName := dir + "russiantrustedca.pem"
	creds, err := credentials.NewClientTLSFromFile(data.Path(FileName), "")
	if err != nil {
		TextError := fmt.Sprint("NewClientTLSFromFile(), error: ", err)
		err = errors.New(TextError)
		return err
	}
	//creds = insecure.NewCredentials()
	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(token)}
	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(creds),
	}

	addr := Settings.URL
	Conn, err = grpc.Dial(addr, opts...)
	if err != nil {
		TextError := fmt.Sprint("Dial(), error: ", err)
		err = errors.New(TextError)
		return err
	}

	ClientChat = protocol.NewChatServiceClient(Conn)
	ClientModels = protocol.NewModelsServiceClient(Conn)

	return err
}

// CloseConnection - закрывает соединение с GigaChat, возвращает ошибку
func CloseConnection_err() error {
	var err error
	if Conn != nil {
		err = Conn.Close()
	}

	return err
}

// CloseConnection - закрывает соединение с GigaChat
func CloseConnection() {
	err := CloseConnection_err()
	if err != nil {
		log.Error("CloseConnection() error: ", err)
	} else {
		log.Info("CloseConnection() OK")
	}
}

// WaitStop - ожидает отмену глобального контекста
func WaitStop() {

	select {
	case <-contextmain.GetContext().Done():
		log.Warn("Context app is canceled. GigaChat")
	}

	//
	stopapp.WaitTotalMessagesSendingNow("GigaChat")

	//
	CloseConnection()

	stopapp.GetWaitGroup_Main().Done()
}

// Start - делает соединение с GigaChat, отключение и др.
func Start() {
	Connect()

	stopapp.GetWaitGroup_Main().Add(1)
	go WaitStop()

}

func ConnectHttp_err() error {
	client := &http.Client{}

	marshalled := make([]byte, 0)

	URL := Settings.URL + "/api/v1/models"

	req, err := http.NewRequest(
		"POST",
		URL,
		bytes.NewReader(marshalled),
	)
	req.Header.Set("Authorization", "Bearer "+Settings.Token)

	if err != nil {
		log.Fatal("Can't build api request")
		return err
	}

	response, err := client.Do(req)

	if err != nil {
		log.Fatal("Can't send api request: ", err)
		return err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Can't read response body")
		return err
	}

	log.Print(string(responseBody))

	//var giga_response *domain.GigaResponse

	//error := json.Unmarshal(responseBody, &giga_response)
	//
	//if error != nil {
	//	log.Fatal(error, response)
	//	return nil, err
	//}

	return err
}

func GetToken_err() (string, error) {
	Otvet := ""
	var err error

	client := &http.Client{}

	marshalled := make([]byte, 0)

	URL := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	//URL := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"

	req, err := http.NewRequest(
		"POST",
		URL,
		bytes.NewReader(marshalled),
	)

	uuid := uuid.New()
	sUID := uuid.String()

	req.Header.Set("Authorization", "Bearer "+Settings.Token)
	req.Header.Set("RqUID", sUID)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatal("Can't build api request")
		return Otvet, err
	}

	response, err := client.Do(req)

	if err != nil {
		TextError := fmt.Sprint("Can't send api request: ", err)
		log.Error(TextError)
		err = errors.New(TextError)
		return Otvet, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Can't read response body")
		return Otvet, err
	}

	log.Print(string(responseBody))

	return Otvet, err
}
