package password

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func faq(c *gin.Context) {
	b := baseData{}
	b.Date = time.Now().Format("2006")
	b.Title = "Password.fun"
	b.Action = "Faq.Index"
	c.HTML(http.StatusOK, "faq/index.html", b)
}
