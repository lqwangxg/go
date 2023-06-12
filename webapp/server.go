package main

import (
  "fmt"
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/lqwangxg/go/controllers"
  "github.com/lqwangxg/go/greetings"
  svg "github.com/ajstarks/svgo"
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
  
   

   e.GET("/hello", func(c echo.Context) error {
    greeting, err := greetings.Hello("Dear guest.")
    if err == nil {
      return c.String(http.StatusOK, greeting)
    } else{
      errmsg :=fmt.Sprintf("error.stack: %v.", err)
      return c.String(http.StatusInternalServerError, "greeting.Hello error" + errmsg )
    }
  })
   e.POST("/price", controllers.GrabPrice) //price endpoint
   //e.POST("/svg", circle) //price endpoint
   //http.Handle("/svg", http.HandlerFunc(circle))
   //Server 
   e.Logger.Fatal(e.Start(":8000"))
}

func circle(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "image/svg+xml")
  s := svg.New(w)
  s.Start(800, 800)
  s.Circle(300, 300, 125, "fill:none;stroke:#3ab60b")
  s.End()
}

