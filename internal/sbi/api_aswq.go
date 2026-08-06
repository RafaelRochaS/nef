package sbi

import (
	"net/http"

	"github.com/free5gc/nef/internal/logger"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/metrics/sbi"
	"github.com/gin-gonic/gin"
)

func (s *Server) getAsSessionWithQoSRoutes() []Route {
	return []Route{
		{
			Method:  http.MethodGet,
			Pattern: "/:afID/subscriptions",
			APIFunc: s.apiGetAsSessionWithQoS,
		},
		{
			Method:  http.MethodPost,
			Pattern: "/:afID/subscriptions",
			APIFunc: s.apiPostAsSessionWithQoS,
		},
		{Method: http.MethodGet, Pattern: "/:afID/subscriptions/:subID", APIFunc: s.apiGetAsSessionWithQoSByID},
		{Method: http.MethodPut, Pattern: "/:afID/subscriptions/:subID", APIFunc: s.apiPutAsSessionWithQoSByID},
		{Method: http.MethodPatch, Pattern: "/:afID/subscriptions/:subID", APIFunc: s.apiPatchAsSessionWithQoSByID},
		{Method: http.MethodDelete, Pattern: "/:afID/subscriptions/:subID", APIFunc: s.apiDeleteAsSessionWithQoSByID},
	}
}

func (s *Server) apiGetAsSessionWithQoS(gc *gin.Context) {
	s.Processor().GetAsSessWithQoS(gc, gc.Param("afID"))
}

func (s *Server) apiPostAsSessionWithQoS(gc *gin.Context) {
	var qosPost models.AsSessionWithQoSSubscription
	reqBody, err := gc.GetRawData()
	if err != nil {
		logger.SBILog.Errorf("Get Request Body error: %+v", err)
		pd := openapi.ProblemDetailsSystemFailure(err.Error())
		gc.Set(sbi.IN_PB_DETAILS_CTX_STR, pd.Cause)
		gc.JSON(http.StatusInternalServerError, pd)
		return
	}

	err = openapi.Deserialize(&qosPost, reqBody, "application/json")
	if err != nil {
		logger.SBILog.Errorf("Deserialize Request Body error: %+v", err)
		pd := openapi.ProblemDetailsMalformedReqSyntax(err.Error())
		gc.Set(sbi.IN_PB_DETAILS_CTX_STR, pd.Cause)
		gc.JSON(http.StatusBadRequest, pd)
		return
	}

	s.Processor().PostAsSessWithQoS(
		gc, gc.Param("afID"), qosPost)
}

func (s *Server) apiGetAsSessionWithQoSByID(gc *gin.Context) {
	s.Processor().GetAsSessWithQoSByID(
		gc, gc.Param("afID"), gc.Param("subID"))
}

func (s *Server) apiPutAsSessionWithQoSByID(gc *gin.Context) {
	var qosPut models.AsSessionWithQoSSubscription
	reqBody, err := gc.GetRawData()
	if err != nil {
		logger.SBILog.Errorf("Get Request Body error: %+v", err)
		pd := openapi.ProblemDetailsSystemFailure(err.Error())
		gc.Set(sbi.IN_PB_DETAILS_CTX_STR, pd.Cause)
		gc.JSON(http.StatusInternalServerError, pd)
		return
	}

	err = openapi.Deserialize(&qosPut, reqBody, "application/json")
	if err != nil {
		logger.SBILog.Errorf("Deserialize Request Body error: %+v", err)
		pd := openapi.ProblemDetailsMalformedReqSyntax(err.Error())
		gc.Set(sbi.IN_PB_DETAILS_CTX_STR, pd.Cause)
		gc.JSON(http.StatusBadRequest, pd)
		return
	}

	s.Processor().PutAsSessWithQoSByID(
		gc, gc.Param("afID"), gc.Param("subID"), qosPut)
}

func (s *Server) apiPatchAsSessionWithQoSByID(gc *gin.Context) {
	var qosPatch models.AsSessionWithQoSSubscriptionPatch
	reqBody, err := gc.GetRawData()
	if err != nil {
		logger.SBILog.Errorf("Get Request Body error: %+v", err)
		pd := openapi.ProblemDetailsSystemFailure(err.Error())
		gc.Set(sbi.IN_PB_DETAILS_CTX_STR, pd.Cause)
		gc.JSON(http.StatusInternalServerError, pd)
		return
	}

	err = openapi.Deserialize(&qosPatch, reqBody, "application/json")
	if err != nil {
		logger.SBILog.Errorf("Deserialize Request Body error: %+v", err)
		pd := openapi.ProblemDetailsMalformedReqSyntax(err.Error())
		gc.Set(sbi.IN_PB_DETAILS_CTX_STR, pd.Cause)
		gc.JSON(http.StatusBadRequest, pd)
		return
	}

	s.Processor().PatchAsSessWithQoSByID(
		gc, gc.Param("afID"), gc.Param("subID"), qosPatch)
}

func (s *Server) apiDeleteAsSessionWithQoSByID(gc *gin.Context) {
	s.Processor().DeleteAsSessWithQoSByID(
		gc, gc.Param("afID"), gc.Param("subID"))
}
