package rest

import (
	"Gateway/pkg/config"
	"Gateway/pkg/http/rest/handlers"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

var (
	CouponsURL string
	BrandURL   string
	ImportURL  string
	UserURL    string
)

func InitUrls(cfg config.URLConfig) {
	CouponsURL = cfg.CouponsURL
	BrandURL = cfg.BrandURL
	ImportURL = cfg.ImportURL
	UserURL = cfg.UserURL
}
func CreateReverseProxy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		host, err := url.Parse(GenerateAddr(getTargetAddress(mux.Vars(r)["service"]), r.URL.Path))
		if err != nil {
			return
		}
		if host.Scheme == "" {
			handlers.RenderResponse(w, http.StatusNotFound, "route not found")
			return
		}
		Proxy(host).ServeHTTP(w, r)
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
	switch hostName {
	case "brands":
		return BrandURL
	case "import":
		return ImportURL
	case "user":
		return UserURL
	case "coupons":
		return CouponsURL
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
