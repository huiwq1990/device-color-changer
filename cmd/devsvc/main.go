package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	// allow CORS request for frontend dev server request
	router.Use(cors.Default())

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("frontend/", true)))
	// if no match route found, take it to index.html for front-end routing
	router.NoRoute(func(c *gin.Context) { c.File("frontend/index.html") })

	// setup route group for the API
	setAPIGroup(router)
	// setup logger output to stdout
	//setLogger(os.Stdout)

	// Start and run the server
	if err := router.Run(":5000"); err != nil {
		panic("gin server fail to start")
	}
}

func setColorGroup(router *gin.Engine) {



	//api := router.Group("/api")
	//{
	//	api.POST("/getOverview", getClusterOverviewHandler)
	//	api.POST("/getNodes", getNodeHandler)
	//	api.POST("/getDeployments", getDeploymentHandler)
	//	api.POST("/getStatefulsets", getStatefulsetHandler)
	//	api.POST("/getJobs", getJobHandler)
	//	api.POST("/getNodepools", getNodepoolHandler)
	//	api.POST("/getPods", getPodHandler)
	//	api.POST("/login", loginHandler)
	//	api.POST("/register", registerHandler)
	//	api.POST("/setNodeAutonomy", setNodeAutonomyHandler)
	//	api.POST("/getApps", getAppHandler)
	//	api.POST("/installApp", installAppHandler)
	//	api.POST("/uninstallApp", uninstallAppHandler)
	//	api.POST("/github", githubLoginHandler)
	//}
}

type color struct {
	Color string `json:"color"`
}

var curColor = "green"

func setAPIGroup(router *gin.Engine) {



	api := router.Group("/")
	{
		api.GET("/_ajaxAutoRefresh", refresh)
		api.POST("/api/v1/device/register", register)
		api.POST("/api/v1/device/:id/changeColor", change)
	}
}

func refresh(c *gin.Context) {

	c.JSON(http.StatusOK,color{
		Color: curColor,
	})
}




func register(c *gin.Context) {

	body := c.Request.Body
	defer body.Close()
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Errorf("read body failed! err: %v", err)
	}



	fmt.Printf("%s",string(bytes))
fmt.Println()
	c.JSON(http.StatusOK,"register succ")
}


//container =



func change(c *gin.Context) {
	id := c.Param("id")

	body := c.Request.Body
	defer body.Close()
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Errorf("read body failed! err: %v", err)
	}

	val := color{}
	if err := json.Unmarshal(bytes,&val);err != nil {
		fmt.Errorf("parse %v",err)
	}



	fmt.Printf("%s %s %s",id,curColor,string(bytes))
	curColor = val.Color
	fmt.Printf("current %s",curColor)
	c.JSON(http.StatusOK,"register succ")
}





