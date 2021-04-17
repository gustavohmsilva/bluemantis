package bluemantis

// Rel represent a generic relationship with a different object, such as
// Categories, Fields, Handlers, Status and etc. It is usually re-labeled to
// match the object name through the package.
type Rel struct {
	ID    uint64 `valid:"optional"`
	Name  string `valid:"required"`
	Label string `valid:"optional"`
}
