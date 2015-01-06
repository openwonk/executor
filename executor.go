package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("ls", "-1", "/home")
	payload, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	users := strings.Split(string(payload), "\n")
	b, err := json.Marshal(users)
	if err != nil {
		log.Println("error:", err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", b)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
