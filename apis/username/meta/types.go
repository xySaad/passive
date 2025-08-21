package meta

import "regexp"

type Resp struct {
	Payload struct {
		Payloads map[string]struct {
			Result struct {
				Type    string `json:"type"`
				Exports struct {
					Meta struct {
						Title string `json:"title"`
					} `json:"meta"`
				} `json:"exports"`
			} `json:"result"`
		}
	} `json:"payload"`
}

var patterns = map[string]string{
	"lsd":      `"lsd":{"name":"lsd","value":"(\w+)"}`,
	"jazoest":  `"jazoest":{"name":"jazoest","value":"(\w+)"}`,
	"__spin_r": `"__spin_r":(\d+)`,
	"__spin_b": `"__spin_b":"(\w+)"`,
	"__ccg":    `{"connectionClass":"(\w+)"}`,
	"__rev":    `{"consistency":{"rev":(\w+)}`,
	"__hsi":    `"hsi":"(\w+)"`,
}

var CompiledRegex = map[string]*regexp.Regexp{}

func init() {
	for name, patt := range patterns {
		CompiledRegex[name] = regexp.MustCompile(patt)
	}
}
