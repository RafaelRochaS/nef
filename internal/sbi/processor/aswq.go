package processor

import (
	"github.com/free5gc/nef/internal/logger"
	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
)

// PostAsSessWithQoS creates a new AS session with QoS subscription
// 3GPP TS 29.122 release 17 version 17.6.0
// Resource structure: 5.14.3
// Request/Response: 5.14.3.2.3.4
func (p *Processor) PostAsSessWithQoS(c *gin.Context, scsAsID string, qosPost models.AsSessionWithQoSSubscription) {
	logger.AsSessQoSLog.Infof("PostAsSessWithQoS - scsAsID: %s", scsAsID)
}

func (p *Processor) GetAsSessWithQoS(c *gin.Context, scsAsID string) {
	logger.AsSessQoSLog.Infof("GetAsSessWithQoS - scsAsID: %s", scsAsID)
}

func (p *Processor) GetAsSessWithQoSByID(c *gin.Context, scsAsID, subID string) {
	logger.AsSessQoSLog.Infof("GetAsSessWithQoSByID - scsAsID: %s, subID: %s", scsAsID, subID)
}

func (p *Processor) PutAsSessWithQoSByID(
	c *gin.Context, scsAsID, subID string, qosPut models.AsSessionWithQoSSubscription,
) {
	logger.AsSessQoSLog.Infof("PutAsSessWithQoSByID - scsAsID: %s, subID: %s", scsAsID, subID)
}

func (p *Processor) PatchAsSessWithQoSByID(
	c *gin.Context, scsAsID, subID string, qosPatch models.AsSessionWithQoSSubscriptionPatch,
) {
	logger.AsSessQoSLog.Infof("PatchAsSessWithQoSByID - scsAsID: %s, subID: %s", scsAsID, subID)
}

func (p *Processor) DeleteAsSessWithQoSByID(c *gin.Context, scsAsID, subID string) {
	logger.AsSessQoSLog.Infof("DeleteAsSessWithQoSByID - scsAsID: %s, subID: %s", scsAsID, subID)
}
