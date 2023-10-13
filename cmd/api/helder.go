package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(s http.ResponseWriter, status int, data any) error {
	
}