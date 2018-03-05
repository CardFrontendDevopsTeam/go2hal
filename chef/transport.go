package chef

import (
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/zamedic/go2hal/gokit"
	"github.com/zamedic/go2hal/machineLearning"
	"net/http"
)

func MakeHandler(service Service, logger kitlog.Logger, ml machineLearning.Service) http.Handler {
	opts := gokit.GetServerOpts(logger, ml)

	chefDeliveryEndpoint := kithttp.NewServer(makeChefDeliveryAlertEndpoint(service), gokit.DecodeString, gokit.EncodeResponse, opts...)

	r := mux.NewRouter()

	r.Handle("/delivery", chefDeliveryEndpoint).Methods("POST")

	return r

}
