package urlshort

import (
	"net/http"
	"fmt"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		// if match path redirect to it
		path := r.URL.Path
		if url, found := pathsToUrls[path]; found {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		// else use fallback
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type Redirect struct {
	Path string `yaml:"path"`
	Url string  `yaml:"url"`
}

type ParsedYaml struct{
	Redirect []Redirect `yaml:",flow"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	out := ParsedYaml{}
	err := yaml.Unmarshal(yml, &out)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(out)
	return MapHandler(pathMap, fallback), nil
}


func buildMap(y ParsedYaml) (map[string]string){
	fmt.Printf("%v",y)
	return make(map[string]string)
}
