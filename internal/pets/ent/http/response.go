// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	time "time"

	"github.com/google/uuid"
	"github.com/masseelch/elk/internal/pets/ent"
	badge "github.com/masseelch/elk/internal/pets/ent/badge"
	pet "github.com/masseelch/elk/internal/pets/ent/pet"
	playgroup "github.com/masseelch/elk/internal/pets/ent/playgroup"
	toy "github.com/masseelch/elk/internal/pets/ent/toy"

	"github.com/mailru/easyjson"
)

// Basic HTTP Error Response
type ErrResponse struct {
	Code   int         `json:"code"`             // http response status code
	Status string      `json:"status"`           // user-level status message
	Errors interface{} `json:"errors,omitempty"` // application-level error
}

func (e ErrResponse) MarshalToHTTPResponseWriter(w http.ResponseWriter) (int, error) {
	d, err := easyjson.Marshal(e)
	if err != nil {
		return 0, err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(d)))
	w.WriteHeader(e.Code)
	return w.Write(d)
}

func BadRequest(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Forbidden(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func InternalServerError(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func NotFound(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Unauthorized(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

type (
	// $n represents the data serialized for the following serialization group combinations:
	// []
	Badge2492344257View struct {
		ID       int            `json:"id,omitempty"`
		Color    badge.Color    `json:"color,omitempty"`
		Material badge.Material `json:"material,omitempty"`
	}
	Badge2492344257Views []*Badge2492344257View
)

func NewBadge2492344257View(e *ent.Badge) *Badge2492344257View {
	if e == nil {
		return nil
	}
	return &Badge2492344257View{
		ID:       e.ID,
		Color:    e.Color,
		Material: e.Material,
	}
}

func NewBadge2492344257Views(es []*ent.Badge) Badge2492344257Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Badge2492344257Views, len(es))
	for i, e := range es {
		r[i] = NewBadge2492344257View(e)
	}
	return r
}

type (
	// $n represents the data serialized for the following serialization group combinations:
	// [pet:list pet:read]
	Pet3217017920View struct {
		ID         int                      `json:"id,omitempty"`
		Height     int                      `json:"height,omitempty"`
		Weight     float64                  `json:"weight,omitempty"`
		Castrated  bool                     `json:"castrated,omitempty"`
		Name       string                   `json:"name,omitempty"`
		Birthday   time.Time                `json:"birthday,omitempty"`
		Nicknames  []string                 `json:"nicknames,omitempty"`
		Sex        pet.Sex                  `json:"sex,omitempty"`
		Chip       uuid.UUID                `json:"chip,omitempty"`
		Badge      *Badge2492344257View     `json:"badge,omitempty"`
		Protege    *Pet340207500View        `json:"protege,omitempty"`
		Spouse     *Pet340207500View        `json:"spouse,omitempty"`
		Toys       Toy36157710Views         `json:"toys,omitempty"`
		Parent     *Pet340207500View        `json:"parent,omitempty"`
		PlayGroups PlayGroup3432834655Views `json:"play_groups,omitempty"`
		Friends    Pet340207500Views        `json:"friends,omitempty"`
	}
	Pet3217017920Views []*Pet3217017920View
)

func NewPet3217017920View(e *ent.Pet) *Pet3217017920View {
	if e == nil {
		return nil
	}
	return &Pet3217017920View{
		ID:         e.ID,
		Height:     e.Height,
		Weight:     e.Weight,
		Castrated:  e.Castrated,
		Name:       e.Name,
		Birthday:   e.Birthday,
		Nicknames:  e.Nicknames,
		Sex:        e.Sex,
		Chip:       e.Chip,
		Badge:      NewBadge2492344257View(e.Edges.Badge),
		Protege:    NewPet340207500View(e.Edges.Protege),
		Spouse:     NewPet340207500View(e.Edges.Spouse),
		Toys:       NewToy36157710Views(e.Edges.Toys),
		Parent:     NewPet340207500View(e.Edges.Parent),
		PlayGroups: NewPlayGroup3432834655Views(e.Edges.PlayGroups),
		Friends:    NewPet340207500Views(e.Edges.Friends),
	}
}

func NewPet3217017920Views(es []*ent.Pet) Pet3217017920Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Pet3217017920Views, len(es))
	for i, e := range es {
		r[i] = NewPet3217017920View(e)
	}
	return r
}

type (
	// $n represents the data serialized for the following serialization group combinations:
	// []
	Pet340207500View struct {
		ID        int       `json:"id,omitempty"`
		Height    int       `json:"height,omitempty"`
		Weight    float64   `json:"weight,omitempty"`
		Castrated bool      `json:"castrated,omitempty"`
		Name      string    `json:"name,omitempty"`
		Birthday  time.Time `json:"birthday,omitempty"`
		Nicknames []string  `json:"nicknames,omitempty"`
		Sex       pet.Sex   `json:"sex,omitempty"`
		Chip      uuid.UUID `json:"chip,omitempty"`
	}
	Pet340207500Views []*Pet340207500View
)

func NewPet340207500View(e *ent.Pet) *Pet340207500View {
	if e == nil {
		return nil
	}
	return &Pet340207500View{
		ID:        e.ID,
		Height:    e.Height,
		Weight:    e.Weight,
		Castrated: e.Castrated,
		Name:      e.Name,
		Birthday:  e.Birthday,
		Nicknames: e.Nicknames,
		Sex:       e.Sex,
		Chip:      e.Chip,
	}
}

func NewPet340207500Views(es []*ent.Pet) Pet340207500Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Pet340207500Views, len(es))
	for i, e := range es {
		r[i] = NewPet340207500View(e)
	}
	return r
}

