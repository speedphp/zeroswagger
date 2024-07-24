package zeroswagger

import (
	"encoding/json"
	"os"
	"path/filepath"

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

func Generator(apiPath string, jsonPath string) error {
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
