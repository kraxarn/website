module github.com/kraxarn/website

go 1.15

require (
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/gorilla/websocket v1.4.2
	github.com/kraxarn/go-watch v0.0.0-20210211220317-ad2780edf953
	github.com/mattn/go-sqlite3 v1.14.6
)

replace github.com/kraxarn/go-watch => ../go-watch
