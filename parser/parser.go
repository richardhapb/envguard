package parser

// Parser defines the interface for parsing operations
type Parser interface {
	// Parse reads and parses content from the configured source
	// Returns a map of key-value pairs and any error encountered
	Parse() (map[string]string, error)

	// String returns a formatted string representation of the parsed content
	String() string
}
