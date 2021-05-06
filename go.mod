module pyke_backend

go 1.14

replace pyke_backend/app/auth v0.0.0 => ./app/auth

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.10.1
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	pyke_backend/app/auth v0.0.0 // indirect
)
