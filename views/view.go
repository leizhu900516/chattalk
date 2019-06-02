package views

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)
const (
	SecretKey = "thisis secretkey"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	Token string `json:"token"`
	Code int `json:"code"`
}

func ChatHome( w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"templates/chat-page.html")
}
func AdminHome( w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"templates/admin.html")
}

func Home(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"templates/login.html")
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	//var user UserCredentials
	//fmt.Println(r.Body)
	fmt.Println(r.FormValue("username"))
	var username = r.FormValue("username")
	var password = r.FormValue("password")

	fmt.Println(username,password)
	if username != "admin" && password !="123456"{
		log.Println("用户密码错误")
		w.WriteHeader(http.StatusForbidden)
		return
	}
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
		//token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		//	func(token *jwt.Token) (interface{}, error) {
		//		return []byte(SecretKey), nil
		//	})

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
			fmt.Println("222")
			if token.Valid {
				next(w, r)
			} else {
				//w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(w,r, "/", http.StatusFound)
				fmt.Fprint(w, "Token is not valid")
			}
		} else {
			fmt.Println("333")
			log.Println(err)
			w.WriteHeader(http.StatusMovedPermanently)
			http.Redirect(w,r, "/", http.StatusFound)
			fmt.Fprint(w, "Unauthorized access to this resource")
		}

	})

}
func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
