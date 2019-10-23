package urlshort

import (
	"net/http"
)

//fallback is for when no path matches in the map
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// If we can match a path, then redirect to it
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok { // ok is bool
			http.Redirect(w, r, dest, http.StatusFound)
			return			
		}
		// If redirect fails
		fallback.ServeHTTP(w, r)
	}
}

func 