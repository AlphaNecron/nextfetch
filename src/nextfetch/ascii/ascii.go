package ascii

import _ "embed"

var (
	//go:embed arch_linux
	arch_linux string

	//go:embed windows
	windows string

	//go:embed generic
	generic string

	ARTS = map[string]Art{
		"Arch Linux": create(arch_linux, "blue"),
		"Windows":    create(windows, "blue"),
		"Generic":    create(generic, "none"),
	}
)

type Art struct {
	Art   string
	Color string
}

func create(art string, color string) Art {
	return Art{art, color}
}
