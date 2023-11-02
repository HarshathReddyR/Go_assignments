package handlers

import (
	"encoding/json"
	"jobnew/internal/middlewares"
	"jobnew/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *handler) JobCreate(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var nj models.NewJob
	err := json.NewDecoder(c.Request.Body).Decode(&nj)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(nj)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide proper job detials"})
		return
	}

	id := c.Param("cid")
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	cmpny, err := h.s.CreateJob(ctx, nj, cid)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "job creation failed"})
		return
	}
	c.JSON(http.StatusOK, cmpny)
}
func (h *handler) JobView(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceID missing from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	job, err := h.s.ViewJob(ctx)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("problem in fetching job")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide proper job detials"})
		return
	}
	c.JSON(http.StatusOK, job)
}
func (h *handler) JobviewById(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceID missing from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	stringjobid := c.Param("id")
	jid, err := strconv.ParseUint(stringjobid, 10, 64)
	if err != nil {
		log.Print("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error founr at conversion"})
		return
	}
	job, err := h.s.ViewJobByID(jid)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("job is not there with that id ")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "job fetching failed"})
		return
	}
	c.JSON(http.StatusOK, job)
}
func (h *handler) JobViewByCompanyId(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceID missing from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	stringCompanyid := c.Param("id")
	cid, err := strconv.ParseUint(stringCompanyid, 10, 64)
	if err != nil {
		log.Print("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error founr at conversion"})
		return
	}
	job, err := h.s.ViiewJobByCompany(cid)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("job is not there with that id ")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "job fetching failed"})
		return
	}
	c.JSON(http.StatusOK, job)
}
