package middleware

import (
	"net/http"

	"github.com/MichaelDeSteven/rum"
)

// create static handler
func StaticHandler(catchAllParamName, absolutePath string, fs http.FileSystem) rum.HandlerFunc {
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *rum.Context) {
		file := c.Param(catchAllParamName)
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
