package bluemantis

type Project struct {
	ID          uint64
	Name        string
	Description string
	Enabled     bool
	FilePath    string
	Status      Rel
	ViewState   Rel
}
