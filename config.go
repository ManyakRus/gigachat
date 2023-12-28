package gigachat

import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	URL   string
	Token string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	//dir := micro.ProgramDir_bin()

	Settings = SettingsINI{}
	Settings.URL = os.Getenv("GIGACHAT_URL")
	Settings.Token = os.Getenv("GIGACHAT_TOKEN")

	if Settings.URL == "" {
		log.Panic("GIGACHAT_URL = ''")
	}

	if Settings.Token == "" {
		log.Panic("GIGACHAT_TOKEN = ''")
	}
}

func LoadSettingsTxt() {
	var err error

	DirBin := micro.ProgramDir_bin()
	Dir := DirBin + micro.SeparatorFile()
	FilenameEnv := Dir + ".env"
	err = config_main.LoadEnv_from_file_err(FilenameEnv)
	if err == nil {
		return
	}

	FilenameSettings := Dir + "settings.txt"
	err = config_main.LoadEnv_from_file_err(FilenameSettings)
	if err != nil {
		log.Panic("LoadSettingsTxt() filename: ", FilenameSettings, " error: ", err)
	}

}
