package html

/*
 * Handle detail action for jobs.html
 * Generate detail.html first
 */

import (
	"github.com/wutaosamuel/bterminal/utils"
)

// Detail for detail html
type Detail struct {
	Name    string // job name
	ID      string // job UUID
	Command string // Command required to run
	Crontab string // cron schedule
	Log     string // log
}

// NewDetail create new job
func NewDetail() *Detail {
	return &Detail{}
}

// GenerateDetail automatically generate
func (d *Detail) GenerateDetail(template, pattern string) (string, error) {
	templateS, err := utils.ReadHTML(template)
	if err != nil {
		return templateS, err
	}
	patternS, err := utils.ReadHTML(pattern)
	if err != nil {
		return templateS, err
	}
	// process pattern first
	p, err := utils.ReplacePattern(patternS, d)
	if err != nil {
		return "", err
	}

	// replease job html
	return utils.ReplaceHTML(templateS, 1, p)
}
