package autocpp

// Pacman implements the PackageSystem interface

type Pacman struct {
}

func (pacman *Pacman) PackagesProvides(shortIncludeName string) ([]string, error) {
	return []string{}, nil
}

func (pacman *Pacman) IncludePathToCXXFlags(string) string {
	return ""
}
