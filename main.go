package main

import (
	"bytes"
	"io"
	"net/http"
	"html/template"
	
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

func uploadNoteHandler(c *gin.Context){
	fileHeader,err:= c.FormFile("markdown_file")
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"failed to upload file"})
		return
	}
	file,_:=fileHeader.Open()
	defer file.Close()
	contentBytes,_:=io.ReadAll(file)
	//content :=string(contentBytes)

	var buf bytes.Buffer
	if err :=goldmark.Convert(contentBytes,&buf);err !=nil{
		c.JSON(500,gin.H{"message":"Failed to parse markdown"})
		return
	}
	safeHTML :=template.HTML(buf.String())

	c.HTML(http.StatusOK,"result.html",gin.H{
		"content":safeHTML,
	})

	//grammarErrors:=checkGrammar(content)

	
	
}



func main(){
	r := gin.Default()


	r.LoadHTMLGlob("templates/*")

	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",gin.H{})
	})

	r.POST("/upload",uploadNoteHandler)

	r.Run()
}