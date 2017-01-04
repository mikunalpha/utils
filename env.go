package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// envs a map stores all key-value pairs in .env file
var envs map[string]string

// LoadEnvs loads .env file
func LoadEnvs(path string) {
	// clean envs
	envs = map[string]string{}

	// read file content
	bytesRaw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot find .env file.")
		return
	}
	rows := strings.Split(string(bytesRaw), "\n")
	for i, _ := range rows {
		rows[i] = strings.Trim(rows[i], " ")

		// parse key and value
		sepIndex := strings.Index(rows[i], "=")
		if sepIndex == -1 {
			continue
		}
		k := strings.Trim(rows[i][:sepIndex], " ")
		v := strings.Trim(rows[i][sepIndex+1:], " ")
		if string(k[0]) == "#" {
			// pass invalid format row and comments
			continue
		}
		envs[k] = v
	}

	// fmt.Println(envs)
}

// EnvString returns string value by key. If key is not exist, it returns spare value
func EnvString(key string, spare string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}
	if s, ok := envs[key]; ok {
		return s
	}
	return spare
}

// EnvBool returns bool value by key. If key is not exist, it returns spare value
func EnvBool(key string, spare bool) bool {
	if s := os.Getenv(key); s != "" {
		if s == "true" {
			return true
		}
		return false
	}
	if s, ok := envs[key]; ok {
		if s == "true" {
			return true
		}
		return false
	}
	return spare
}

// EnvInt returns int value by key. If key is not exist, it returns spare value
func EnvInt(key string, spare int64) int64 {
	if s := os.Getenv(key); s != "" {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return spare
		}
		return i
	}
	if s, ok := envs[key]; ok {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return spare
		}
		return i
	}
	return spare
}

// EnvFloat returns float value by key. If key is not exist, it returns spare value
func EnvFloat(key string, spare float64) float64 {
	if s := os.Getenv(key); s != "" {
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return spare
		}
		return i
	}
	if s, ok := envs[key]; ok {
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return spare
		}
		return i
	}
	return spare
}
