package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

//go:generate sh -c "go run example_data/main.go && go fmt example-data.json.go"

const (
	listenAddr = ":9010"
	redisAddr  = "localhost:6379"
	redisKey   = "ModifiedFields"
)

func main() {
	go updater(os.Stdin)
	fmt.Printf("Listen on %s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, http.HandlerFunc(exampleHandler)))
}

func updater(r io.Reader) {
	p := &redis.Pool{
		MaxActive:   100,
		Wait:        true,
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", redisAddr) },
	}

	conn := p.Get()
	defer conn.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		input := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		args := []interface{}{redisKey, "*"}
		for _, d := range input {
			keyValue := strings.SplitN(d, "=", 2)
			if len(keyValue) != 2 {
				continue
			}
			args = append(args, keyValue[0], keyValue[1])
			exampleData[keyValue[0]] = []byte(keyValue[1])
		}

		if len(args) == 2 {
			continue
		}

		if _, err := conn.Do("XADD", args...); err != nil {
			log.Fatalf("Can not send command to redis: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner failed: %v", err)
	}
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Keys []string `json:"requests"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
		return
	}
	log.Println(data.Keys)

	responceData := make(map[string]map[string]map[string]json.RawMessage)
	for _, key := range data.Keys {
		value, ok := exampleData[key]

		if !ok {
			continue
		}

		keyParts := strings.SplitN(key, "/", 3)

		if _, ok := responceData[keyParts[0]]; !ok {
			responceData[keyParts[0]] = make(map[string]map[string]json.RawMessage)
		}

		if _, ok := responceData[keyParts[0]][keyParts[1]]; !ok {
			responceData[keyParts[0]][keyParts[1]] = make(map[string]json.RawMessage)
		}
		responceData[keyParts[0]][keyParts[1]][keyParts[2]] = json.RawMessage(value)
	}

	json.NewEncoder(w).Encode(responceData)
}
