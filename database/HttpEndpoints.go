package database

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
	"fmt"
	"github.com/CardFrontendDevopsTeam/GoMongo"
)

type HTTPEndpoint struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string
	Endpoint   string
	Method     string
	Parameters []Parameters
	Threshold  int

	LastChecked time.Time
	LastSuccess time.Time
	ErrorCount  int
	Passing     bool
	Error       string
}

type Parameters struct {
	Name, Value string
}

//AddHTMLEndpoint allows for a new endpoint to be added
func AddHTMLEndpoint(endpoint HTTPEndpoint) {
	c := database.Mongo.C("MonitorHtmlEndpoints")
	c.Insert(endpoint)
}

//GetHTMLEndpoints returns a list of HTML Endpoints
func GetHTMLEndpoints() []HTTPEndpoint {
	c := database.Mongo.C("MonitorHtmlEndpoints")
	q := c.Find(nil)
	i, err := q.Count()
	if err != nil {
		log.Println(err)
		return nil
	}
	r := make([]HTTPEndpoint, i)
	err = q.All(&r)
	if err != nil {
		log.Println(err)
		return nil
	}
	return r
}

/*
SuccessfulEndpointTest will update the mongo element with the ID with the latest details to show it passed successfully
 */
func SuccessfulEndpointTest(endpoint *HTTPEndpoint) error {
	c := database.Mongo.C("MonitorHtmlEndpoints")

	endpoint.LastChecked = time.Now()
	endpoint.LastSuccess = time.Now()
	endpoint.Passing = true
	endpoint.Error = ""
	endpoint.ErrorCount = 0

	err := c.UpdateId(endpoint.ID, endpoint)
	if err != nil {
		return fmt.Errorf("error saving endpoint with success details: %s", err.Error())
	}
	return nil
}

/*
FailedEndpointTest will update the mongo element with the failed details
 */
func FailedEndpointTest(endpoint *HTTPEndpoint, errorMessage string) error {
	c := database.Mongo.C("MonitorHtmlEndpoints")
	result := HTTPEndpoint{}
	err := c.FindId(endpoint.ID).One(&result);
	if err != nil {
		return fmt.Errorf("error retreiving endpoint with success details: %s", err.Error())
	}

	result.LastChecked = time.Now()
	result.Passing = false
	result.Error = errorMessage
	result.ErrorCount++

	err = c.UpdateId(endpoint.ID, result)
	if err != nil {
		return fmt.Errorf("error saving endpoint with success details: %s", err.Error())
	}
	return nil
}
