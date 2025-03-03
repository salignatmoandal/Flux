package cloud

import (
	"time"
)

type Resource struct {
	ID           string
	Type         string
	Provider     string
	LastUsed     time.Time
	CostPerHour  float64
	UsagePercent float64
}

type Client interface {
	GetUnusedResources() ([]Resource, error)
	GetCosts(start, end time.Time) (float64, error)
	Initialize() error
}
