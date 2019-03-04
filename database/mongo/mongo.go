package mongo

import (
	"gopkg.in/mgo.v2"
	"strings"
)

type MongoGuy struct {
	Session *mgo.Session
}

// NewSession returns a new mongodb session instance
func NewSession(servers ...string) (*mgo.Session, error) {
	session, err := mgo.Dial(strings.Join(servers,","))
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
