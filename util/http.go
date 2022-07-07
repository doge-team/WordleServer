package util

import (
	"fmt"
	"net/http"
)

type AllowCROS struct {
}

func (AllowCROS) SetupCORS(w *http.ResponseWriter) {
	if w == nil {
		return
	}
	fmt.Println("set CORS succeed")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
