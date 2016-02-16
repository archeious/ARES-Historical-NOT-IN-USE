package controllers

import (
	"github.com/revel/revel"
	"webfrontend/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User {
	//	if c.RenderArgs["user"] != nil {
	//		return c.RenderArgs["user"].(*models.User)
	//	}
	if username, ok := c.Session["user"]; ok {
		return models.GetUser(username)
	}
	return nil
}

func (c App) Login(username, password string) revel.Result {
	if username != "" {
		user := models.GetUser(username)
		if user != nil {
			if user.VerifyPassword(password) {
				c.Session["user"] = username
				c.Flash.Success("Welcome, " + username)
				return c.Redirect("/")
			}
		}
		c.Flash.Out["username"] = username
		c.Flash.Error("Login Failed")
		return c.Redirect("/")
	}
	return c.Render()
}

func (c App) Logout() revel.Result {
	for i := range c.Session {
		delete(c.Session, i)
	}
	return c.Redirect("/")
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) List() revel.Result {
	series := models.GetAllSeries()
	return c.Render(series)
}
