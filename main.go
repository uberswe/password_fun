package password

import (
	"embed"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-acme/lego/log"
	"html/template"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//go:embed public/*
var static embed.FS

//go:embed templates/*
var templates embed.FS

func Run() {
	port := ":80"
	sessionSecret := "secret"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if envSessionSecret := os.Getenv("SESSION_SECRET"); envSessionSecret != "" {
		sessionSecret = envSessionSecret
	}

	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	store := cookie.NewStore([]byte(sessionSecret))
	r.Use(sessions.Sessions("password_fun_session", store))
	r.StaticFS("/assets", http.FS(static))

	r.GET("/", index)
	r.GET("/faq", faq)
	r.GET("/api/v1", api)

	r.Any("/api/v1/passwords", passwords)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "errors/404.html", nil)
	})

	log.Println(fmt.Sprintf("Listening on %s", port))
	err = r.Run(port)
	if err != nil {
		panic(err)
	}
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	err := fs.WalkDir(templates, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		f, err := templates.Open(path)
		if err != nil {
			return err
		}
		h, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		name := strings.Replace(path, "templates/", "", 1)
		log.Println(name)
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return err
		}
		return nil
	})
	return t, err
}
