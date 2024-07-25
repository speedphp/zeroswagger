package zeroswagger

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/goctl-swagger/generate"
)

type Plugin struct {
	Api         *spec.ApiSpec
	ApiFilePath string
	Style       string
	Dir         string
}

type JsonPath struct {
	JsonFile     string `json:"url"`
	Name         string `json:"name"`
	RealFileName string
}

func GenerateApi(docPath string, basePath string, tempDir string) []JsonPath {
	if basePath == "" {
		basePath = "."
	}
	exeDir := filepath.Dir(basePath)
	files, _ := findFilesWithExt(exeDir, ".api")

	if tempDir == "" {
		tempDir = "./temp"
	}
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		panic(err)
	}

	var resultMap []JsonPath
	for _, apiPath := range files {
		jsonFile := strings.ReplaceAll(apiPath, "/", "-") + ".json"
		jsonPath := filepath.Join(tempDir, jsonFile)
		if err := generator(apiPath, jsonPath); err != nil {
			panic(err)
		}
		resultMap = append(resultMap, JsonPath{
			JsonFile:     docPath + "/api-" + jsonFile,
			Name:         apiPath,
			RealFileName: jsonFile,
		})
	}
	return resultMap
}

func generator(apiPath string, jsonPath string) error {
	apiFilePath, _ := filepath.Abs(apiPath)
	jsonFilePath, _ := filepath.Abs(jsonPath)
	dir := filepath.Dir(jsonFilePath)
	fileName := filepath.Base(jsonFilePath)

	data, err := prepareArgs(apiFilePath, dir, "")
	if err != nil {
		return err
	}

	tmpFile, _ := os.CreateTemp("", "tempfile")
	defer os.Remove(tmpFile.Name())
	tmpFile.Write(data)
	tmpFile.Seek(0, 0)
	originalStdin := os.Stdin
	os.Stdin = tmpFile

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	os.Stdin = originalStdin
	return generate.Do(fileName, "", "", "", p)
}

func prepareArgs(apiPath string, dir string, style string) ([]byte, error) {
	var transferData Plugin
	if len(apiPath) > 0 && pathx.FileExists(apiPath) {
		api, err := parser.Parse(apiPath)
		if err != nil {
			return nil, err
		}

		transferData.Api = api
	}

	absApiFilePath, err := filepath.Abs(apiPath)
	if err != nil {
		return nil, err
	}

	transferData.ApiFilePath = absApiFilePath
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	transferData.Dir = dirAbs
	transferData.Style = style
	data, err := json.Marshal(transferData)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func findFilesWithExt(rootDir, ext string) ([]string, error) {
	var result []string
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ext {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
