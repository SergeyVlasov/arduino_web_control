package helpers

import(
	"encoding/json"
	"log"
	"fmt"
	"os"
)

func openJSONconfig() (user, connection, port string) {
	type configuration struct {
		User string `json:"user"`
		Connection string `json:"connection"`
		Port string `json:"port"`
	}
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
		fmt.Println("неудачная попытка чтения JSON файла конфигурации")
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	user = config.User
	connection = config.Connection
	port = config.Port
	return user, connection, port
}