type (
	// $n represents the data serialized for the following serialization group combinations:
	// [pet:list]
	Pet45794832View struct {
		ID    int                  `json:"id,omitempty"`
		Name  string               `json:"name,omitempty"`
		Sex   pet.Sex              `json:"sex,omitempty"`
		Chip  uuid.UUID            `json:"chip,omitempty"`
		Badge *Badge2492344257View `json:"badge,omitempty"`
	}
	Pet45794832Views []*Pet45794832View
)

func NewPet45794832View(e *ent.Pet) *Pet45794832View {
	if e == nil {
		return nil
	}
	return &Pet45794832View{
		ID:    e.ID,
		Name:  e.Name,
		Sex:   e.Sex,
		Chip:  e.Chip,
		Badge: NewBadge2492344257View(e.Edges.Badge),
	}
}

func NewPet45794832Views(es []*ent.Pet) Pet45794832Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Pet45794832Views, len(es))
	for i, e := range es {
		r[i] = NewPet45794832View(e)
	}
	return r
}

type (
	// $n represents the data serialized for the following serialization group combinations:
	// []
	PlayGroup3432834655View struct {
		ID          int               `json:"id,omitempty"`
		Title       string            `json:"title,omitempty"`
		Description string            `json:"description,omitempty"`
		Weekday     playgroup.Weekday `json:"weekday,omitempty"`
	}
	PlayGroup3432834655Views []*PlayGroup3432834655View
)

func NewPlayGroup3432834655View(e *ent.PlayGroup) *PlayGroup3432834655View {
	if e == nil {
		return nil
	}
	return &PlayGroup3432834655View{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Weekday:     e.Weekday,
	}
}

func NewPlayGroup3432834655Views(es []*ent.PlayGroup) PlayGroup3432834655Views {
	if len(es) == 0 {
		return nil
	}
	r := make(PlayGroup3432834655Views, len(es))
	for i, e := range es {
		r[i] = NewPlayGroup3432834655View(e)
	}
	return r
}

type (
	// $n represents the data serialized for the following serialization group combinations:
	// []
	Toy36157710View struct {
		ID       int          `json:"id,omitempty"`
		Color    toy.Color    `json:"color,omitempty"`
		Material toy.Material `json:"material,omitempty"`
		Title    string       `json:"title,omitempty"`
	}
	Toy36157710Views []*Toy36157710View
)

func NewToy36157710View(e *ent.Toy) *Toy36157710View {
	if e == nil {
		return nil
	}
	return &Toy36157710View{
		ID:       e.ID,
		Color:    e.Color,
		Material: e.Material,
		Title:    e.Title,
	}
}

func NewToy36157710Views(es []*ent.Toy) Toy36157710Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Toy36157710Views, len(es))
	for i, e := range es {
		r[i] = NewToy36157710View(e)
	}
	return r
}
