package httx

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// Mount httx to a directory in the local file system
func Mount(root string, secrets map[string]string) *Directory {
	return &Directory{root, secrets}
}

// Directory is a pointer to the file system where a website lives
type Directory struct {
	root    string
	secrets map[string]string
}

// ServeHTTP fulfilling http.Handler interface from net/http
func (d *Directory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check for private and hidden files to not be served
	ps := strings.Split(r.URL.Path[1:], "/")
	for _, p := range ps {
		if strings.HasPrefix(p, ".") || strings.HasPrefix(p, "_") {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
	}

	// Default to index if no path specified
	file := strings.Join(ps[0:], "/")
	if file == "" {
		file = "index"
	}

	// Serve file is files is found via Handler
	if h := d.Open(file); h != nil {
		h.ServeHTTP(w, r)
		return
	}

	// Fallback to local file starting with colon
	if name, h := d.fallback(file); h != nil {
		url := strings.Split(file, "/")
		log.Println("URL_"+name[1:len(name)-3], url[len(url)-1])
		r.Header.Set("URL_"+name[1:len(name)-3], url[len(url)-1])
		h.ServeHTTP(w, r)
		return // future proofing
	}

	// Page not found
	http.Error(w, "page not found", http.StatusNotFound)
}

// Open a file inside direcory if it exists or nil
func (d *Directory) Open(file string) *Handler {
	for _, ext := range []string{
		"",
		".sh",
		"/index.sh",
		"/index.html",
	} {
		// Check that file exists in file system
		if f, err := os.Stat(path.Join(d.root, file+ext)); err != nil {
			continue
		} else

		// Check that file isn't a directory
		if f.IsDir() {
			continue
		}

		return Open(d.root, file+ext, d.secrets)
	}
	return nil
}

// Fallback function for path variables
func (d *Directory) fallback(file string) (string, *Handler) {
	parts := strings.Split(path.Join(d.root, file), "/")
	files, err := os.ReadDir(strings.Join(parts[:len(parts)-1], "/"))
	if err != nil {
		return "", nil
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ":") {
			parts := strings.Split(file, "/")
			pdir := strings.Join(parts[:len(parts)-1], "/")
			return f.Name(), d.Open(path.Join(pdir, f.Name()))
		}
	}
	return "", nil
}
