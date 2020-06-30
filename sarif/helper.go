package sarif

// NewSarif returns new Sarif instance with default value filled in for Version and
// Schema fields.
func NewSarif() *Sarif {
	return &Sarif{
		Version: The210,
		Schema:  String("http://json.schemastore.org/sarif-2.1.0-rtm.4"),
	}
}

// Helpers to get pointer to primitive types.

func Bool(v bool) *bool          { return &v }
func Int(v int) *int             { return &v }
func Int64(v int64) *int64       { return &v }
func Float64(v float64) *float64 { return &v }
func String(v string) *string    { return &v }

// Ptr are helpers to get a pointer to enum value as we cannot get a reference
// to defined const values.

func (v Role) Ptr() *Role                       { return &v }
func (v Level) Ptr() *Level                     { return &v }
func (v Content) Ptr() *Content                 { return &v }
func (v BaselineState) Ptr() *BaselineState     { return &v }
func (v Importance) Ptr() *Importance           { return &v }
func (v ResultKind) Ptr() *ResultKind           { return &v }
func (v SuppressionKind) Ptr() *SuppressionKind { return &v }
func (v State) Ptr() *State                     { return &v }
func (v Version) Ptr() *Version                 { return &v }
func (v ColumnKind) Ptr() *ColumnKind           { return &v }
