package models

import (
	"time"

	"github.com/golang-module/carbon"
	"github.com/judegiordano/gogetem/pkg/mongo"
	"github.com/judegiordano/gogetem/pkg/nanoid"
)

type Health struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	Message   string    `bson:"message,omitempty" json:"message"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}

func (d *Health) Save() (*Health, error) {
	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}
	d.Id = id
	d.CreatedAt = carbon.Now().Carbon2Time().UTC()
	d.UpdatedAt = carbon.Now().Carbon2Time().UTC()
	return mongo.Insert(*d)
}
