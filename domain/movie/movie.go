package movie

import (
	"context"
	"time"

	"github.com/gilcrest/go-api-basic/domain/errs"
	"github.com/gilcrest/go-api-basic/domain/user"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Writer is used to create/update a Movie
type Writer interface {
	Add(ctx context.Context) error
	Update(ctx context.Context, id string) error
}

// NewMovie initializes a Movie struct for use in Movie creation
func NewMovie(id uuid.UUID, extlID string, u *user.User) (*Movie, error) {
	switch {
	case id == uuid.Nil:
		return nil, errs.E(errs.Validation, errs.Parameter("ID"), errors.New(errs.MissingField("ID").Error()))
	case extlID == "":
		return nil, errs.E(errs.Validation, errs.Parameter("ID"), errors.New(errs.MissingField("ID").Error()))
	case !u.IsValid():
		return nil, errs.E(errs.Validation, errs.Parameter("User"), errors.New("User is invalid"))
	}

	now := time.Now()

	return &Movie{
		ID:         id,
		ExternalID: extlID,
		CreateUser: u,
		CreateTime: now,
		UpdateUser: u,
		UpdateTime: now,
	}, nil
}

// Movie holds details of a movie
type Movie struct {
	ID         uuid.UUID
	ExternalID string
	Title      string
	Rated      string
	Released   time.Time
	RunTime    int
	Director   string
	Writer     string
	CreateUser *user.User
	CreateTime time.Time
	UpdateUser *user.User
	UpdateTime time.Time
}

func (m *Movie) SetTitle(t string) *Movie {
	m.Title = t
	return m
}

func (m *Movie) SetRated(r string) *Movie {
	m.Rated = r
	return m
}

func (m *Movie) SetReleased(t time.Time) *Movie {
	m.Released = t
	return m
}

func (m *Movie) SetRunTime(rt int) *Movie {
	m.RunTime = rt
	return m
}

func (m *Movie) SetDirector(d string) *Movie {
	m.Director = d
	return m
}

func (m *Movie) SetWriter(w string) *Movie {
	m.Writer = w
	return m
}

// IsValid performs validation of the struct
func (m *Movie) IsValid() error {
	switch {
	case m.Title == "":
		return errs.E(errs.Validation, errs.Parameter("Title"), errs.MissingField("Title"))
	case m.Rated == "":
		return errs.E(errs.Validation, errs.Parameter("Rated"), errs.MissingField("Rated"))
	case m.Released.IsZero() == true:
		return errs.E(errs.Validation, errs.Parameter("ReleaseDate"), "Released must have a value")
	case m.RunTime <= 0:
		return errs.E(errs.Validation, errs.Parameter("RunTime"), "Run time must be greater than zero")
	case m.Director == "":
		return errs.E(errs.Validation, errs.Parameter("Director"), errs.MissingField("Director"))
	case m.Writer == "":
		return errs.E(errs.Validation, errs.Parameter("Writer"), errs.MissingField("Writer"))
	}

	return nil
}

// Update performs business validations prior to writing to the db
func (m *Movie) Update(ctx context.Context, id string) error {
	m.UpdateTime = time.Now().UTC()

	return nil
}
