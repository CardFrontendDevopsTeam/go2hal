package chef

import (
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/weAutomateEverything/go2hal/gokit"
	"github.com/weAutomateEverything/go2hal/machineLearning"
	"github.com/weAutomateEverything/go2hal/telegram"
	"net/http"
)

//MakeHandler returns a restful http handler for the chef delivery service
//the Machine Learning service can be set to nil if you do not wish to log the http requests

func MakeHandler(service Service, logger kitlog.Logger, ml machineLearning.Service) http.Handler {
	opts := gokit.GetServerOpts(logger, ml)

	chefDeliveryEndpoint := kithttp.NewServer(makeChefDeliveryAlertEndpoint(service), gokit.DecodeString, gokit.EncodeResponse, opts...)

	addChefRecipeToGroup := kithttp.NewServer(gokitjwt.NewParser(gokit.GetJWTKeys(), jwt.SigningMethodHS256,
		telegram.CustomClaimFactory)(makeAddRecipeToGroupEndpoint(service)), decodeAddChefRequest, gokit.EncodeResponse, opts...)

	getChefRecipesForGroup := kithttp.NewServer(gokitjwt.NewParser(gokit.GetJWTKeys(), jwt.SigningMethodHS256,
		telegram.CustomClaimFactory)(makeGetAllGrouRecipesEndpoint(service)), gokit.DecodeString, gokit.EncodeResponse, opts...)

	addEnvironmentToGroup := kithttp.NewServer(gokitjwt.NewParser(gokit.GetJWTKeys(), jwt.SigningMethodHS256,
		telegram.CustomClaimFactory)(makeAddEnvironmentToGroupEndpoint(service)), decodeAddEnvironmentfRequest, gokit.EncodeResponse, opts...)

	getEnvironmentForGroup := kithttp.NewServer(gokitjwt.NewParser(gokit.GetJWTKeys(), jwt.SigningMethodHS256,
		telegram.CustomClaimFactory)(makeGetEnvironmentForGroupEndpoint(service)), gokit.DecodeString, gokit.EncodeResponse, opts...)

	r := mux.NewRouter()

	// swagger:operation POST /chef/delivery/{group} sendDeliveryMessage
	//
	// Sends a delivery notifcation to the chat group
	// //
	// ---
	// produces:
	// parameters:
	// responses:
	//   '200':
	//     description: Message has been sent
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorModel"
	r.Handle("/chef/delivery/{chatid:[0-9]+}", chefDeliveryEndpoint).Methods("POST")
	// swagger:operation GET /pets getPet
	//
	// Returns all pets from the system that the user has access to
	//
	// Could be any pet
	//
	// ---
	// produces:
	// - application/json
	// - application/xml
	// - text/xml
	// - text/html
	// parameters:
	// - name: tags
	//   in: query
	//   description: tags to filter by
	//   required: false
	//   type: array
	//   items:
	//     type: string
	//   collectionFormat: csv
	// - name: limit
	//   in: query
	//   description: maximum number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// responses:
	//   '200':
	//     description: pet response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/pet"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorModel"
	r.Handle("/chef/recipe", addChefRecipeToGroup).Methods("POST")
	// swagger:operation GET /pets getPet
	//
	// Returns all pets from the system that the user has access to
	//
	// Could be any pet
	//
	// ---
	// produces:
	// - application/json
	// - application/xml
	// - text/xml
	// - text/html
	// parameters:
	// - name: tags
	//   in: query
	//   description: tags to filter by
	//   required: false
	//   type: array
	//   items:
	//     type: string
	//   collectionFormat: csv
	// - name: limit
	//   in: query
	//   description: maximum number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// responses:
	//   '200':
	//     description: pet response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/pet"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorModel"
	r.Handle("/chef/recipes", getChefRecipesForGroup).Methods("GET")
	// swagger:operation GET /pets getPet
	//
	// Returns all pets from the system that the user has access to
	//
	// Could be any pet
	//
	// ---
	// produces:
	// - application/json
	// - application/xml
	// - text/xml
	// - text/html
	// parameters:
	// - name: tags
	//   in: query
	//   description: tags to filter by
	//   required: false
	//   type: array
	//   items:
	//     type: string
	//   collectionFormat: csv
	// - name: limit
	//   in: query
	//   description: maximum number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// responses:
	//   '200':
	//     description: pet response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/pet"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorModel"
	r.Handle("/chef/environment", addEnvironmentToGroup).Methods("POST")
	// swagger:operation GET /pets getPet
	//
	// Returns all pets from the system that the user has access to
	//
	// Could be any pet
	//
	// ---
	// produces:
	// - application/json
	// - application/xml
	// - text/xml
	// - text/html
	// parameters:
	// - name: tags
	//   in: query
	//   description: tags to filter by
	//   required: false
	//   type: array
	//   items:
	//     type: string
	//   collectionFormat: csv
	// - name: limit
	//   in: query
	//   description: maximum number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// responses:
	//   '200':
	//     description: pet response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/pet"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorModel"
	r.Handle("/chef/environments", getEnvironmentForGroup).Methods("GET")

	return r

}

func decodeAddChefRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var q = &addRecipeRequest{}
	err := json.NewDecoder(r.Body).Decode(&q)
	return q, err
}

func decodeAddEnvironmentfRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var q = &addEnvironmentRequest{}
	err := json.NewDecoder(r.Body).Decode(&q)
	return q, err
}
