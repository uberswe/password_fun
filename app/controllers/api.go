package controllers

import (
	"github.com/revel/revel"
	"io"
	"crypto/rand"
	random "math/rand"
	"net/url"
	"fmt"
	"strconv"
	"time"
)

type Api struct {
	*revel.Controller
}

type Password struct {
	Password string ` json:"password" xml:"password" `
	Length int ` json:"length" xml:"length" `
}

type Params struct {
	url.Values
}

func (c Api) Index() revel.Result {
	// Show some documentation here
	var action string = c.Action
	var date string = time.Now().Format("2006")
	var title string = "Password.fun"
	return c.Render(action, title, date)
}

func (c Api) Passwords() revel.Result {
	var paramCount int
	var paramLength int
	var paramCapital bool
	var paramLower bool
	var paramSpecial bool
	var paramSpaces bool
	var paramNumbers bool
	var paramRemember bool
	var paramHighlight bool
	c.Params.Bind(&paramCount, "count") // Sets the number of passwords
	c.Params.Bind(&paramLength, "length") // Sets the length of the password(s)
	c.Params.Bind(&paramCapital, "capital") // Should contain capital letters?
	c.Params.Bind(&paramLower, "lower")
	c.Params.Bind(&paramSpecial, "special")
	c.Params.Bind(&paramSpaces, "spaces")
	c.Params.Bind(&paramNumbers, "numbers")
	c.Params.Bind(&paramRemember, "remember") // Store values in session cookie
	c.Params.Bind(&paramHighlight, "highlight") // Highlight random string with JS

	if (paramRemember) {
		c.Session["count"] = intToString(paramCount);
		c.Session["length"] = intToString(paramLength);
		c.Session["capital"] = boolToString(paramCapital);
		c.Session["lower"] = boolToString(paramLower);
		c.Session["special"] = boolToString(paramSpecial);
		c.Session["spaces"] = boolToString(paramSpaces);
		c.Session["numbers"] = boolToString(paramNumbers);
		c.Session["highlight"] = boolToString(paramHighlight);
		c.Session["remember"] = boolToString(paramRemember);
	} else {
		delete(c.Session, "count")
		delete(c.Session, "length")
		delete(c.Session, "capital")
		delete(c.Session, "lower")
		delete(c.Session, "special")
		delete(c.Session, "spaces")
		delete(c.Session, "numbers")
		delete(c.Session, "highlight")
		delete(c.Session, "remember")
	}

	var chars string = ""
	if (paramCapital) {

		chars = chars + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if (paramLower) {
		chars = chars + "abcdefghijklmnopqrstuvwxyz"
	}
	if (paramSpecial) {
		chars = chars + "!@#$%^&*()-_=+,.?/:;{}[]`~"
	}
	if (paramSpaces) {
		chars = chars + " "
	}
	if (paramNumbers) {
		chars = chars + "1234567890"
	}

	data := make(map[string]interface{})
	data["error"] = nil
	var length = random.Intn(50)
	if (length < 5) {
		length = 5
	}

	// We don't allow passwords over 100 in length
	if (paramLength > 100) {
		paramLength = 100
	}
	if (paramLength > 0) {
		length = paramLength
	}
	var bytes = []byte(chars)
	if (len(bytes) <= 0) {
		bytes = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()-_=+,.?/:;{}[]`~ 1234567890")
	}
	passwords := []Password{
		{Password: rand_char(length, bytes), Length: length},
	}
	var numpass = random.Intn(50)
	if (numpass < 5) {
		numpass = 5
	}
	// What would you possibly need 100 passwords for?
	if (paramCount > 100) {
		paramCount = 100
	}
	if (paramCount > 0) {
		numpass = paramCount-1
	}
	for i := 0; i < numpass; i++ {
		length = random.Intn(50)
		if (paramLength > 0) {
			length = paramLength
		}
		passwords = append(passwords,Password{Password: rand_char(length, bytes), Length: length})
	}
	data["href"] = "https://lösenord.xyz" + "/api/v1/passwords"
	data["passwords"] = passwords
	data["count"] = len(passwords)
	return c.RenderJSON(data)
	//return c.RenderXML(data)
}

func rand_char(length int, chars []byte) string {
	// Some ideas and credit for this function is given to https://github.com/cmiceli/password-generator-go
	new_pword := make([]byte, length)
	random_data := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, random_data); err != nil {
			panic(err)
		}
		for _, c := range random_data {
			if c >= maxrb {
				continue
			}
			new_pword[i] = chars[c%clen]
			i++
			if i == length {
				return string(new_pword)
			}
		}
	}
	panic("unreachable")
}

func intToString(b int) string {
	return fmt.Sprintf("%d", b)
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}