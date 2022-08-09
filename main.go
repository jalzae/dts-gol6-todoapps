package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/controllers"
	filter "rest/middleware"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "connected")
	})

	employeRepo := controllers.EmployeControll()
	r.POST("/employe", employeRepo.Createemploye)
	r.GET("/employe", employeRepo.Getemployes)
	r.GET("/employe/:id", employeRepo.Getemploye)
	r.PUT("/employe/:id", employeRepo.Updateemploye)
	r.DELETE("/employe/:id", employeRepo.Deleteemploye)

	taskRepo := controllers.TaskControll()
	r.POST("/task", taskRepo.Createtask)
	r.GET("/task", taskRepo.Gettasks)
	r.GET("/task/:id", taskRepo.Gettask)
	r.PUT("/task/:id", taskRepo.Updatetask)
	r.DELETE("/task/:id", taskRepo.Deletetask)

	return r

}

//Cors handles cross-domain requests and supports options access
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT,DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//Release all OPTIONS methods
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// process request
		c.Next()
	}
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	key := filter.Getkey()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}
