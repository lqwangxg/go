module github.com/lqwangxg/go/server

go 1.15

replace github.com/lqwangxg/go/controllers => ../controllers

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/lqwangxg/go/controllers v0.0.0-00010101000000-000000000000
)
