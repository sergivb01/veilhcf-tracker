package message

import (
	"regexp"
)

var ChatRegex = regexp.MustCompile("\\[(.*?)\\] \\(\\d+\\) \\w+: \\w+")
var KillRegex = regexp.MustCompile("\\w+\\[\\d+\\] was slain by \\w+\\[\\d+\\] using (.)+\\.")
