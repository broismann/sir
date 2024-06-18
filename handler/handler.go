package handler

import (
	"io"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "THis is a server using io.")
}
