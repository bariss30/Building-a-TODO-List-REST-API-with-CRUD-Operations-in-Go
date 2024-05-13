package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Kullanıcı modeli
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // Kullanıcı rolü (admin, user1 veya user2)
}

// Anahtar
var jwtKey = []byte("my_secret_key")

// Kullanıcıları saklayan bir dilim (slice)
var users = []User{
	{Username: "user1", Password: "password1", Role: "user1"},
	{Username: "user2", Password: "password2", Role: "user2"},
	{Username: "admin", Password: "adminpass", Role: "admin"},
}

// Token yapısı
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Kullanıcının JWT'sini oluşturur
func generateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Login işlemi
func login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			token, err := generateJWT(u.Username, u.Role)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			response := fmt.Sprintf("Hoşgeldin, %s! Token: %s", u.Role, token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Geçersiz kullanıcı adı veya şifre"))
}

// Ana sayfa
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ana Sayfa")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/home", homePage).Methods("GET")

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
