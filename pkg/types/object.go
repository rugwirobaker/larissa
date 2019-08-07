package types

// Object describes a larissa obkect
type Object struct {
	Name    string
	Content []byte
}

// Serialize turns content bytes into a savable larissa Object
func (obj *Object) Serialize() error { return nil }
