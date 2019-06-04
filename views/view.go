package views

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"log"
	"net/http"
	"time"
)
const (
	SecretKey = "thisis secretkey"
)


//jwt success return data
type Token struct {
	Token string `json:"token"`
	Code int `json:"code"`
}

func ChatHome( w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"templates/chat-page.html")
}
func AdminHome( w http.ResponseWriter, r *http.Request){
	serviceId := "1000"
	username, _ := r.Cookie("_uname")
	if username.Value == "admin" {
		fmt.Println("admin")
		serviceId = "1000"
	}else {
		serviceId = "1000"
	}

	myTemplate,_ := template.ParseFiles("templates/admin.html")
	myTemplate.Execute(w,serviceId)
}

func Home(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"templates/login.html")
}

//测试页
func TestHandle(w http.ResponseWriter,r *http.Request){
	myTemplate,err := template.ParseFiles("templates/test.html")
	if err!= nil{
		fmt.Println(err)
	}
	p:=1000
	myTemplate.Execute(w,p)
}

func ServicePeopleHandle(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"templates/servicep.html")
}
func LoginHandle(w http.ResponseWriter, r *http.Request) {

	var username = r.FormValue("username")
	var password = r.FormValue("password")

	fmt.Println(username,password)
	if username != "admin" && password !="123456"{
		log.Println("用户密码错误")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	http.SetCookie(w,&http.Cookie{Name:"_uname",Value:username,HttpOnly:true})
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour*time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString,err := token.SignedString([]byte(SecretKey))
	if err!= nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("500")
	}


	response := Token{Token:tokenString,Code:0}
	JsonResponse(response, w)
}

func ValidateTokenMiddleware(next  http.HandlerFunc) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString ,err:= r.Cookie("token")
		fmt.Println("cookie=",tokenString)
		if err!=nil{
			http.Redirect(w,r, "/", http.StatusFound)
			return
		}

		token,err:=jwt.Parse(tokenString.Value,func(token *jwt.Token)(interface{},error){
			return []byte(SecretKey),nil
		})
		if err == nil {
			//fmt.Println("222")
			if token.Valid {
				next(w, r)
			} else {
				//w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(w,r, "/", http.StatusFound)
				fmt.Fprint(w, "Token is not valid")
				return

			}
		} else {
			//w.WriteHeader(http.StatusMovedPermanently)
			http.Redirect(w,r, "/", http.StatusFound)
			fmt.Fprint(w, "Unauthorized access to this resource")
			return

		}

	})

}
func JsonResponse(response interface{}, w http.ResponseWriter) {

	_json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(_json)
}
