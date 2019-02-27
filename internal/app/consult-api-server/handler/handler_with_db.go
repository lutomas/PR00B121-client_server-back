package handler

import (
	"github.com/Sirupsen/logrus"
	"github.com/globalsign/mgo"
	"github.com/go-openapi/runtime/middleware"
	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/models"
	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/restapi/operations/consultation"
)

const collectionConsultancy = "consultations"

type DBHandler struct {
	host   string
	dbName string
}

func NewDBHandler(host, dbName string) *DBHandler {
	return &DBHandler{
		host:   host,
		dbName: dbName,
	}
}

//
func (x *DBHandler) ConsultationAddConsultationHandler(params *consultation.AddConsultationParams) middleware.Responder {
	logrus.Debugf("Received: %v \n", params.Body)

	//get session
	session, err := mgo.Dial(x.host)
	if err != nil {
		errorResponse := &models.Error{
			Reason: int64(DbAccessError),
			Msg:    err.Error(),
		}
		return consultation.NewAddConsultationBadRequest().WithPayload(errorResponse)
	}

	defer session.Close()

	//error check on every access
	session.SetSafe(&mgo.Safe{})

	//get collection
	collection := session.DB(x.dbName).C(collectionConsultancy)

	err = collection.Insert(params.Body)
	if err != nil {
		errorResponse := &models.Error{
			Reason: int64(DbInsertError),
			Msg:    err.Error(),
		}
		return consultation.NewAddConsultationBadRequest().WithPayload(errorResponse)
	}

	return consultation.NewAddConsultationCreated()
}

//TODO: Comment-Out 6
//func (x *DBHandler) ConsultationGetConsultationsHandler(params *consultation.GetConsultationsParams) middleware.Responder {
//	errorResponse := &models.Error{
//		Reason: int64(NotImplemented),
//		Msg:    "Nerealizuota",
//	}
//	return consultation.NewGetConsultationsBadRequest().WithPayload(errorResponse)
//}

//TODO: Uncomment 6
func (x *DBHandler) ConsultationGetConsultationsHandler(params *consultation.GetConsultationsParams) middleware.Responder {
	//get session
	session, err := mgo.Dial(x.host)
	if err != nil {
		errorResponse := &models.Error{
			Reason: int64(DbAccessError),
			Msg:    err.Error(),
		}
		return consultation.NewAddConsultationBadRequest().WithPayload(errorResponse)
	}

	defer session.Close()

	//error check on every access
	session.SetSafe(&mgo.Safe{})

	//get collection
	collection := session.DB(x.dbName).C(collectionConsultancy)

	var result []*models.ConsultationResp
	if err = collection.Find(nil).All(&result); err != nil {
		errorResponse := &models.Error{
			Reason: int64(DbReadError),
			Msg:    err.Error(),
		}
		consultation.NewGetConsultationsBadRequest().WithPayload(errorResponse)
	}

	return consultation.NewGetConsultationsOK().WithPayload(result)
}
