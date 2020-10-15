package controllers
import (
   "context"
   "fmt"
   "strconv"
   "time"
   
   "github.com/konrawitAEK/app/ent"
   "github.com/konrawitAEK/app/ent/positionassingment"
   "github.com/konrawitAEK/app/ent/physician"
   "github.com/konrawitAEK/app/ent/department"
   "github.com/konrawitAEK/app/ent/position"
   "github.com/gin-gonic/gin"
) 
// PositionassingmentController defines the struct for the positionassingment controller
type PositionassingmentController struct {
   client *ent.Client
   router gin.IRouter
}

type Positionassingment struct{
	Physicianid    int
	Departmentid   int
	Positionid     int
	DayStart       string
}

// CreatePositionassingment handles POST requests for adding positionassingment entities
// @Summary Create positionassingment  
// @Description Create positionassingment  
// @ID create-positionassingment  
// @Accept   json
// @Produce  json
// @Param positionassingment body Positionassingment true "Positionassingment entity"
// @Success 200 {object} Positionassingment  
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positionassingments [post]
func (ctl *PositionassingmentController) CreatePositionassingment(c *gin.Context) {
	obj := Positionassingment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "positionassingment binding failed",
		})
		return
	}
  
	p, err := ctl.client.Physician.
		Query().
		Where(physician.IDEQ(int(obj.Physicianid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "physician not found",
		})
		return
	}

	d, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Departmentid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "department not found",
		})
		return
	}

	po, err := ctl.client.Position.
		Query().
		Where(position.IDEQ(int(obj.Positionid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "position not found",
		})
		return
	}
	times, err := time.Parse(time.RFC3339, obj.DayStart)

	pa, err := ctl.client.Positionassingment.
		Create().
		SetUser(p).
		SetPosition(po).
		SetDepartment(d).
		SetDayStart(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, pa)
 }
// GetPositionassingment handles GET requests to retrieve a positionassingment entity
// @Summary Get a positionassingment entity by ID
// @Description get positionassingment by ID
// @ID get-positionassingment
// @Produce  json
// @Param id path int true "Positionassingment ID"
// @Success 200 {object} ent.Positionassingment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positionassingments/{id} [get]
func (ctl *PositionassingmentController) GetPositionassingment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pa, err := ctl.client.Positionassingment.
		Query().
		Where(positionassingment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, pa)
 }
 // ListPositionassingment handles request to get a list of positionassingment entities
// @Summary List positionassingment entities
// @Description list positionassingment entities
// @ID list-positionassingment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Positionassingment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positionassingments [get]
func (ctl *PositionassingmentController) ListPositionassingment(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
	positionassingments, err := ctl.client.Positionassingment.
		Query().
		WithUser().
		WithDepartment().
		WithPosition().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
	c.JSON(200, positionassingments)
 }
// DeletePositionassingment handles DELETE requests to delete a positionassingment entity
// @Summary Delete a positionassingment entity by ID
// @Description get positionassingment by ID
// @ID delete-positionassingment
// @Produce  json
// @Param id path int true "Positionassingment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positionassingments/{id} [delete]
func (ctl *PositionassingmentController) DeletePositionassingment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Positionassingment.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
// UpdatePositionassingment handles PUT requests to update a positionassingment entity
// @Summary Update a positionassingment entity by ID
// @Description update positionassingment by ID
// @ID update-positionassingment
// @Accept   json
// @Produce  json
// @Param id path int true "Positionassingment ID"
// @Param positionassingment body ent.Positionassingment true "Positionassingment entity"
// @Success 200 {object} ent.Positionassingment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positionassingments/{id} [put]

func (ctl *PositionassingmentController) UpdatePositionassingment(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	obj := ent.Positionassingment{}
 
	if err := c.ShouldBind(&obj); err != nil {
 
		c.JSON(400, gin.H{
 
			"error": "positionassingment binding failed",
 
		})
 
		return
 
	}
 
	obj.ID = int(id)
 
	u, err := ctl.client.Positionassingment.
 
		UpdateOne(&obj).
 
		Save(context.Background())
 
	if err != nil {
 
		c.JSON(400, gin.H{"error": "update failed",})
 
		return
 
	}
 
 
	c.JSON(200, u)
 
 }
// NewPositionassingmentController creates and registers handles for the positionassingment controller
func NewPositionassingmentController(router gin.IRouter, client *ent.Client) *PositionassingmentController {
	uc := &PositionassingmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
// InitPositionassingmentController registers routes to the main engine
 func (ctl *PositionassingmentController) register() {
	positionassingments := ctl.router.Group("/positionassingments")
	
	positionassingments.GET("", ctl.ListPositionassingment)
	// CRUD
	positionassingments.POST("", ctl.CreatePositionassingment)
	positionassingments.GET(":id", ctl.GetPositionassingment)
	positionassingments.PUT(":id", ctl.UpdatePositionassingment)
	positionassingments.DELETE(":id", ctl.DeletePositionassingment)
 }