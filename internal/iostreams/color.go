package iostreams

import "os"

func EnvColorDisabled() bool {
	// See: https://bixense.com/clicolors/
	return os.Getenv("NO_COLOR") != "" || os.Getenv("CLICOLOR") == "0"
}

func EnvColorForced() bool {
	// See: https://bixense.com/clicolors/
	return os.Getenv("CLICOLOR_FORCE") != "" && os.Getenv("CLICOLOR_FORCE") != "0"
}
