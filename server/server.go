package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Server Starts the server for the client side
func Server(port string) error {

	r := gin.Default()

	r.Static("/assets", "./templates/assets")

	r.LoadHTMLGlob("./templates/*.html")
	// htmlRender.TemplatesDir = "templates/" // default
	// htmlRender.Ext = ".html"               // default

	// Tell gin to use our html render

	// Initialise connection to Sqlite DB
	connect, err := Connect()
	if err != nil {
		return err
	}

	// Gets default information about the game sessions
	r.GET("/", func(c *gin.Context) {
		// Get client session ID
		session, err := c.Cookie("SessionID")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
		}

		// User information based on session ID
		User, err := GetUserInformation(connect, session)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
		}

		// Get information of all game sessions available
		GameSessions, err := DisplayGameSessions(connect)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":        "Xplane11-WebRTC",
			"GameSessions": GameSessions,
			"User":         User,
		})
	})

	// Opens the forgot password page
	r.GET("/forgotpassword", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forgotpassword.html", gin.H{
			"title": "Xplane11-WebRTC",
		})
	})

	// Opens the login page
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Xplane11-WebRTC",
		})
	})

	// Validate login information
	r.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		// source: https://github.com/gin-gonic/gin#multiparturlencoded-binding
		var users Users
		// in this case proper binding will be automatically selected
		if err := c.ShouldBind(&users); err != nil {
			c.String(http.StatusBadRequest, "bad request")
		}

		result, err := AuthLogin(connect, &users)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		} else {
			if result == "success" {
				// Create Login session
				session, err := CreateLoginSession(connect, &users)
				if err != nil {
					c.String(http.StatusExpectationFailed, "error")
				}
				// TODO redirect to homepage
				// Set client Cookie as Session ID
				// To be changed when we use HTTPS
				c.SetCookie("SessionID", session.SessionKey, 3600, "/", "localhost", false, true)
				// redirects to the home page
				//c.Redirect(http.StatusFound, "/")

				// Sends message success. which then redirects back to the home page
				c.String(http.StatusOK, "Success!")
			} else {
				c.String(http.StatusOK, "Something is wrong with from our side. Email us: me@akilan.io to find out more. ")
			}
		}
	})

	// Route for logout
	r.GET("/logout", func(c *gin.Context) {
		// Get client session ID
		session, err := c.Cookie("SessionID")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
		}

		// Remove Session ID from the database
		RemoveLoginSession(connect, session)

		// redirect to the login page
		c.Redirect(http.StatusFound, "/login")
	})

	// Opens the registration page
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registration.html", gin.H{
			"title": "Xplane11-WebRTC",
		})
	})

	// Register user information required
	r.POST("/register", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		// source: https://github.com/gin-gonic/gin#multiparturlencoded-binding
		var users Users
		// in this case proper binding will be automatically selected
		if err := c.ShouldBind(&users); err != nil {
			c.String(http.StatusBadRequest, "bad request")
		}

		err = RegisterUser(connect, &users)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		} else {
			c.String(http.StatusOK, "Success! Head back to the login page")
		}
	})

	// Add new server information
	r.POST("/AddGameSession", func(c *gin.Context) {

		// TODO: Binding
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		// source: https://github.com/gin-gonic/gin#multiparturlencoded-binding
		var gameSession GameSession

		gameSession.Rate, _ = strconv.ParseFloat(c.PostForm("Rate"), 8)
		gameSession.Server.GPU = c.PostForm("GPU")
		gameSession.Link = c.PostForm("Link")
		gameSession.Server.CPU = c.PostForm("CPU")
		gameSession.Server.GPU = c.PostForm("GPU")
		gameSession.Server.Platform = c.PostForm("Platform")
		gameSession.Server.Hostname = c.PostForm("HostName")
		gameSession.Server.RAM = c.PostForm("RAM")
		gameSession.Server.Disk = c.PostForm("Disk")

		// in this case proper binding will be automatically selected
		//if err := c.ShouldBind(&gameSession); err != nil {
		//	c.String(http.StatusBadRequest, "bad request")
		//}

		err = AddServerSpecs(connect, &gameSession)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		} else {
			c.String(http.StatusOK, "Success")
		}
	})

	// Run gin server on the specified port
	err = r.Run(":" + port)
	if err != nil {
		return err
	}

	return nil
}