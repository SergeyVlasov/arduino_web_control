package handlers
 
import (
    "fmt"
    "net/http"
    
    helpers "../helpers"
    authuser "../authuser"
    cookie "../cookie"

	"strings"
    //pages_handle "../pages_handle"

    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "html/template"
    "strconv"
)
 
 



// --------------------------------AUTORIZATION------------------------------------
// for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {

    //userName := GetUserName(request)    
    //fmt.Println("Сейчас вошел пользователь:    ",strings.ToLower(userName)) // debug

    var body, _ = helpers.LoadFile("templates/login.html")
    fmt.Fprintf(response, body)
}
 
// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("name")
    //name1 := request.FormValue("name1")    // for debug ( only name no working)
    pass := request.FormValue("password")
    redirectTarget := "/"
    _userIsValid := authuser.UserIsValid(name, pass)
    //fmt.Println("name: ", name, "\n", "name1: ", name1, "\n", "pass: ", pass, "\n", _userIsValid)    // for debug
   
    if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
        // Database check for user data!
        if _userIsValid {
            cookie.SetCookie(name, response)
            redirectTarget = "/main"    // если имя и пароль валидны то перенаправляем на главную страницу
        } 
    }
    
    //fmt.Println(redirectTarget)
    http.Redirect(response, request, redirectTarget, 302)
}
 


var database *sql.DB
type Rele struct {
    rele_name string
    rele_status bool
  }
 
// for GET
func MainPageHandler(response http.ResponseWriter, request *http.Request) {
    userName := cookie.GetUserName(request)  //  имя пользоваеля из куки
    if !helpers.IsEmpty(userName) {

        //--------------page after autorization-------------

            db, err := sql.Open("postgres", "dbname=arduino user=postgres password=poilo777 host=localhost sslmode=disable")
        
            if err != nil {
                log.Fatal(err)
            }
            database = db
            defer db.Close()
        
            //fmt.Println("id_name   ", id_name)
            rows, err := db.Query("SELECT rele_name, rele_status FROM public.arduinodevices where name = '" + strings.ToLower(userName) + "';")
            if err != nil {
                log.Fatal(err)
            }
            defer rows.Close()
        
        
            dev := make([]*Rele, 0) // slice of struct Device
            for rows.Next() {
            dv := new(Rele)
            err := rows.Scan(&dv.rele_name, &dv.rele_status)
            if err != nil {
              log.Fatal(err)
            }
            dev = append(dev, dv)
            }
            if err = rows.Err(); err != nil {
                log.Fatal(err)
            }
        
            type Rele_view struct {
                Rele_number string                // struct to add data from SQL-query to HTML
                Rele_name string
                Rele_status bool
                }
               
              type Data struct {                             // additional struct for loop 
                  Items []Rele_view
            }
            
              data := Data{}
              i := 1
              for _, dv := range dev {                          // loop
                view := Rele_view{
                    Rele_number: strconv.Itoa(i),
                    Rele_name: dv.rele_name,
                    Rele_status: dv.rele_status,
                }	

                    i = i + 1
                    data.Items = append(data.Items, view)
              }
              tmpl, _ := template.ParseFiles("./templates/index.html")     // parsing HTML to web page
              tmpl.Execute(response, data)


        //--------------------------------------------------

    } else {
        http.Redirect(response, request, "/", 302)
    }
}
 
// for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
    cookie.ClearCookie(response)
    http.Redirect(response, request, "/", 302)
}
 

func SetStatusHandler(response http.ResponseWriter, request *http.Request) {

    //userName := GetUserName(request)  //  имя пользоваеля из куки

    name := request.FormValue("rele1")
    fmt.Println(name)
    http.Redirect(response, request, "/main", 302)
}

