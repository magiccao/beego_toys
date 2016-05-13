package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	filename := path.Join("static/", c.Ctx.Request.URL.Path)
	if filename == "static" || filename == "static/" {
		filename = "static/index.html"
	}
	f, e := os.Open(filename)
	if e != nil {
		beego.Error("open index.html err:%v", e.Error())
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("server internal error")
		return
	}

	if strings.HasSuffix(filename, ".css") {
		c.Ctx.Output.Header("Content-Type", "text/css")
	} else if strings.HasSuffix(filename, ".js") {
		c.Ctx.Output.Header("Content-Type", "application/x-javascript; charset=UTF-8")
	}

	_, e = io.Copy(c.Ctx.ResponseWriter, f)
	if e != nil {
		beego.Error("write client err: %v", e.Error())
		return
	}
}

func (c *MainController) Statistics() {
    c.Ctx.Request.ParseForm()
    name := c.Ctx.Request.Form.Get("qtalk")
    department := c.Ctx.Request.Form.Get("department")
    office := c.Ctx.Request.Form.Get("office")
    sex := c.Ctx.Request.Form.Get("size")

    size := 0
    if sex == "female" {
        size = 1
    }

    user := &models.User{Name:name}
    user.Read("name")
    
    user.Size = size
    user.Department = department
    user.Office = office

    if user.Id != 0 {
        user.Update()
    } else {
        err := user.Insert()
        if err != nil {
            beego.Warn("insert user err: ", err.Error())
        }
    }

    c.Ctx.Output.SetStatus(302)
    c.Ctx.Output.Header("Location", "/index.html")
}

func (c *MainController) Show() {
	users, err := models.ReadUsers()
	if err != nil {
		beego.Warn("read user table err: ", err.Error())
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	html := `
        <html>
        <head>
            <style type="text/css">
            table {margin-left: 50px; width: 450px;}
            th {width: 100px; background-color: #aaa; border: bold, 2px, #999; font-weight: bold;}
            td {text-align: center; background-color: #ccc; color: #111;}
            </style>
        </head>
        <body>
            <table>
                <tr>
                    <th>id</th>
                    <th>qtalk</th>
                    <th>department</th>
                    <th>office</th>
                    <th>style</th>
                </tr>
    `

	for i, user := range users {
        sex := "男装"
        if user.Size == 1 {
            sex = "女装"
        }
        html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", 
            i+1, user.Name, user.Department, user.Office, sex)
	}

	html += `
            </table>
        </body>
        </html>
    `

	c.Ctx.WriteString(html)
}
