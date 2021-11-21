package password

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	random "math/rand"
	"net/http"
	"strconv"
	"time"
)

type Password struct {
	Password string ` json:"password" xml:"password" `
	Length   int    ` json:"Length" xml:"Length" `
}

func api(c *gin.Context) {
	b := baseData{}
	// Show some documentation here
	b.Action = "Api.Index"
	b.Date = time.Now().Format("2006")
	b.Title = "Password.fun"

	c.HTML(http.StatusOK, "api/index.html", b)
}

type passwordForm struct {
	Count     int  `form:"count" binding:"required"`
	Length    int  `form:"length" binding:"required"`
	Capital   bool `form:"capital"`
	Lower     bool `form:"lower"`
	Special   bool `form:"special"`
	Spaces    bool `form:"spaces"`
	Numbers   bool `form:"numbers"`
	Highlight bool `form:"highlight"`
	Remember  bool `form:"remember"`
}

func passwords(c *gin.Context) {
	form := passwordForm{}
	err := c.Bind(&form)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Invalid form data")
		return
	}

	s := sessions.Default(c)
	if form.Remember {
		s.Set("Count", intToString(form.Count))
		s.Set("Length", intToString(form.Length))
		s.Set("Capital", boolToString(form.Capital))
		s.Set("Lower", boolToString(form.Lower))
		s.Set("Special", boolToString(form.Special))
		s.Set("Spaces", boolToString(form.Spaces))
		s.Set("Numbers", boolToString(form.Numbers))
		s.Set("Highlight", boolToString(form.Highlight))
		s.Set("Remember", boolToString(form.Remember))
	} else {
		s.Delete("Count")
		s.Delete("length")
		s.Delete("Capital")
		s.Delete("Lower")
		s.Delete("Special")
		s.Delete("Spaces")
		s.Delete("Numbers")
		s.Delete("Highlight")
		s.Delete("Remember")
	}

	chars := ""
	if form.Capital {
		chars = chars + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if form.Lower {
		chars = chars + "abcdefghijklmnopqrstuvwxyz"
	}
	if form.Special {
		chars = chars + "!@#$%^&*()-_=+,.?/:;{}[]`~"
	}
	if form.Spaces {
		chars = chars + " "
	}
	if form.Numbers {
		chars = chars + "1234567890"
	}

	data := make(map[string]interface{})
	data["error"] = nil
	var length = 25

	// We don't allow passwords over 100 in length
	if form.Length > 100 {
		form.Length = 100
	}
	if form.Length > 0 {
		length = form.Length
	}

	var bytes = []byte(chars)
	if len(bytes) <= 0 {
		bytes = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()-_=+,.?/:;{}[]`~ 1234567890")
	}
	passwords := []Password{
		{Password: rand_char(length, bytes), Length: length},
	}

	var numpass = 5

	// What would you possibly need 100 passwords for?
	if form.Count > 100 {
		form.Count = 100
	}
	if form.Count > 0 {
		numpass = form.Count - 1
	}

	for i := 0; i < numpass; i++ {
		length = random.Intn(50)
		if form.Length > 0 {
			length = form.Length
		}
		passwords = append(passwords, Password{Password: rand_char(length, bytes), Length: length})
	}
	data["href"] = "https://password.fun" + "/api/v1/passwords"
	data["passwords"] = passwords
	data["Count"] = len(passwords)
	c.JSON(http.StatusOK, data)
}

func rand_char(length int, chars []byte) string {
	// Some ideas and credit for this function is given to https://github.com/cmiceli/password-generator-go
	newPword := make([]byte, length)
	randomData := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPword)
			}
		}
	}
}

func intToString(b int) string {
	return fmt.Sprintf("%d", b)
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}
