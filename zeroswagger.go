package zeroswagger

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest"
)

//go:embed assets/*
var staticFiles embed.FS
var tempDir = "./temp" // 存放JSON的临时目录

type ZeroSwaggerHandler struct {
	jsonList []JsonPath
	docPath  string
}

func New(path string) *ZeroSwaggerHandler {
	return &ZeroSwaggerHandler{
		jsonList: GenerateApi(path, ".", tempDir),
		docPath:  path,
	}
}

func (m *ZeroSwaggerHandler) Route() rest.Route {
	return rest.Route{
		Method: http.MethodGet,
		Path:   m.docPath + "/:file",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if path == m.docPath+"/swagger-initializer.js" {
				serveInitFile(w, r, m.jsonList)
			} else if strings.HasPrefix(path, m.docPath+"/api-") {
				for _, jsonPath := range m.jsonList {
					if path == jsonPath.JsonFile {
						http.ServeFile(w, r, tempDir+"/"+jsonPath.RealFileName)
					}
				}
			} else {
				staticContent, _ := fs.Sub(staticFiles, "assets")
				http.StripPrefix(m.docPath+"/", http.FileServer(http.FS(staticContent))).ServeHTTP(w, r)
			}
		}),
	}
}

func serveInitFile(w http.ResponseWriter, _ *http.Request, jsonList []JsonPath) {
	fmt.Print(jsonList)
	content, _ := staticFiles.ReadFile("assets/swagger-initializer.js")
	jsonContents, _ := json.Marshal(jsonList)
	fmt.Print(string(jsonContents))
	modifiedContent := strings.Replace(string(content), "###JSONLIST###", string(jsonContents), -1)
	w.Header().Set("Content-Type", "application/x-javascript")
	w.Write([]byte(modifiedContent))
}
