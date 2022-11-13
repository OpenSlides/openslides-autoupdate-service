package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := environment.InterruptContext()
	defer cancel()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s AMOUNT", os.Args[0])
		os.Exit(2)
	}

	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Usage: %s AMOUNT", os.Args[0])
		os.Exit(2)
	}

	if err := run(ctx, amount); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if err := memProfile(); err != nil {
		fmt.Printf("Error on profile: %v", err)
	}

}

func memProfile() error {
	f, err := os.Create("mem.profile")
	if err != nil {
		return fmt.Errorf("creating file for memory profile: %w", err)
	}
	defer f.Close()
	if err := pprof.WriteHeapProfile(f); err != nil {
		return fmt.Errorf("write memory profile: %w", err)
	}
	return nil
}

func run(ctx context.Context, amount int) error {
	ds, usernameIndex, err := buildDB()
	if err != nil {
		return fmt.Errorf("build db: %w", err)
	}

	au, _ := autoupdate.New(ds, restrict.Middleware)

	eg, ctx := errgroup.WithContext(ctx)

	for i := 0; i < amount; i++ {
		i := i
		eg.Go(func() error {
			kb, err := buildRequest()
			if err != nil {
				return fmt.Errorf("building request: %w", err)
			}

			userID := usernameIndex[fmt.Sprintf(`"m96dummy%d"`, i+1)]
			next, _ := au.Connect(userID, kb)()
			if _, err := next(context.Background()); err != nil {
				return fmt.Errorf("next: %w", err)
			}
			return nil
		})
	}

	return eg.Wait()
}

func buildDB() (autoupdate.Datastore, map[string]int, error) {
	fd, err := os.Open("db.json")
	if err != nil {
		return nil, nil, fmt.Errorf("open db file: %w", err)
	}

	var rawData map[string]json.RawMessage
	if err := json.NewDecoder(fd).Decode(&rawData); err != nil {
		return nil, nil, fmt.Errorf("decoding data: %w", err)
	}

	data := make(map[dskey.Key][]byte, len(rawData))
	for k, v := range rawData {
		key, err := dskey.FromString(k)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid key %s: %w", k, err)
		}

		data[key] = v
	}

	ds, _ := dsmock.NewMockDatastore(data)

	usernameIndex := make(map[string]int)
	for k, v := range data {
		if k.Field == "username" {
			usernameIndex[string(v)] = k.ID
		}
	}

	return ds, usernameIndex, nil
}

func buildRequest() (autoupdate.KeysBuilder, error) {
	fd, err := os.Open("request.json")
	if err != nil {
		return nil, fmt.Errorf("open request file: %w", err)
	}

	kb, err := keysbuilder.ManyFromJSON(fd)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	return kb, nil
}
