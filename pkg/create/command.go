package create

type Command struct {
	Command string `json:"command"`
	Alias string `json:"alias,omitempty"`
	Template string `json:"template,omitempty"`
	Subcommands []*Command `json:"subcommands,omitempty"`
	File string `json:"file,omitempty"`
}