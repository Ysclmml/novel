package binding

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// Bind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//     "application/json" --> JSON binding
//     "application/xml"  --> XML binding
// otherwise --> returns an error.
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
func Bind(c *gin.Context, obj interface{}) error {
	b := Default(c.Request.Method, c.ContentType())
	return MustBindWith(c, obj, b)
}

// BindJSON is a shortcut for c.MustBindWith(obj, JSON).
func BindJSON(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, JSON)
}

// BindXML is a shortcut for c.MustBindWith(obj, BindXML).
func BindXML(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, XML)
}

// BindQuery is a shortcut for c.MustBindWith(obj, Query).
func BindQuery(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, Query)
}

// BindYAML is a shortcut for c.MustBindWith(obj, YAML).
func BindYAML(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, YAML)
}

// BindHeader is a shortcut for c.MustBindWith(obj, Header).
func BindHeader(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, Header)
}

// BindUri binds the passed struct pointer using Uri.
// It will abort the request with HTTP 400 if any error occurs.
func BindUri(c *gin.Context, obj interface{}) error {
	if err := ShouldBindUri(c, obj); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind) // nolint: errcheck
		return err
	}
	return nil
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func MustBindWith(c *gin.Context, obj interface{}, b Binding) error {
	if err := c.ShouldBindWith(obj, b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind) // nolint: errcheck
		return err
	}
	return nil
}

// ShouldBind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//     "application/json" --> JSON binding
//     "application/xml"  --> XML binding
// otherwise --> returns an error
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.Bind() but this method does not set the response status code to 400 and abort if the json is not valid.
func ShouldBind(c *gin.Context, obj interface{}) error {
	b := Default(c.Request.Method, c.ContentType())
	return ShouldBindWith(c, obj, b)
}

// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, JSON).
func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, JSON)
}

// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, XML).
func ShouldBindXML(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, XML)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, Query).
func ShouldBindQuery(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, Query)
}

// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, YAML).
func ShouldBindYAML(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, YAML)
}

// ShouldBindHeader is a shortcut for c.ShouldBindWith(obj, Header).
func ShouldBindHeader(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, Header)
}

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func ShouldBindUri(c *gin.Context, obj interface{}) error {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return Uri.BindUri(m, obj)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(c *gin.Context, obj interface{}, b Binding) error {
	return b.Bind(c.Request, obj)
}

// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before  So you should use
// ShouldBindWith for better performance if you need to call only once.
func ShouldBindBodyWith(c *gin.Context, obj interface{}, bb BindingBody) (err error) {
	var body []byte
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, err = ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		c.Set(gin.BodyBytesKey, body)
	}
	return bb.BindBody(body, obj)
}
