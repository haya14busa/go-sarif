package sarif

// New returns new Sarif instance with default value filled in for Version and
// Schema fields.
func New() *Sarif {
	return &Sarif{
		Version: "2.1.0",
		Schema:  "http://json.schemastore.org/sarif-2.1.0-rtm.4",
	}
}
