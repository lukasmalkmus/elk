// Code generated by entc, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/masseelch/elk/examples/pets/ent"
	"go.uber.org/zap"
)

// handler has some convenience methods used on node-handlers.
type handler struct{}

// GroupHandler handles http crud operations on ent.Group.
type GroupHandler struct {
	handler

	client    *ent.Client
	log       *zap.Logger
	validator *validator.Validate
}

func NewGroupHandler(c *ent.Client, l *zap.Logger, v *validator.Validate) *GroupHandler {
	return &GroupHandler{
		client:    c,
		log:       l.With(zap.String("handler", "GroupHandler")),
		validator: v,
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *GroupHandler) RegisterHandlers(r chi.Router) {
	// Do no use r.Route() to avoid wildcard matching.
	r.Get("/", h.List)
	r.Post("/", h.Create)
	r.Get("/{id}", h.Read)
	r.Patch("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/{id}/users", h.Users)
	r.Get("/{id}/admin", h.Admin)
}

// PetHandler handles http crud operations on ent.Pet.
type PetHandler struct {
	handler

	client    *ent.Client
	log       *zap.Logger
	validator *validator.Validate
}

func NewPetHandler(c *ent.Client, l *zap.Logger, v *validator.Validate) *PetHandler {
	return &PetHandler{
		client:    c,
		log:       l.With(zap.String("handler", "PetHandler")),
		validator: v,
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *PetHandler) RegisterHandlers(r chi.Router) {
	// Do no use r.Route() to avoid wildcard matching.
	r.Get("/", h.List)
	r.Post("/", h.Create)
	r.Get("/{id}", h.Read)
	r.Patch("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/{id}/friends", h.Friends)
	r.Get("/{id}/owner", h.Owner)
}

// UserHandler handles http crud operations on ent.User.
type UserHandler struct {
	handler

	client    *ent.Client
	log       *zap.Logger
	validator *validator.Validate
}

func NewUserHandler(c *ent.Client, l *zap.Logger, v *validator.Validate) *UserHandler {
	return &UserHandler{
		client:    c,
		log:       l.With(zap.String("handler", "UserHandler")),
		validator: v,
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *UserHandler) RegisterHandlers(r chi.Router) {
	// Do no use r.Route() to avoid wildcard matching.
	r.Get("/", h.List)
	r.Post("/", h.Create)
	r.Get("/{id}", h.Read)
	r.Patch("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/{id}/pets", h.Pets)
	r.Get("/{id}/friends", h.Friends)
	r.Get("/{id}/groups", h.Groups)
	r.Get("/{id}/manage", h.Manage)
}

func (h handler) stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}
