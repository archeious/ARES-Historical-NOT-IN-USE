package controllers

import (
	"encoding/json"
	"github.com/mrjones/oauth"
	"github.com/revel/revel"
	"io/ioutil"
	"webfrontend/app/models"
)

var TWITTER = oauth.NewConsumer(
	"xNnVVJhP0bnBcZdU9dVQ09xh0",
	"hiL0E7ETAwPvC54MkOQUPTdJ7HEMNisTgk64bXmCU7JmqAio00",
	oauth.ServiceProvider{
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	},
)

type Account struct {
	App
}

func (c Account) Index() revel.Result {
	user := getUser()
	if user.AccessToken == nil {
		return c.Render()
	}

	resp, err := TWITTER.Get(
		"https://api.twitter.com/1.1/statuses/mentions_timeline.json",
		map[string]string{"count": "10"},
		user.AccessToken)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Render()
	}
	defer resp.Body.Close()

	mentions := []struct {
		Text string `json:test`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&mentions)
	if err != nil {
		revel.ERROR.Println(err)
	}
	revel.INFO.Println(mentions)
	return c.Render(mentions)
}

func (c Account) SetStatus(status string) revel.Result {
	resp, err := TWITTER.PostForm(
		"http://api.twitter.com/1.1/statuses/update.json",
		map[string]string{"status": status},
		getUser().AccessToken,
	)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	revel.INFO.Println(string(bodyBytes))
	c.Response.ContentType = "application/json"
	return c.RenderText(string(bodyBytes))
}

func (c Account) Authenticate(oauth_verifier string) revel.Result {
	user := getUser()
	if oauth_verifier != "" {
		accessToken, err := TWITTER.AuthorizeToken(user.RequestToken, oauth_verifier)
		if err == nil {
			user.AccessToken = accessToken
		} else {
			revel.ERROR.Println("Error connecting to twitter:", err)
		}
		return c.Redirect(Account.Index)
	}

	requestToken, url, err := TWITTER.GetRequestTokenAndUrl("http://ares.datistry.com/user/auth/twitter")
	if err == nil {
		user.RequestToken = requestToken
		return c.Redirect(url)
	} else {
		revel.ERROR.Println("Error connecting to twitter:", err)
	}
	return c.Redirect(Account.Index)
}

func getUser() *models.User {
	return models.FindOrCreate("guest")
}

func init() {
	TWITTER.Debug(true)
}
