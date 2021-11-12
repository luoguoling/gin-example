package whiteip

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := context.ClientIP()
		fmt.Println(clientIP)
		for _, host := range ipList {
			if clientIP == host {
				flag = true
			}
		}
		if !flag {
			context.String(401, "%s,not in iplist", clientIP)
			context.Abort()
			return
		}
		context.Next()

	}
}
