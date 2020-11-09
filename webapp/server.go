package main

import (
"net/http"
"github.com/labstack/echo"
"github.com/labstack/echo/middleware"
"./controllers"
)

func main() {
   e := echo.New()

   e.Use(middleware.Logger())
   e.Use(middleware.Recover())
   e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
     AllowOrigins: []string{"*"},
     AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
   }))
   e.GET("/", func(c echo.Context) error {
     return c.String(http.StatusOK,"Hello, wolrd!, this is from server written by golang.")
   })
   e.POST("/price", controllers.GrabPrice) //price endpoint

   //Server 
   e.Logger.Fatal(e.Start(":8000"))
}


