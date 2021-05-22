package tests

import (
	"testing"
)

func TestCases(t *testing.T) {
	files, err := walk("testcases")
	if err != nil {
		t.Fatalf("Can not work test case files: %v", err)
	}
	for _, file := range files {
		c, err := loadFile(file)
		if err != nil {
			t.Fatalf("Can not load test case file %s: %v", file, err)
		}

		c.walk(func(c *Case) {
			t.Run(c.Name, func(t *testing.T) {
				c.test(t)
			})
		})

	}
}
