package main

import (
    "github.com/DigWing/go-api-boilerplate/db"
    "github.com/joho/godotenv"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "net/http"
)

func main() {
    e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
    }))

    e.Logger.Fatal(e.Start(":8080"))
}

func init() {
    _ = godotenv.Load()
    db.Connect()
}
