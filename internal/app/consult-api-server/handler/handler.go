package handler

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-openapi/runtime/middleware"
	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/restapi/operations/consultation"
)

func ConsultationAddConsultationHandler(params *consultation.AddConsultationParams) middleware.Responder {
	logrus.Debugf("Received: %+v", params.Body)

	return consultation.NewAddConsultationCreated()
}
