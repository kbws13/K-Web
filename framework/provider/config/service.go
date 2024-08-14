package config

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"bytes"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type KConfig struct {
	c framework.Container

	folder   string
	env      string
	keyDelim string

	// envMap
	envMaps  map[string]string
	confMaps map[string]interface{}
	confRaws map[string][]byte
}

func NewKConfig(params ...interface{}) (interface{}, error) {
	if len(params) != 4 {
		return nil, errors.New("NewKConfig params error")
	}

	folder := params[0].(string)
	envMaps := params[1].(map[string]string)
	env := params[2].(string)

	c := params[3].(framework.Container)

	envFolder := filepath.Join(folder, env)
	// check folder exist
	if _, err := os.Stat(envFolder); os.IsNotExist(err) {
		return nil, errors.New("folder " + envFolder + " not exist: " + err.Error())
	}

	kConf := &KConfig{
		c:        c,
		folder:   folder,
		env:      env,
		keyDelim: ".",
		envMaps:  envMaps,
		confMaps: map[string]interface{}{},
		confRaws: map[string][]byte{},
	}

	// read all yml/yaml files in folder
	files, err := os.ReadDir(envFolder)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	for _, file := range files {
		s := strings.Split(file.Name(), ".")
		if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
			name := s[0]

			// read file bytes
			bf, err := os.ReadFile(filepath.Join(envFolder, file.Name()))
			if err != nil {
				continue
			}
			kConf.confRaws[name] = bf
			// do replace
			bf = replace(bf, envMaps)
			// parse yaml
			c := map[string]interface{}{}
			if err := yaml.Unmarshal(bf, &c); err != nil {
				continue
			}
			kConf.confMaps[name] = c
		}
	}

	// init app path
	if kConf.IsExist("app.path") && c.IsBind(contract.AppKey) {
		appPaths := kConf.GetStringMapString("app.path")
		appService := c.MustMake(contract.AppKey).(contract.App)
		appService.LoadAppConfig(appPaths)
	}
	return kConf, nil
}

func replace(content []byte, maps map[string]string) []byte {
	if maps == nil {
		return content
	}

	for key, val := range maps {
		reKey := "env(" + key + ")"
		content = bytes.ReplaceAll(content, []byte(reKey), []byte(val))
	}

	return content
}

func searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	next, ok := source[path[0]]
	if ok {
		// Fast path
		if len(path) == 1 {
			return next
		}

		// Nested case
		switch next.(type) {
		case map[interface{}]interface{}:
			return searchMap(cast.ToStringMap(next), path[1:])
		case map[string]interface{}:
			// Type assertion is safe here since it is only reached
			// if the type of `next` is the same as the type being asserted
			return searchMap(next.(map[string]interface{}), path[1:])
		default:
			// got a value but nested key expected, return "nil" for not found
			return nil
		}
	}
	return nil
}

func (conf *KConfig) find(key string) interface{} {
	return searchMap(conf.confMaps, strings.Split(key, conf.keyDelim))
}

// IsExist check setting is exist
func (conf *KConfig) IsExist(key string) bool {
	return conf.find(key) != nil
}

// Get a new interface
func (conf *KConfig) Get(key string) interface{} {
	return conf.find(key)
}

// GetBool get bool type
func (conf *KConfig) GetBool(key string) bool {
	return cast.ToBool(conf.find(key))
}

// GetInt get Int type
func (conf *KConfig) GetInt(key string) int {
	return cast.ToInt(conf.find(key))
}

// GetFloat64 get float64
func (conf *KConfig) GetFloat64(key string) float64 {
	return cast.ToFloat64(conf.find(key))
}

// GetTime get time type
func (conf *KConfig) GetTime(key string) time.Time {
	return cast.ToTime(conf.find(key))
}

// GetString get string typen
func (conf *KConfig) GetString(key string) string {
	return cast.ToString(conf.find(key))
}

// GetIntSlice get int slice type
func (conf *KConfig) GetIntSlice(key string) []int {
	return cast.ToIntSlice(conf.find(key))
}

// GetStringSlice get string slice type
func (conf *KConfig) GetStringSlice(key string) []string {
	return cast.ToStringSlice(conf.find(key))
}

// GetStringMap get map which key is string, value is interface
func (conf *KConfig) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(conf.find(key))
}

// GetStringMapString get map which key is string, value is string
func (conf *KConfig) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(conf.find(key))
}

// GetStringMapStringSlice get map which key is string, value is string slice
func (conf *KConfig) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(conf.find(key))
}

// Load a config to a struct, val should be an pointer
func (conf *KConfig) Load(key string, val interface{}) error {
	return mapstructure.Decode(conf.find(key), val)
}
