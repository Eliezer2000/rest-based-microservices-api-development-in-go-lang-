package time_http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eliezer2000/task01/internal/ports"
)

type TimeHandler struct {
	TimeService ports.TimeProvider
}

func (h *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tzParam := r.URL.Query().Get("tz")
	timeZones := strings.Split(tzParam, ",")

	if len(timeZones) == 0 || len(timeZones) == 1 && timeZones[0] == "" {
		timeZones = []string{"UTC"}
	}

	response := make(map[string]string)

	for _, tz := range timeZones {
		if tz == "" {
			tz = "UTC"
		}
		currentTime, err := h.TimeService.GetCurrentTime(tz)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response[tz] = currentTime
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
