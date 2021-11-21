package password

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type baseData struct {
	Title string
	MoreStyles []string
	MoreScripts []string
	Action string
	Date string
}

type appData struct {
	baseData
	Count     string
	Length    string
	Capital   string
	Lower     string
	Special   string
	Numbers   string
	Spaces    string
	Remember  string
	Highlight string
}

func index(c *gin.Context) {
	result := appData{
		Count:     "5",
		Length:    "25",
		Capital:   "true",
		Lower:     "true",
		Special:   "true",
		Numbers:   "true",
		Spaces:    "false",
		Remember:  "false",
		Highlight: "false",
	}
	s := sessions.Default(c)
	result.Date = time.Now().Format("2006")
	result.Title = "Password.fun"
	if s.Get("Count") != nil {
		result.Count = s.Get("Count").(string)
	}
	if s.Get("Length") != nil {
		result.Length = s.Get("Length").(string)
	}
	if s.Get("Capital") != nil {
		result.Capital = s.Get("Capital").(string)
	}
	if s.Get("Lower") != nil {
		result.Lower = s.Get("Lower").(string)
	}
	if s.Get("Special") != nil {
		result.Special = s.Get("Special").(string)
	}
	if s.Get("Spaces") != nil {
		result.Spaces = s.Get("Spaces").(string)
	}
	if s.Get("Numbers") != nil {
		result.Numbers = s.Get("Numbers").(string)
	}
	if s.Get("Highlight") != nil {
		result.Highlight = s.Get("Highlight").(string)
	}
	if s.Get("Remember") != nil {
		result.Remember = s.Get("Remember").(string)
	}

	// Should be moved to new controller and all controllers inherit
	result.Action = "App.Index"
	c.HTML(http.StatusOK, "app/index.html", result)
}
