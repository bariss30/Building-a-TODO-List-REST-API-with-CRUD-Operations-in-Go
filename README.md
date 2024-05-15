# To-Do-list-with-go


Go ile ilgili hiçbir fikrim yok

[GoLang  ](https://medium.com/i%CC%87yi-programlama/go-programala-dili-temel-bilgiler-6f7b0b8597)


                                                 --------------------------------------------
JSON Web Tokens (JWTs) are a popular method for dealing with online authentication, and you can implement JWT authentication in any server-side programming language.

[JWT nedir kullanım amacı  ](https://tugrulbayrak.medium.com/jwt-json-web-tokens-nedir-nasil-calisir-5ca6ebc1584a)

[JWT Kullanım ]( https://blog.logrocket.com/jwt-authentication-go/)

Getting started with the Golang-JWT package
Setting up a web server in Go
Generating JWTs for authentication using the Golang-JWT package
.
.
.
.






 
                                                  --------------------------------------------


Login
Kullanıcı bir kullanıcı adı ve parola yollar eğer bilgiler doğru ise bir JWT oturum anahtarı alır.


![Alternatif Metin](https://github.com/bariss30/To-Do-list-with-go/blob/main/sehsarha.png)


![Alternatif Metin](https://github.com/bariss30/To-Do-list-with-go/blob/main/erhaerh.png)


![Alternatif Metin](https://github.com/bariss30/To-Do-list-with-go/blob/main/seharsh.png)


![Alternatif Metin](https://github.com/bariss30/To-Do-list-with-go/blob/main/sehsarha.png)




                                                   --------------------------------------------

[Postman Mock Servis Kullanımı ](https://hakankaplan.medium.com/postman-ile-mock-servis-bfdbcad89284#:~:text=Bu%20yaz%C4%B1mda%20mock%20servisin%20ne%20oldu%C4%9Funu%2C%20neden%20ihtiya%C3%A7,taklit%20edebiliriz%20ya%20da%20sahte%20bir%20servis%20olu%C5%9Fturabiliriz.)


 Bu, verilerinizi yönetmek için geçici bir yöntem olarak hizmet ederken, uygulamanızı geliştirirken veya test ederken işleri kolaylaştırabilir.

 [!!! BÜYÜK YARDIMCI KAYNAK !!!](https://dev.to/permify/implementing-jwt-authentication-in-a-golang-application-onf)


                                                  --------------------------------------------




Middleware Kullanımı: authenticate adlı bir ara yazılım, JWT doğrulamasını gerçekleştiriyor ve yetkilendirilmiş kullanıcı bilgilerini istek bağlamına (Context) ekliyor. Bu, belirli rotalara erişimi sınırlamak için kullanılıyor.













[Top 5 Golang Frameworks](https://masteringbackend.com/posts/top-5-golang-frameworks)


Framework ler kulanarak zenginleştirilebilir




                                        ---------------------------------------------------


   ![Yetkilendirme kullanım örnek ](https://github.com/bariss30/To-Do-list-with-go/blob/main/agewbgao.gif)












                                     ---------------------------------------------------

Fonksiyonların Endpoint'lerle kullanımı


"/todo/lists" Methods("GET")
"/todo/update/{id}" Methods("PUT")              // PUT request ile güncelleme
"/todo/delete/{id}"  Methods("DELETE")
"/todo/completion/{id}" Methods("PUT")            body{"completionPercentage": 75}    // Tamamlanma yüzdesi güncelleme endpointi

                                    
                                       ---------------------------------------------------

                                         
işte kodun özeti:

Bu kod bir HTTP sunucusu uygulamasıdır. Uygulama, kullanıcıların giriş yapmasına, to-do listeleri oluşturmasına, güncellemesine, silmesine ve to-do listelerini listelemesine olanak tanır. Ayrıca, to-do listelerinin tamamlanma yüzdesini güncellemek için bir endpoint sağlar. Uygulama, Gorilla Mux ve JWT Go kütüphanelerini kullanır.


Bu kod, kullanıcı yönetimi ve temel to-do listesi işlemlerini gerçekleştiren basit bir RESTful API sunar.
