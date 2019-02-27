package handler

import (
	"fmt"
	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/models"
	"testing"
	"time"
)

func TestDBHandler_doConsultationAddConsultationHandler(t *testing.T) {
	underTestHandler := NewDBHandler("localhost", "TEST-PR00B121")

	testTask := fmt.Sprintf("test-%s", time.Now())

	body := &models.ConsultationReq{
		Task: testTask,
	}
	result := underTestHandler.doConsultationAddConsultationHandler(body)
	fmt.Printf("result: %+v\n", result)

	//assert.True(t, result.t, fmt.Sprintf("Neturetu buti klaida: %s", err))

}
