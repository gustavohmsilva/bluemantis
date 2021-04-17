package bluemantis

// Project represents a instance of a project, as a single MantisBT application
// can support logging bugs and errors from multiple projects at the same time.
type Project struct {
	ID          uint64 `valid:"optional"`
	Name        string `valid:"required"`
	Description string `valid:"optional"`
	Enabled     bool   `valid:"optional"`
	FilePath    string `valid:"optional"`
	Status      Rel    `valid:"-"`
	ViewState   Rel    `valid:"-"`
}
