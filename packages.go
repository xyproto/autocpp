package autocpp

type PackageSystem interface {
	PackageProvides(string) ([]string, error)
	IncludePathToCXXFlags(string) string
}
