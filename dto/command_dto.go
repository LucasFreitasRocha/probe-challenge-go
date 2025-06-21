package dto

type CommandDTO struct {
	IdProbe uint   `json:"probe_id"`
	Command string `json:"command"`
}