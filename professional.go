package main

import (
	"fmt"
)

const (
	// StatusLevelInfo defines minimum score for Info level
	StatusLevelInfo float64 = 70
	// StatusLevelWarning defines minimum score for Warning level
	StatusLevelWarning float64 = 50
)

// Status is a type defining status of the reviwe
type Status int

const (
	// StatusInfo means review is valid
	StatusInfo Status = iota
	// StatusWarning means review should be investigated
	StatusWarning
	// StatusAlert means review should not be trusted
	StatusAlert
	// StatusInvalid means review is not valid
	StatusInvalid
)

func (s Status) String() string {
	switch s {
	case StatusInfo:
		return "Info"
	case StatusWarning:
		return "Warning"
	case StatusAlert:
		return "Alert"
	case StatusInvalid:
		return "Invalid"
	}
	return "Unknown"
}

// Professional contains definition of a professional
type Professional struct {
	Name   string  `json:"name"`
	status Status  `json:"status"`
	score  float64 `json:"score"`
}

// SetScore sets score of the review and updates review status accordingly
func (p *Professional) SetScore(score float64) {
	p.score = score
	p.sanitizeScore()
	p.updateStatus()
}

// AdjustScore adjusts score by fraction of score
func (p *Professional) AdjustScore(f float64) {
	p.score += f * 100
	p.sanitizeScore()
	p.updateStatus()
}

func (p *Professional) sanitizeScore() {
	if p.score > 100.0 {
		p.score = 100.0
	}
	if p.score < 0.0 {
		p.score = 0.0
	}
}

func (p *Professional) updateStatus() {
	if p.score >= StatusLevelInfo {
		p.status = StatusInfo
		return
	}
	if p.score >= StatusLevelWarning {
		p.status = StatusWarning
		return
	}
	p.status = StatusAlert
}

// GetScore return current review score
func (p Professional) GetScore() float64 {
	return p.score
}

// GetStatus returns current review status
func (p Professional) GetStatus() Status {
	return p.status
}

func (p Professional) String() string {
	switch p.status {
	case StatusInfo:
		fallthrough
	case StatusWarning:
		return fmt.Sprintf("%s: %s has a trusted review score of %g", p.status, p.Name, p.score)
	case StatusAlert:
		return fmt.Sprintf("%s: %s has been de-activated due to a low trusted review score", p.status, p.Name)
	}

	return ""
}

var professionals = []*Professional{}

// GetProfessional creates new professional or gets from "database"
func GetProfessional(name string) *Professional {
	p := GetProfessionalByName(name)
	if p == nil {
		p = &Professional{
			Name: name,
		}
		p.SetScore(100)
		professionals = append(professionals, p)
	}
	return p
}

// GetProfessionalByName gets professional of given name from "database"
func GetProfessionalByName(name string) *Professional {
	for i := range professionals {
		if professionals[i].Name == name {
			return professionals[i]
		}
	}
	return nil
}
