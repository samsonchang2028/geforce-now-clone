package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func handleJob(w http.ResponseWriter, r *http.Request, q *Queue) {
	var job Job
	//write json to address of job from r
	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	job.ID = uuid.NewString()

	q.Enqueue(job)

	//prep HTTP response back to frontend, tels body im sending json
	w.Header().Set("Content-Type", "application/json")
	//sends 202 accepted, but job isnt done
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(map[string]string{
		"job_id":  string(job.ID),
		"message": "job queued",
	})
}

func main() {
	q := &Queue{}
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		handleJob(w, r, q)
	})

	http.ListenAndServe(":8080", nil)

}
