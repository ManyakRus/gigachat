package gigachat

import (
	"errors"
	"fmt"
	protocol "github.com/ManyakRus/gigachat/api/protocol"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/stopapp"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/examples/data"
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
