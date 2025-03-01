package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	cmd := exec.Command("go", "run", "main.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		t.Fatalf("can't run server: %v", err)
	}
	defer cmd.Process.Kill()

	time.Sleep(time.Second * 10)

	resp, err := http.Get("http://localhost:3000/")
	if err != nil {
		t.Fatalf("can't connect: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("can't read body: %v", err)
	}

	// Checks response
	if !strings.Contains(string(body), "response") {
		t.Errorf("expect 'response', received: %s", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect status 200, received: %d", resp.StatusCode)
	}

	fmt.Println("Server work's", resp.Status, string(body))
}
