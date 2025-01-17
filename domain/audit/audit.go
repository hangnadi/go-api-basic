package audit

import (
	"net/http"
	"time"

	"github.com/gilcrest/go-api-basic/domain/app"
	"github.com/gilcrest/go-api-basic/domain/user"
)

// Audit represents the moment an app/user interacted with the system
type Audit struct {
	App    app.App
	User   user.User
	Moment time.Time
}

// FromRequest is a convenience function that retrieves the App
// and User structs from the request context. The moment is also
// set to time.Now
func FromRequest(r *http.Request) (Audit, error) {
	var (
		a   app.App
		u   user.User
		err error
	)

	a, err = app.FromRequest(r)
	if err != nil {
		return Audit{}, err
	}

	u, err = user.FromRequest(r)
	if err != nil {
		return Audit{}, err
	}

	return Audit{App: a, User: u, Moment: time.Now()}, nil
}
