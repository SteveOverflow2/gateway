package rest

import (
	"Gateway/pkg/config"
	"Gateway/pkg/http/rest/handlers"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

var (
	PostURL string
)

func InitUrls(cfg config.URLConfig) {
	PostURL = cfg.PostURL
}
func CreateReverseProxy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		host, err := url.Parse(GenerateAddr(getTargetAddress(mux.Vars(r)["service"]), r.URL.Path))
		if err != nil {
			return
		}
		if host.Scheme == "" {
			fmt.Printf("r.Host: %v\n", r.Host+r.URL.Path)
			handlers.RenderResponse(w, http.StatusNotFound, "route not found")
			return
		}
		fmt.Printf("proxying : %v\n", host)
		Proxy(host).ServeHTTP(w, r)
	}
}

func PrintRequest() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---------------------------------------------")
		fmt.Printf("r.Host: %v\n", r.Host)
		fmt.Println("404 came in")
		fmt.Println("---------------------------------------------")

	}
}

func GenerateAddr(targetAddr string, urlParts string) string {
	parts := strings.Split(strings.TrimPrefix(urlParts, "/"), "/")
	for i := 1; i < len(parts); i++ {
		targetAddr += "/" + parts[i]
	}
	return targetAddr
}

func getTargetAddress(hostName string) string {
	fmt.Printf("Checking hostname: %v\n", hostName)
	switch hostName {
	case "posts":
		fmt.Printf("Returning URL: %v\n", PostURL)
		return PostURL
	default:
		return ""
	}
}

func Proxy(address *url.URL) *httputil.ReverseProxy {
	p := httputil.NewSingleHostReverseProxy(address)
	p.Director = func(request *http.Request) {
		request.Host = address.Host
		request.URL.Scheme = address.Scheme
		request.URL.Host = address.Host
		request.URL.Path = address.Path
	}
	return p
}
