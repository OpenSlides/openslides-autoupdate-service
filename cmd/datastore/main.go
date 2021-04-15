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

	models "github.com/OpenSlides/openslides-models-to-go"
	"github.com/gomodule/redigo/redis"
)

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
		if len(scanner.Text()) == 0 {
			continue
		}

		var data map[string]json.RawMessage
		if err := json.Unmarshal(scanner.Bytes(), &data); err != nil {
			log.Printf("Invalid json input: %v", err)
			continue
		}

		args := []interface{}{redisKey, "*"}
		for key, value := range data {
			args = append(args, key, string(value))
			models.ExampleData[key] = value
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
		value, ok := models.ExampleData[key]

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
		responceData[keyParts[0]][keyParts[1]][keyParts[2]] = value
	}

	json.NewEncoder(w).Encode(responceData)
}
