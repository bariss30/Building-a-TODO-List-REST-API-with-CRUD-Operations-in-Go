package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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




// TO-DO Listesi modeli
type TodoList struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	CreationDate time.Time  `json:"creation_date"`
	ModifiedDate time.Time  `json:"modified_date"`
	DeletedDate  *time.Time `json:"deleted_date,omitempty"`
	Completion   int        `json:"completion"`
	DeletionDate         *time.Time `json:"deletionDate,omitempty"` // İşaretçi olarak tanımla
	CompletionPercentage int        `json:"completionPercentage"`
	Username             string     `json:"username"` // Kullanıcı adını sakla
}
















// Kullanıcıları saklayan bir dilim (slice)
var users = []User{
	{Username: "user1", Password: "password1", Role: "user1"},
	{Username: "user2", Password: "password2", Role: "user2"},
	{Username: "admin", Password: "adminpass", Role: "admin"},
}

// TO-DO Listelerini saklayan bir dilim (slice)
var todoLists = make(map[string]TodoList)









// Anahtar
var jwtKey = []byte("my_secret_key")

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









//fonk start




// TO-DO listesi oluşturma
func createTodoList(w http.ResponseWriter, r *http.Request) {
	var newList TodoList
	err := json.NewDecoder(r.Body).Decode(&newList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Kullanıcının adını al
	userClaims := r.Context().Value("user").(*Claims)
	newList.Username = userClaims.Username // TO-DO listesine kullanıcı adını ekle

	// ID oluşturma
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	newList.ID = id
	newList.CreationDate = time.Now()
	newList.ModifiedDate = time.Now()

	// TODO listesini saklama
	todoLists[id] = newList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newList)
}
















// TO-DO listelerini listeleme
func getTodoLists(w http.ResponseWriter, r *http.Request) {
	// Kullanıcının adını al
	userClaims := r.Context().Value("user").(*Claims)
	username := userClaims.Username
	role := userClaims.Role

	// Yönetici (admin) ise, tüm kullanıcıların TO-DO listelerini alabilir
	if role == "admin" {
		var allTodoLists []TodoList
		for _, todo := range todoLists {
			if todo.DeletionDate == nil {
				allTodoLists = append(allTodoLists, todo)
			}
		}

		// JSON yanıtı hazırlayalım
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allTodoLists)
		return
	}

	// Diğer kullanıcılar sadece kendi TO-DO listelerini alabilir
	var userTodoLists []TodoList
	for _, todo := range todoLists {
		if todo.DeletionDate == nil && todo.Username == username {
			userTodoLists = append(userTodoLists, todo)
		}
	}

	// JSON yanıtı hazırlayalım
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userTodoLists)
}



















// TO-DO listesini güncelleme
func updateTodoList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // URL parametrelerini al

	// URL'den gelen ID'yi al
	id := params["id"]

	var updatedList TodoList
	err := json.NewDecoder(r.Body).Decode(&updatedList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ID'ye sahip liste var mı kontrol et
	_, ok := todoLists[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Güncelleme tarihi ekleyerek listeyi güncelle
	updatedList.ID = id
	updatedList.ModifiedDate = time.Now()
	todoLists[id] = updatedList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedList)
}
















// TO-DO listesini silme
func deleteTodoList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // URL parametrelerini al

	// URL'den gelen ID'yi al
	id := params["id"]

	// ID'ye sahip liste var mı kontrol et
	todo, ok := todoLists[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Silinme tarihini ayarla
	deletionTime := time.Now()
	todo.DeletionDate = &deletionTime // İşaretçi olarak atanması gerekiyor

	// Güncellenmiş TO-DO listesini tekrar map'e ekle
	todoLists[id] = todo

	w.WriteHeader(http.StatusNoContent) // Başarılı bir yanıt, içerik yok
}
















// TO-DO listesinin tamamlanma yüzdesini güncelleme
func updateCompletionPercentage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var todo TodoList
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ID'ye sahip liste var mı kontrol et
	_, ok := todoLists[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// CompletionPercentage alanını güncelle
	todoList := todoLists[id]
	todoList.CompletionPercentage = todo.CompletionPercentage
	todoLists[id] = todoList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoLists[id])
}



// fonks end













func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// İstek başlıklarından Authorization başlığını al
		tokenString := r.Header.Get("Authorization")

		// "Bearer " ön eki ile gelen JWT'yi ayır
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Token doğrulama işlemi
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Token'ı işleme
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Context'e kullanıcı bilgilerini ekle
		ctx := context.WithValue(r.Context(), "user", claims)

		// Sonraki işlemi devam ettir
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
















func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/home", homePage).Methods("GET")
	r.HandleFunc("/todo/create", authenticate(createTodoList)).Methods("POST") // authenticate middleware kullanılacak
	r.HandleFunc("/todo/lists", authenticate(getTodoLists)).Methods("GET")
	r.HandleFunc("/todo/update/{id}", authenticate(updateTodoList)).Methods("PUT") // PUT request ile güncelleme
	r.HandleFunc("/todo/delete/{id}", authenticate(deleteTodoList)).Methods("DELETE")
	r.HandleFunc("/todo/completion/{id}", authenticate(updateCompletionPercentage)).Methods("PUT") // Tamamlanma yüzdesi güncelleme endpointi

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}

