package app

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

var DB *sql.DB

func initDB() {
	revel.INFO.Println("DB Loading")

	var err error
	DB, err = sql.Open("sqlite3", "anime.sql3")
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		sessionFilter,
		revel.InterceptorFilter, // Run interceptors around the action.
		revel.CompressFilter,    // Compress the result.
		revel.ActionInvoker,     // Invoke the action.
	}
	revel.OnAppStart(initDB)

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(FillCache)
}

var sessionFilter = func(c *revel.Controller, fc []revel.Filter) {
	if user, ok := c.Session["user"]; ok {
		c.RenderArgs["user"] = user
	}

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
