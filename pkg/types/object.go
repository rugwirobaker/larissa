package types

// Object describes a larissa obkect
type Object struct {
	Name    string
	Content []byte
}

// ObjectPage ...
type ObjectPage struct {
	Objects []string `json:"objects"`
}
