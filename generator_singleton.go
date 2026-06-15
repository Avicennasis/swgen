package swgen

import "net/http"

// singleton package generator
var gen = NewGenerator()

// EnableCORS enables HTTP handler to support CORS
func EnableCORS(b bool, allowHeaders ...string) *Generator {
	return gen.EnableCORS(b, allowHeaders...)
}

// SetHost sets host info for swagger specification
func SetHost(host string) *Generator {
	return gen.SetHost(host)
}

// SetBasePath sets base path info for swagger specification
func SetBasePath(basePath string) *Generator {
	return gen.SetBasePath(basePath)
}

// SetContact sets contact information for API
func SetContact(name, url, email string) *Generator {
	return gen.SetContact(name, url, email)
}

// SetInfo sets information about API
func SetInfo(title, description, term, version string) *Generator {
	return gen.SetInfo(title, description, term, version)
}

// SetLicense sets license information for API
func SetLicense(name, url string) *Generator {
	return gen.SetLicense(name, url)
}

// AddExtendedField adds vendor extension field to document
func AddExtendedField(name string, value interface{}) *Generator {
	return gen.AddExtendedField(name, value)
}

// AddTypeMap adds a rule to use dst interface instead of src
func AddTypeMap(src interface{}, dst interface{}) *Generator {
	return gen.AddTypeMap(src, dst)
}

// GenDocument returns document specification in JSON string (in []byte)
func GenDocument() ([]byte, error) {
	return gen.GenDocument()
}

// ServeHTTP implements http.HandlerFunc to serve swagger.json document
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gen.ServeHTTP(w, r)
}
