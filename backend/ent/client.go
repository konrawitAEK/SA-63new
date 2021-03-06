// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/konrawitAEK/app/ent/migrate"

	"github.com/konrawitAEK/app/ent/department"
	"github.com/konrawitAEK/app/ent/physician"
	"github.com/konrawitAEK/app/ent/position"
	"github.com/konrawitAEK/app/ent/positionassingment"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Physician is the client for interacting with the Physician builders.
	Physician *PhysicianClient
	// Position is the client for interacting with the Position builders.
	Position *PositionClient
	// Positionassingment is the client for interacting with the Positionassingment builders.
	Positionassingment *PositionassingmentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Department = NewDepartmentClient(c.config)
	c.Physician = NewPhysicianClient(c.config)
	c.Position = NewPositionClient(c.config)
	c.Positionassingment = NewPositionassingmentClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Department:         NewDepartmentClient(cfg),
		Physician:          NewPhysicianClient(cfg),
		Position:           NewPositionClient(cfg),
		Positionassingment: NewPositionassingmentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:             cfg,
		Department:         NewDepartmentClient(cfg),
		Physician:          NewPhysicianClient(cfg),
		Position:           NewPositionClient(cfg),
		Positionassingment: NewPositionassingmentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Department.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Department.Use(hooks...)
	c.Physician.Use(hooks...)
	c.Position.Use(hooks...)
	c.Positionassingment.Use(hooks...)
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Create returns a create builder for Department.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Create returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{config: c.config}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryFormdepartment queries the formdepartment edge of a Department.
func (c *DepartmentClient) QueryFormdepartment(d *Department) *PositionassingmentQuery {
	query := &PositionassingmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(positionassingment.Table, positionassingment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.FormdepartmentTable, department.FormdepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// PhysicianClient is a client for the Physician schema.
type PhysicianClient struct {
	config
}

// NewPhysicianClient returns a client for the Physician from the given config.
func NewPhysicianClient(c config) *PhysicianClient {
	return &PhysicianClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `physician.Hooks(f(g(h())))`.
func (c *PhysicianClient) Use(hooks ...Hook) {
	c.hooks.Physician = append(c.hooks.Physician, hooks...)
}

// Create returns a create builder for Physician.
func (c *PhysicianClient) Create() *PhysicianCreate {
	mutation := newPhysicianMutation(c.config, OpCreate)
	return &PhysicianCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Physician.
func (c *PhysicianClient) Update() *PhysicianUpdate {
	mutation := newPhysicianMutation(c.config, OpUpdate)
	return &PhysicianUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PhysicianClient) UpdateOne(ph *Physician) *PhysicianUpdateOne {
	mutation := newPhysicianMutation(c.config, OpUpdateOne, withPhysician(ph))
	return &PhysicianUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PhysicianClient) UpdateOneID(id int) *PhysicianUpdateOne {
	mutation := newPhysicianMutation(c.config, OpUpdateOne, withPhysicianID(id))
	return &PhysicianUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Physician.
func (c *PhysicianClient) Delete() *PhysicianDelete {
	mutation := newPhysicianMutation(c.config, OpDelete)
	return &PhysicianDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PhysicianClient) DeleteOne(ph *Physician) *PhysicianDeleteOne {
	return c.DeleteOneID(ph.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PhysicianClient) DeleteOneID(id int) *PhysicianDeleteOne {
	builder := c.Delete().Where(physician.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PhysicianDeleteOne{builder}
}

// Create returns a query builder for Physician.
func (c *PhysicianClient) Query() *PhysicianQuery {
	return &PhysicianQuery{config: c.config}
}

// Get returns a Physician entity by its id.
func (c *PhysicianClient) Get(ctx context.Context, id int) (*Physician, error) {
	return c.Query().Where(physician.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PhysicianClient) GetX(ctx context.Context, id int) *Physician {
	ph, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ph
}

// QueryFormuser queries the formuser edge of a Physician.
func (c *PhysicianClient) QueryFormuser(ph *Physician) *PositionassingmentQuery {
	query := &PositionassingmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ph.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(physician.Table, physician.FieldID, id),
			sqlgraph.To(positionassingment.Table, positionassingment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, physician.FormuserTable, physician.FormuserColumn),
		)
		fromV = sqlgraph.Neighbors(ph.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PhysicianClient) Hooks() []Hook {
	return c.hooks.Physician
}

// PositionClient is a client for the Position schema.
type PositionClient struct {
	config
}

// NewPositionClient returns a client for the Position from the given config.
func NewPositionClient(c config) *PositionClient {
	return &PositionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `position.Hooks(f(g(h())))`.
func (c *PositionClient) Use(hooks ...Hook) {
	c.hooks.Position = append(c.hooks.Position, hooks...)
}

// Create returns a create builder for Position.
func (c *PositionClient) Create() *PositionCreate {
	mutation := newPositionMutation(c.config, OpCreate)
	return &PositionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Position.
func (c *PositionClient) Update() *PositionUpdate {
	mutation := newPositionMutation(c.config, OpUpdate)
	return &PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PositionClient) UpdateOne(po *Position) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPosition(po))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PositionClient) UpdateOneID(id int) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPositionID(id))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Position.
func (c *PositionClient) Delete() *PositionDelete {
	mutation := newPositionMutation(c.config, OpDelete)
	return &PositionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PositionClient) DeleteOne(po *Position) *PositionDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PositionClient) DeleteOneID(id int) *PositionDeleteOne {
	builder := c.Delete().Where(position.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PositionDeleteOne{builder}
}

// Create returns a query builder for Position.
func (c *PositionClient) Query() *PositionQuery {
	return &PositionQuery{config: c.config}
}

// Get returns a Position entity by its id.
func (c *PositionClient) Get(ctx context.Context, id int) (*Position, error) {
	return c.Query().Where(position.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PositionClient) GetX(ctx context.Context, id int) *Position {
	po, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return po
}

// QueryFormposition queries the formposition edge of a Position.
func (c *PositionClient) QueryFormposition(po *Position) *PositionassingmentQuery {
	query := &PositionassingmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(positionassingment.Table, positionassingment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, position.FormpositionTable, position.FormpositionColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PositionClient) Hooks() []Hook {
	return c.hooks.Position
}

// PositionassingmentClient is a client for the Positionassingment schema.
type PositionassingmentClient struct {
	config
}

// NewPositionassingmentClient returns a client for the Positionassingment from the given config.
func NewPositionassingmentClient(c config) *PositionassingmentClient {
	return &PositionassingmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `positionassingment.Hooks(f(g(h())))`.
func (c *PositionassingmentClient) Use(hooks ...Hook) {
	c.hooks.Positionassingment = append(c.hooks.Positionassingment, hooks...)
}

// Create returns a create builder for Positionassingment.
func (c *PositionassingmentClient) Create() *PositionassingmentCreate {
	mutation := newPositionassingmentMutation(c.config, OpCreate)
	return &PositionassingmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Positionassingment.
func (c *PositionassingmentClient) Update() *PositionassingmentUpdate {
	mutation := newPositionassingmentMutation(c.config, OpUpdate)
	return &PositionassingmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PositionassingmentClient) UpdateOne(po *Positionassingment) *PositionassingmentUpdateOne {
	mutation := newPositionassingmentMutation(c.config, OpUpdateOne, withPositionassingment(po))
	return &PositionassingmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PositionassingmentClient) UpdateOneID(id int) *PositionassingmentUpdateOne {
	mutation := newPositionassingmentMutation(c.config, OpUpdateOne, withPositionassingmentID(id))
	return &PositionassingmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Positionassingment.
func (c *PositionassingmentClient) Delete() *PositionassingmentDelete {
	mutation := newPositionassingmentMutation(c.config, OpDelete)
	return &PositionassingmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PositionassingmentClient) DeleteOne(po *Positionassingment) *PositionassingmentDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PositionassingmentClient) DeleteOneID(id int) *PositionassingmentDeleteOne {
	builder := c.Delete().Where(positionassingment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PositionassingmentDeleteOne{builder}
}

// Create returns a query builder for Positionassingment.
func (c *PositionassingmentClient) Query() *PositionassingmentQuery {
	return &PositionassingmentQuery{config: c.config}
}

// Get returns a Positionassingment entity by its id.
func (c *PositionassingmentClient) Get(ctx context.Context, id int) (*Positionassingment, error) {
	return c.Query().Where(positionassingment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PositionassingmentClient) GetX(ctx context.Context, id int) *Positionassingment {
	po, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return po
}

// QueryUser queries the user edge of a Positionassingment.
func (c *PositionassingmentClient) QueryUser(po *Positionassingment) *PhysicianQuery {
	query := &PhysicianQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(positionassingment.Table, positionassingment.FieldID, id),
			sqlgraph.To(physician.Table, physician.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, positionassingment.UserTable, positionassingment.UserColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDepartment queries the department edge of a Positionassingment.
func (c *PositionassingmentClient) QueryDepartment(po *Positionassingment) *DepartmentQuery {
	query := &DepartmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(positionassingment.Table, positionassingment.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, positionassingment.DepartmentTable, positionassingment.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPosition queries the position edge of a Positionassingment.
func (c *PositionassingmentClient) QueryPosition(po *Positionassingment) *PositionQuery {
	query := &PositionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(positionassingment.Table, positionassingment.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, positionassingment.PositionTable, positionassingment.PositionColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PositionassingmentClient) Hooks() []Hook {
	return c.hooks.Positionassingment
}
