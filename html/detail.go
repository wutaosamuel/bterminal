package html

/*
 * Handle detail action for jobs.html
 * Generate detail.html first
 */

import (
	"../utils"
)

// Detail for detail html
type Detail struct {
	Name    string // job name
	ID      string // job UUID
	Command string // Command required to run
	Crontab string // cron schedule
	LogName string // log
}

// NewDetail create new job
func NewDetail() *Detail {
	return &Detail{}
}

// GenerateDetail automatically generate
func (d *Detail) GenerateDetail(template, pattern string) (string, error) {
	// process pattern first
	p, err := utils.ReplacePattern(pattern, d)
	if err != nil {
		return "", err
	}

	// replease job html
	return utils.ReplaceHTML(template, 1, p)
}
