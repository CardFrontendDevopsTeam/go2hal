package database

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
)

type HTTPEndpoint struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string
	Endpoint   string
	Method     string
	Parameters []Parameters

	lastChecked time.Time
	lastSuccess time.Time
	Passing     bool
	Error       string
}

type Parameters struct {
	Name, Value string
}

//AddHTMLEndpoint allows for a new endpoint to be added
func AddHTMLEndpoint(endpoint HTTPEndpoint) {
	c := database.C("MonitorHtmlEndpoints")
	c.Insert(endpoint)
}

//GetHTMLEndpoints returns a list of HTML Endpoints
func GetHTMLEndpoints() []HTTPEndpoint {
	c := database.C("MonitorHtmlEndpoints")
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
func SuccessfulEndpointTest(id string) {
	c := database.C("MonitorHtmlEndpoints")
	result := HTTPEndpoint{}
	err := c.FindId(id).One(&result);
	if err != nil {
		log.Printf("Error retreiving endpoint with success details: %s", err.Error())
		return
	}

	result.lastChecked = time.Now()
	result.lastSuccess = time.Now()
	result.Passing = true
	result.Error = ""

	err = c.UpdateId(id, result)
	if err != nil {
		log.Printf("Error saving endpoint with success details: %s", err.Error())
	}
}

/*
FailedEndpointTest will update the mongo element with the failed details
 */
func FailedEndpointTest(endpoint HTTPEndpoint, errorMessage string) {
	c := database.C("MonitorHtmlEndpoints")
	result := HTTPEndpoint{}
	err := c.FindId(endpoint.ID).One(&result);
	if err != nil {
		log.Printf("Error retreiving endpoint with success details: %s", err.Error())
		return
	}

	result.lastChecked = time.Now()
	result.Passing = false
	result.Error = errorMessage

	err = c.UpdateId(endpoint.ID, result)
	if err != nil {
		log.Printf("Error saving endpoint with success details: %s", err.Error())
	}
}
