package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	//r.GET("/", func(context *gin.Context) {
	//	context.String(200, "Hello Geektutu")
	//})

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "who are you?")
	})

	r.GET("/user/name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	r.POST("/form", func(context *gin.Context) {
		username := context.PostFormArray("username")
		password := context.DefaultQuery("password", "000000")

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		name := context.PostFormMap("names")
		context.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": name,
		})
	})

	r.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("goindex", func(context *gin.Context) {
		context.Request.URL.Path = "/"
		r.HandleContext(context)
	})

	r.POST("upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	r.POST("upload2", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

		}
		context.String(http.StatusOK, "%d file upload", len(files))
	})
	_ = r.Run(":9999")

}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("geekutu", "1111")
		context.Next()
		latency := time.Since(t)
		log.Print(latency)
	}
}
