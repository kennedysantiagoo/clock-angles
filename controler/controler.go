package controler

import (
	"log"
	"net/http"

	"github.com/kennedysantiagoo/service"

	"github.com/kennedysantiagoo/helper"

	"github.com/gorilla/mux"
)

func GetAngle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	stringHour, stringMinute := helper.GetParams(vars)

	hour, minute := helper.ConvertParamsToInt(stringHour, stringMinute)

	angle, err := service.GetAngle(hour, minute)

	var responseBody helper.ResponseBody

	if err != nil {
		log.Println("An error ocurred while getting angle ", err)
		responseBody = helper.BuildResponseBody("An error ocurred while getting clock angle", false, 0)
		helper.WriteResponse(w, http.StatusInternalServerError, responseBody)
	}

	responseBody = helper.BuildResponseBody("Succesfully finished", true, angle)
	helper.WriteResponse(w, http.StatusOK, responseBody)

}
