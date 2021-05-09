package bluemantis

// Project represents a instance of a project, as a single MantisBT application
// can support logging bugs and errors from multiple projects at the same time.
type Project struct {
	ID          uint64 `valid:"optional" json:"id"`
	Name        string `valid:"required" json:"name"`
	Description string `valid:"optional" json:"description"`
	Enabled     bool   `valid:"optional" json:"enabled"`
	Status      Status `valid:"-" json:"status"`
	ViewState   Rel    `valid:"-" json:"view_state"`
	FilePath    string `valid:"_" json:"file_path"`
}
