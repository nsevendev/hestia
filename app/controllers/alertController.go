package controllers

import (
	"html/template"
	"net/http"
)

// FlashMessage structure
type FlashMessage struct {
	Type    string // "success", "warning", "error"
	Message string
}

// RenderTemplate est une fonction réutilisable pour afficher un template avec un message d'alerte
func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, tplName string, data map[string]interface{}, flash *FlashMessage) {
	if flash != nil {
		data["Flash"] = flash
	}
	err := tmpl.ExecuteTemplate(w, tplName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
