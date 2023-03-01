package file

import "github.com/gin-gonic/gin"

// UploadHandler is the handler for uploading a file.
func UploadHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("File name is required"))
		return
	}
	
}