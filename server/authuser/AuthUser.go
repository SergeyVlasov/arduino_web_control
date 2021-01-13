package authuser

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "fmt"
    "strings"
    helpers "../helpers"
)
 

var database *sql.DB

type device struct{
    id_name string
    password string
}



func UserIsValid(id_name, pwd string) bool {
    
    isValid := false

    user, connection, _ := helpers.OpenJSONconfig()
    db, err := sql.Open(user, connection)
    
    if err != nil {
		log.Fatal(err)
	}
	database = db
	defer db.Close()

    rows, err := db.Query("SELECT name, password FROM public.arduinodevices where name = '" + strings.ToLower(id_name) + "' and password = '" + strings.ToLower(pwd) + "';")
	if err != nil {
        //return "", err
        log.Fatal(err)
	}
    defer rows.Close()
    

    devices := []device{}
    for rows.Next(){
        d := device{}
        err = rows.Scan(&d.id_name, &d.password)
        if err != nil{
            fmt.Println(err)
            continue
        }
        devices = append(devices, d)
    }

    for _, d := range devices{
        if d.id_name != ""{    // чтобы не гонять попусту если у нас несколько записей
            if id_name == d.id_name && pwd == d.password {
                isValid = true
            } else {
                isValid = false
            }
        }
    }



    return isValid
}
