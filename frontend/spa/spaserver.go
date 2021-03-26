package spa

import (
	"net/http"
	"os"
	"path/filepath"
)


type spaHandler struct {
	staticPath string
	indexFile string
}


func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(h.staticPath, filepath.Clean(r.URL.Path))

	if info, err := os.Stat(p); err != nil {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexFile))
		return
	} else if info.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexFile))
		return
	}

	http.ServeFile(w, r, p)
}


func SpaHandler(publicDir string, indexFile string) http.Handler {
	return &spaHandler{publicDir, indexFile}
}
