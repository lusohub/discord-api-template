package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gtuk/discordwebhook"
)

func main() {
	discordURL := os.Getenv("DISCORD_URL")
	if discordURL == "" {
		log.Fatal("DISCORD_URL environment variable is required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	http.HandleFunc("/send/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Path[6:]
		if name == "" {
			http.Error(w, "Nome é obrigatório na URL. Exemplo: /joao, /maria", http.StatusBadRequest)
			return
		}

		content := fmt.Sprintf("Olá, %s", name)
		username := "API Bot"
		
		message := discordwebhook.Message{
			Username: &username,
			Content: &content,
   		}

		err := discordwebhook.SendMessage(discordURL, message)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Mensagem enviada para Discord com sucesso", "content": "%s"}`, content)
	})

	port := os.Getenv("PORT")	
	if port == "" {
		port = "8080"
	}

	fmt.Printf("OK: http://localhost:%s/[nome]\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}