package bluemantis

// Status represents a specific type of Rel (relation). This should probably
// be deprecated soon as the package increase size, because the relations will
// differ a little bit more than just ID, Name and Label.
type Status Rel

// ViewState represents a specific type of Rel (relation). This should probably
// be deprecated soon as the package increase size, because the relations will
// differ a little bit more than just ID, Name and Label.
type ViewState Rel

// Project represents a instance of a project, as a single MantisBT application
// can support logging bugs and errors from multiple projects at the same time.
type Project struct {
	ID          uint64    `valid:"optional" json:"id"`
	Name        string    `valid:"required" json:"name"`
	Description string    `valid:"optional" json:"description"`
	Enabled     bool      `valid:"optional" json:"enabled"`
	Status      Status    `valid:"-" json:"status"`
	ViewState   ViewState `valid:"-" json:"view_state"`
}
