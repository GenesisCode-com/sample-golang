package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

func main() {
	http.HandleFunc("/trk", func(w http.ResponseWriter, r *http.Request) {
		// Chuyển hướng người dùng sang trang web của Google
		http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently)
	})

	
	
}
