package autocpp

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Sources struct {
	verbose            bool
	rootPath           string
	absFilenamesHeader []string
	absFilenamesCPP    []string
	absFilenamesC      []string
	entireSource       []byte
	foundMap           map[string]string // from a short include name to the full path, if the include was found
}

func NewSources(rootPath string, verbose bool) (*Sources, error) {
	var src Sources
	src.rootPath = rootPath
	src.verbose = verbose
	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		switch strings.ToLower(filepath.Ext(path)) {
		case ".h", ".hpp", ".hh", ".h++":
			if verbose {
				fmt.Printf("added: %q\n", path)
			}
			src.absFilenamesHeader = append(src.absFilenamesHeader, path)
		case ".c":
			if verbose {
				fmt.Printf("added: %q\n", path)
			}
			src.absFilenamesC = append(src.absFilenamesC, path)
		case ".cpp", ".cc", ".cxx", ".c++":
			if verbose {
				fmt.Printf("added: %q\n", path)
			}
			src.absFilenamesCPP = append(src.absFilenamesCPP, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	src.ReadAll()
	src.foundMap = make(map[string]string)
	return &src, nil
}

func (src *Sources) AllFilenames() []string {
	var allFilenames []string
	allFilenames = append(allFilenames, src.absFilenamesHeader...)
	allFilenames = append(allFilenames, src.absFilenamesC...)
	allFilenames = append(allFilenames, src.absFilenamesCPP...)
	return allFilenames
}

func (src *Sources) ReadAll() error {
	allFilenames := src.AllFilenames()
	lenall := len(allFilenames)
	for i, path := range allFilenames {
		if src.verbose {
			fmt.Printf("[%d/%d, %.2f%%] Reading %s...\n", i+1, lenall, math.Round((float64(i+1)*100.0)/float64(lenall)), path)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		src.entireSource = append(src.entireSource, []byte("\n")...)
		src.entireSource = append(src.entireSource, data...)
	}
	return nil
}

func (src *Sources) String() string {
	return string(src.entireSource)
}

func ForEachTrimmedLine(data []byte, f func(string) error) error {
	for _, line := range strings.Split(string(data), "\n") {
		if err := f(strings.TrimSpace(line)); err != nil {
			return err
		}
	}
	return nil
}

func (src *Sources) IncludeLines() []string {
	var includes []string
	ForEachTrimmedLine(src.entireSource, func(trimmedLine string) error {
		if strings.HasPrefix(trimmedLine, "#include") {
			includes = append(includes, trimmedLine)
		}
		return nil
	})
	return includes
}

func hasS(xs []string, e string) bool {
	for _, x := range xs {
		if x == e {
			return true
		}
	}
	return false
}

// ShortIncludes returns a slice of unique and sorted include filenames, as written in the header files,
// but without the surrounding "#include <...>" or "#include \"...\"".
func (src *Sources) ShortIncludes() []string {
	var includes []string
	for _, includeLine := range src.IncludeLines() {
		include := strings.TrimPrefix(includeLine, "#include")
		include = strings.TrimSpace(include)
		include = strings.TrimPrefix(include, "\"")
		include = strings.TrimPrefix(include, "<")
		include = strings.TrimSuffix(include, "\"")
		include = strings.TrimSuffix(include, ">")
		if !hasS(includes, include) {
			includes = append(includes, include)
		}
	}
	sort.Strings(includes)
	return includes
}

// exists checks if the given path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func shortest(xs []string) string {
	minlen := -1
	s := ""
	for _, x := range xs {
		if minlen == -1 || len(x) < minlen {
			minlen = len(x)
			s = x
		}
	}
	if minlen == -1 {
		return "" // none were shortest, no elements
	}
	return s
}

// shotestButPreferKeyword tries to find the shortest string in the given slice,
// but if a candidate has the given keyword, prefer that over shorter strings.
// xs is expected to be sorted with sort.Strings first.
func shortestButPreferKeyword(xs []string, keyword string) string {
	minlen := -1
	s := ""
	hasKeyword := false
	foundKeywordAlready := false
	// loop in reverse order, assume xs is sorted and has the lowest version numbers first
	for i := len(xs) - 1; i > 0; i-- {
		x := xs[i]
		hasKeyword = strings.Contains(x, keyword)
		foundKeywordAlready = strings.Contains(s, keyword)
		if (minlen == -1) || !foundKeywordAlready || (hasKeyword && len(x) < minlen) {
			minlen = len(x)
			s = x
		}
	}
	if minlen == -1 {
		return "" // none were shortest, no elements
	}
	return s
}

// FindIncludePaths fills src.foundMap with short include names and their corresponding paths.
// It will also return a slice of the short include names that were not found.
func (src *Sources) FindIncludePaths(locsys *LocalSystem) []string {
	var notFound []string
OUT:
	for _, include := range src.ShortIncludes() {
		// First search system directories
		for _, includeDirectory := range locsys.systemIncludeDirectories {
			path := filepath.Join(includeDirectory, include)
			if hasS(locsys.includeFiles, path) {
				src.foundMap[include] = path
				continue OUT
			}
		}
		// Then search local directories
		for _, includeDirectory := range locsys.localIncludeDirectories {
			path := filepath.Join(includeDirectory, include)
			if exists(path) {
				src.foundMap[include] = path
				continue OUT
			}
		}
		// Then look for candidates in the include files that has been found on the system
		var candidates []string
		for _, includeFile := range locsys.includeFiles {
			if strings.HasSuffix(includeFile, "/"+include) {
				candidates = append(candidates, includeFile)
			}
		}
		if len(candidates) == 0 {
			notFound = append(notFound, include)
			continue
		}
		sort.Strings(candidates)
		// candidates are now sorted
		if src.verbose {
			fmt.Printf("Candidates for %s:\n", include)
			for _, candidate := range candidates {
				fmt.Printf("\t%s\n", candidate)
			}
		}
		path := shortestButPreferKeyword(candidates, "++")
		if src.verbose {
			fmt.Printf("\tChose:\n\t%s\n", path)
		}
		src.foundMap[include] = path
	}
	return notFound
}

func (src *Sources) FindAndPrintIncludePaths(locsys *LocalSystem) {
	notFound := src.FindIncludePaths(locsys)
	for _, path := range src.foundMap {
		fmt.Printf("FOUND: %s\n", path)
	}
	for _, include := range notFound {
		fmt.Printf("NOT FOUND: %s\n", include)
	}
}

// Thread returns true if "-ldl -pthread -lpthread" should be added
func (src *Sources) Thread() bool {
	for _, include := range src.ShortIncludes() {
		switch include {
		case "condition_variable", "future", "mutex", "new", "pthread.h", "thread", "dlfcn.h":
			return true
		}
	}
	return false
}
