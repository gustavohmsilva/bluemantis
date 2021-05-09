package bluemantis

import "time"

// Meta stores time sensitive information, such as last time a given object was
// created, edited, deleted and by whom those operations have been done.
type Meta struct {
	CreatedAt time.Time `valid:"-" json:"created_at"`
	UpdatedAt time.Time `valid:"-" json:"updated_at"`
}
