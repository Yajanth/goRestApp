package model

import (
	"time"
)

type Playlist struct {
	Name       *string   `json:"name"`
	Songs      []string  `json:"songs"`
	Start_time time.Time `json:"start_time"`
}
