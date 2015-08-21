package hexseeker

import (
	"io"
)

type Hexer struct{ io.ReadSeeker }

