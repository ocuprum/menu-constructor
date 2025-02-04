package http

import (
	"net/http"
	"strings"
)

func GetPath(path string) string {
	return "GET " + path
}

func PostPath(path string) string {
	return "POST " + path
}

func PutPath(path string) string {
	return "PUT " + path
}

func DeletePath(path string) string {
	return "DELETE " + path
}

func WriteResponse(resp http.ResponseWriter, statusCode int, explanations ...string) {
	resp.WriteHeader(statusCode)
	resp.Write([]byte(http.StatusText(statusCode)))

	if len(explanations) != 0 {
		resp.Write([]byte("\n"))
		fullText := strings.Join(explanations, "\n")
		resp.Write([]byte(fullText))
	}
}