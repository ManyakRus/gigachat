package gigachat

import (
	"github.com/ManyakRus/gigachat"
	"github.com/ManyakRus/starter/stopapp"
)

func main() {
	StartApp()
}

func StartApp() {
	gigachat.LoadSettingsTxt()

	gigachat.FillSettings()

	stopapp.StartWaitStop()

	stopapp.GetWaitGroup_Main().Wait()

}
