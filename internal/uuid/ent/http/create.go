// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/uuid/ent"
	"github.com/masseelch/elk/internal/uuid/ent/user"
	"go.uber.org/zap"
)

// Create creates a new ent.User and stores it in the database.
func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d UserCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.User.Create()
	if d.UUID != nil {
		b.SetUUID(*d.UUID)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create user", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", e.ID))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", e.ID))
			BadRequest(w, msg)
		default:
			l.Error("could not read user", zap.Error(err), zap.Int("id", e.ID))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("user rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewUser3451555716View(e), w)
}
