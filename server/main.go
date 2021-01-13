package main
 
import (
    "net/http"
  
    pagehandler "./pagehandler"
	helpers "./helpers"

	"strings"
    "log"
	"fmt"
	"database/sql"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "encoding/json"
)
 
var router = mux.NewRouter()
 
func main() {
	_, _, port := helpers.OpenJSONconfig()

	router.HandleFunc("/", pagehandler.LoginPageHandler) // GET страница для пользователей
	router.HandleFunc("/status_set", pagehandler.StatusSetPageHandler) // GET страница для пользователей

	router.HandleFunc("/main", pagehandler.MainPageHandler) // GET  страница для пользователей
	router.HandleFunc("/setdevicestatus", pagehandler.SetStatusHandler) // POST страница для пользователей
	
    router.HandleFunc("/login", pagehandler.LoginHandler).Methods("POST") //  страница для пользователей
    router.HandleFunc("/logout", pagehandler.LogoutHandler).Methods("POST")  //  страница для пользователей
 
    router.HandleFunc("/pagefordevices/{id_device}", query_to_db) // страница для устройств

    http.Handle("/", router)
    http.ListenAndServe(port, nil)
}






// ------------------------for device ---------------------------------


var database *sql.DB


func query_to_db(w http.ResponseWriter, r *http.Request)  {
	
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id_device := str_handle(vars["id_device"]) // id (name) девайса из URL адреса

	rows, _ := getJSON(id_device)
	fmt.Fprintf(w, rows)
	
	
}


func getJSON(id_name string) (string, error) {

	user, connection, _ := helpers.OpenJSONconfig()

	db, err := sql.Open(user, connection)

	if err != nil {
		log.Fatal(err)
	}
	database = db
	defer db.Close()

    //fmt.Println("id_name   ", id_name)
	rows, err := db.Query("SELECT rele_name, rele_status FROM public.arduinodevices where name = '" + strings.ToLower(id_name) + "';")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func str_handle(inpt string) (outpt string) {  //some sequrity

	filter0 := strings.ToTitle(inpt)

	filter1 := strings.Replace(filter0, ";", "", -1)
	filter2 := strings.Replace(filter1, "'", "", -1)
	filter3 := strings.Replace(filter2, "%", "", -1)
	filter4 := strings.Replace(filter3, "&", "", -1)
	filter5 := strings.Replace(filter4, "?", "", -1)
	filter6 := strings.Replace(filter5, "drop", "", -1)
	filter7 := strings.Replace(filter6, "table", "", -1)
	filter8 := strings.Replace(filter7, "delete", "", -1)
	filter9 := strings.Replace(filter8, "alter", "", -1)
	outpt = filter9
	return outpt
}
