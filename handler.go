package httx

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pkg/errors"
)

// Open file handler for serving file
func Open(root, file string, secrets map[string]string) *Handler {
	return &Handler{root, file, secrets}
}

// Handler handles an individual file
type Handler struct {
	root, file string
	secrets    map[string]string
}

// ServeHTTP fulfilling http.Handler interface from net/http
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(h.file, ".sh") {
		fileResponse(path.Join(h.root, h.file)).ServeHTTP(w, r)
		return
	}

	res, err := h.Exec(h.env(r), r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	res.ServeHTTP(w, r)
}

// Parse environment variables from requets from net/http
func (h *Handler) env(r *http.Request) (env []string) {
	env = []string{"PATH=" + os.Getenv("PATH")}

	// Merge request headers
	for k, vs := range r.Header {
		v := strings.Join(vs, " ")
		v = strings.ToUpper(strings.ReplaceAll(v, "-", "_"))
		k = strings.ToUpper(strings.ReplaceAll(k, "-", "_"))
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	// Merge request cookies
	for _, c := range r.Cookies() {
		k := fmt.Sprintf("COOKIE_%s", c.Name)
		k = strings.ToUpper(strings.ReplaceAll(k, "-", "_"))
		env = append(env, fmt.Sprintf("%s=%s", k, c.Value))
	}

	// Merge request params
	for k, vs := range r.URL.Query() {
		v := strings.Join(vs, " ")
		k = strings.ToUpper(strings.ReplaceAll(k, "-", "_"))
		env = append(env, fmt.Sprintf("QUERY_%s=%s", k, v))
	}

	// Finally add secrets
	for k, v := range h.secrets {
		k = strings.ToUpper(strings.ReplaceAll(k, " ", "_"))
		k = strings.ToUpper(strings.ReplaceAll(k, "-", "_"))
		env = append(env, fmt.Sprintf("SECRET_%s=%s", k, v))
	}

	return env
}

// Execute file underlying handler
func (h *Handler) Exec(env []string, stdin io.Reader) (*Response, error) {
	var (
		cmd    *exec.Cmd
		err    error
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	// Setup subprocess bash command
	cmd = exec.Command("bash", h.file)
	cmd.Dir = h.root
	cmd.Env = env
	cmd.Stdin = stdin
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	// Run command returning errors
	if err = cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "failed to exec")
	}

	// Capture stderr from bash process
	go func() {
		b, _ := ioutil.ReadAll(&errBuf)
		if len(b) > 0 {
			log.Printf("Error:\n%s", string(b))
		}
	}()

	// Capture stdout to new response
	return loadResponse(&outBuf), nil
}
