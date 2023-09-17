package gde

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/deeper-x/golang-dot-env/settings"
)

// EnvFile object representing .env file
type EnvFile struct {
	content map[string]string
}

// New EnvFile builder, given an input file
func New(fpath ...string) (EnvFile, error) {
	res := EnvFile{}
	dotEnvFile := settings.DefaultFile

	// check for paramters: 0 or 1 are accepted
	// 0 param makes default value to be set
	// 1 param makes it to value
	if len(fpath) > 1 {
		return res, errors.New("<= 1 parameter is required")
	}

	if len(fpath) == 1 {
		dotEnvFile = fpath[0]
	}

	// 1. create Env File Content
	efc, err := readEnvFile(dotEnvFile)
	if err != nil {
		return res, err
	}

	// 2. build a k:v map object
	kv, err := buildKV(efc)
	if err != nil {
		return res, err
	}

	// 3. build EnvFile object
	res.content = kv

	return res, nil
}

// Get returns a string value, for an input key
func (ef *EnvFile) Get(k string) (string, error) {
	v := ef.content[k]
	if v == "" {
		return "", errors.New("no given value")
	}

	return v, nil
}

func checkRowsFormat(values []string) (bool, error) {
	res := false
	rxp, err := regexp.Compile(settings.RegexRow)
	if err != nil {
		return res, err
	}

	for i := range values {
		if !rxp.MatchString(values[i]) {
			return res, nil
		}
	}

	res = true
	return res, nil
}

// readEnvFile read .env content
func readEnvFile(fpath ...string) ([]string, error) {
	dotEnvFile := settings.DefaultFile
	res := []string{}

	// check for paramters: 0 or 1 are accepted
	// 0 param makes default value to be set
	// 1 param makes it to value
	if len(fpath) > 1 {
		return res, errors.New("<= 1 parameter is required")
	}

	if len(fpath) == 1 {
		dotEnvFile = fpath[0]
	}

	fcontent, err := os.Open(dotEnvFile)
	if err != nil {
		return res, err
	}

	fscanner := bufio.NewScanner(fcontent)
	fscanner.Split(bufio.ScanLines)

	for fscanner.Scan() {
		res = append(res, fscanner.Text())
	}

	return res, nil
}

// buildKV returns a k:v map with all environment vars
func buildKV(values []string) (map[string]string, error) {
	res := map[string]string{}

	for i := range values {
		cursl := strings.Split(values[i], "=")

		k := cursl[0]
		v := cursl[1]

		res[k] = v
	}

	return res, nil
}
