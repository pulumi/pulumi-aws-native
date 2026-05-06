package resources

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// pathSet stores slash-delimited metadata paths and wildcard patterns.
type pathSet struct {
	paths []string
}

// newPathSet copies a metadata path slice into a pathSet.
func newPathSet(paths []string) pathSet {
	return pathSet{paths: append([]string(nil), paths...)}
}

// add records another slash-delimited metadata path.
func (s *pathSet) add(path string) {
	s.paths = append(s.paths, path)
}

// matches reports whether path is matched by any stored metadata path.
//
// For example, a stored pattern
// "defaultActions/*/authenticateOidcConfig/clientSecret" matches the concrete
// path "defaultActions/0/authenticateOidcConfig/clientSecret".
func (s pathSet) matches(path string) bool {
	for _, pattern := range s.paths {
		if pathMatches(pattern, path) {
			return true
		}
	}
	return false
}

// covers reports whether path is equal to or nested under any stored metadata
// path pattern.
//
// For example, a stored create-only path "configurationProperties" covers
// "configurationProperties/*/type".
func (s pathSet) covers(path string) bool {
	for _, pattern := range s.paths {
		if slashPath(pattern).Contains(slashPath(path)) {
			return true
		}
	}
	return false
}

// pathMatches reports whether one metadata path pattern matches one concrete
// path with the same number of segments.
func pathMatches(pattern, path string) bool {
	return slashPath(pattern).Contains(slashPath(path)) && len(strings.Split(pattern, "/")) == len(strings.Split(path, "/"))
}

// slashPath converts a slash-delimited metadata path into a PropertyPath.
//
// Unlike slashPathForValue, this keeps every segment as a string. That preserves
// wildcard segments for PropertyPath.Contains when matching patterns like
// "rules/*/statement" against concrete metadata paths.
func slashPath(path string) resource.PropertyPath {
	segments := strings.Split(path, "/")
	result := make(resource.PropertyPath, 0, len(segments))
	for _, segment := range segments {
		result = append(result, segment)
	}
	return result
}
