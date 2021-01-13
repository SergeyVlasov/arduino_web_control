package helpers
 
import (
    "io/ioutil"
"encoding/json"
"log"
"fmt"
"os"
)
 
func LoadFile(fileName string) (string, error) {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}


func IsEmpty(data string) bool {
    if len(data) <= 0 {
        return true
    } else {
        return false
    }
}

func OpenJSONconfig() (user, connection, port string) {
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


