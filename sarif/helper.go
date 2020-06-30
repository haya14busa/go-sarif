package sarif

// NewSarif returns new Sarif instance with default value filled in for Version and
// Schema fields.
func NewSarif() *Sarif {
	return &Sarif{
		Version: The210,
		Schema:  String("http://json.schemastore.org/sarif-2.1.0-rtm.4"),
	}
}

func Bool(v bool) *bool          { return &v }
func Int(v int) *int             { return &v }
func Int64(v int64) *int64       { return &v }
func Float64(v float64) *float64 { return &v }
func String(v string) *string    { return &v }
