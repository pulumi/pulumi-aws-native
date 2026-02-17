// Copyright 2024, Pulumi Corporation.

// Known issue: AWS CloudFormation schemas contain 'imcp' typo (should be 'icmp')
// in EC2 SecurityGroup propertyTransform expressions. This is upstream in AWS.
package resources

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/blues/jsonata-go"
	"github.com/golang/glog"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// TransformCache holds compiled JSONata expressions per resource type.
// It provides thread-safe lazy compilation of expressions.
type TransformCache struct {
	mu         sync.RWMutex
	transforms map[string][]CompiledTransform
}

// CompiledTransform holds a pre-compiled JSONata expression for a property path.
type CompiledTransform struct {
	// Path is the SDK property path (e.g., "securityGroupEgress/*/ipProtocol")
	Path string

	// Expression is the compiled JSONata expression
	Expression *jsonata.Expr

	// RawExpr is the original JSONata expression string (for debugging)
	RawExpr string

	// PathMatcher helps match concrete paths against wildcard patterns
	PathMatcher *PathMatcher
}

// PathMatcher helps match concrete property paths against patterns with wildcards.
// Pattern: "securityGroupEgress/*/fromPort" matches "securityGroupEgress/0/fromPort"
type PathMatcher struct {
	segments    []string
	hasWildcard bool
}

// NewPathMatcher creates a PathMatcher from a pattern string.
func NewPathMatcher(pattern string) *PathMatcher {
	segments := strings.Split(pattern, "/")
	hasWildcard := false
	for _, s := range segments {
		if s == "*" {
			hasWildcard = true
			break
		}
	}
	return &PathMatcher{
		segments:    segments,
		hasWildcard: hasWildcard,
	}
}

// Match checks if a concrete path matches this pattern.
// Pattern "a/*/b" matches "a/0/b", "a/1/b", etc.
func (pm *PathMatcher) Match(concretePath string) bool {
	concreteSegments := strings.Split(concretePath, "/")
	if len(concreteSegments) != len(pm.segments) {
		return false
	}

	for i, patternSeg := range pm.segments {
		if patternSeg == "*" {
			// Wildcard matches any segment (typically array index)
			continue
		}
		if patternSeg != concreteSegments[i] {
			return false
		}
	}
	return true
}

// HasWildcard returns true if this pattern contains wildcards.
func (pm *PathMatcher) HasWildcard() bool {
	return pm.hasWildcard
}

// GetPropertyName returns the last segment of the path (the property name).
func (pm *PathMatcher) GetPropertyName() string {
	if len(pm.segments) == 0 {
		return ""
	}
	return pm.segments[len(pm.segments)-1]
}

// NewTransformCache creates a new TransformCache instance.
// Prefer creating a cache per-provider instance rather than using GlobalTransformCache.
func NewTransformCache() *TransformCache {
	return &TransformCache{
		transforms: make(map[string][]CompiledTransform),
	}
}

// GlobalTransformCache is a singleton cache for compiled transforms.
// Deprecated: Use NewTransformCache() and dependency injection instead.
// This global is kept for backward compatibility with tests.
var GlobalTransformCache = NewTransformCache()

// GetOrCompile returns compiled transforms for a resource type, compiling on first access.
// Returns nil if the resource has no property transforms.
func (tc *TransformCache) GetOrCompile(resourceToken string, spec *metadata.CloudAPIResource) []CompiledTransform {
	if spec == nil || len(spec.PropertyTransforms) == 0 {
		return nil
	}

	// Fast path: check if already compiled
	tc.mu.RLock()
	if transforms, ok := tc.transforms[resourceToken]; ok {
		tc.mu.RUnlock()
		return transforms
	}
	tc.mu.RUnlock()

	// Slow path: compile and cache
	tc.mu.Lock()
	defer tc.mu.Unlock()

	// Double-check after acquiring write lock
	if transforms, ok := tc.transforms[resourceToken]; ok {
		return transforms
	}

	transforms := tc.compile(resourceToken, spec.PropertyTransforms)
	tc.transforms[resourceToken] = transforms
	return transforms
}

// compile compiles all property transforms for a resource.
func (tc *TransformCache) compile(resourceToken string, transformSpecs map[string]string) []CompiledTransform {
	var transforms []CompiledTransform

	for path, exprStr := range transformSpecs {
		expr, err := jsonata.Compile(exprStr)
		if err != nil {
			// Log warning and skip this transform
			glog.Warningf("Failed to compile JSONata expression for %s path %s: %v", resourceToken, path, err)
			continue
		}

		transforms = append(transforms, CompiledTransform{
			Path:        path,
			Expression:  expr,
			RawExpr:     exprStr,
			PathMatcher: NewPathMatcher(path),
		})
	}

	return transforms
}

// EvaluateTransform evaluates a JSONata expression against a property value.
// The context map provides the full object context for expressions that reference
// sibling properties (e.g., the IpProtocol lookup in SecurityGroup transforms).
//
// Returns the transformed value and any error. ErrUndefined from JSONata is treated
// as "no transformation needed" and returns the original value. Nil results are also
// treated as "no transformation" since functions like $lookup return nil for missing keys.
func EvaluateTransform(transform CompiledTransform, value interface{}, context map[string]interface{}) (interface{}, error) {
	// Build the evaluation context by merging the value with sibling context
	evalContext := make(map[string]interface{})
	for k, v := range context {
		evalContext[k] = v
	}

	// The property being transformed is typically referenced by its CloudFormation name
	// in the expression. We need to provide the context that includes sibling properties.
	result, err := transform.Expression.Eval(evalContext)
	if err != nil {
		// Check for undefined - this means the transform doesn't apply
		if isUndefinedError(err) {
			glog.V(9).Infof("JSONata transform returned undefined for path %s, keeping original value", transform.Path)
			return value, nil
		}
		return value, fmt.Errorf("JSONata evaluation failed: %w", err)
	}

	// If result is nil (e.g., $lookup didn't find the key), return original value
	if result == nil {
		glog.V(9).Infof("JSONata transform returned nil for path %s, keeping original value", transform.Path)
		return value, nil
	}

	return result, nil
}

// isUndefinedError checks if an error is JSONata's "undefined" result.
// FRAGILE: This relies on string matching because jsonata-go doesn't export typed errors.
// If the library changes error messages, this may break silently. Monitor for issues after
// jsonata-go dependency updates.
func isUndefinedError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	// JSONata-go returns "no results found" when expression evaluates to undefined.
	// Also check for "undefined" for forward compatibility with potential library changes.
	return strings.Contains(errMsg, "undefined") || strings.Contains(errMsg, "no results found")
}

// regexCache caches compiled regex patterns to avoid recompilation in hot paths.
var regexCache = &sync.Map{}

// getOrCompileRegex retrieves a cached compiled regex or compiles and caches it.
func getOrCompileRegex(pattern string) (*regexp.Regexp, error) {
	fullPattern := "^" + pattern + "$"
	if cached, ok := regexCache.Load(fullPattern); ok {
		return cached.(*regexp.Regexp), nil
	}
	re, err := regexp.Compile(fullPattern)
	if err != nil {
		return nil, err
	}
	regexCache.Store(fullPattern, re)
	return re, nil
}

// ValuesEquivalent compares two values for semantic equivalence after applying
// CloudFormation propertyTransform normalization.
//
// Comparison strategies (in order of precedence):
//
//  1. Pulumi DeepEquals - Primary comparison using Pulumi's PropertyValue equality.
//     Handles type normalization (e.g., int(6) and float64(6.0) both become NumberProperty).
//
//  2. Regex pattern matching - For CloudFormation transforms that produce regex patterns.
//     Example transforms that produce patterns:
//     - engineVersion: $join([$string(EngineVersion), ".*"]) → "5.8.*" matches "5.8.1"
//     - principal: $join(["^arn:aws[a-zA-Z-]*:iam::",Principal,":[a-zA-Z-]*"]) → regex
func ValuesEquivalent(original, transformedNew interface{}) bool {
	// 1. Primary: Use Pulumi's DeepEquals for type-safe comparison.
	// This handles numeric type variations correctly (int vs float64).
	origPV := resource.NewPropertyValue(original)
	newPV := resource.NewPropertyValue(transformedNew)
	if origPV.DeepEquals(newPV) {
		return true
	}

	// 2. Secondary: Regex pattern matching for CloudFormation transforms.
	// Some transforms produce regex patterns to match version ranges or ARN formats.
	if pattern, ok := transformedNew.(string); ok {
		// Remove wrapping quotes if present (some CF transforms produce quoted regexes)
		pattern = strings.Trim(pattern, "\"")

		// Try to match as regex if it looks like a pattern
		if looksLikeRegex(pattern) {
			if originalStr, ok := original.(string); ok {
				re, err := getOrCompileRegex(pattern)
				if err != nil {
					glog.V(9).Infof("Failed to compile regex pattern %q from propertyTransform: %v", pattern, err)
					return false
				}
				if re.MatchString(originalStr) {
					return true
				}
			}
		}
	}

	return false
}

// looksLikeRegex returns true if the string appears to be a regex pattern
// rather than a literal value. Used to avoid compiling non-regex strings.
func looksLikeRegex(s string) bool {
	return strings.Contains(s, ".*") || strings.Contains(s, ".+") ||
		strings.Contains(s, "[") || strings.Contains(s, "\\")
}

// FindTransformForPath finds a compiled transform that matches the given concrete path.
// Returns nil if no matching transform is found.
func FindTransformForPath(transforms []CompiledTransform, concretePath string) *CompiledTransform {
	for i := range transforms {
		if transforms[i].PathMatcher.Match(concretePath) {
			return &transforms[i]
		}
	}
	return nil
}

// ExtractPropertyContext extracts the context object for evaluating a transform.
// For array element properties like "securityGroupEgress/0/ipProtocol", this
// returns the parent array element (securityGroupEgress[0]) as a map.
//
// The context is needed because JSONata expressions often reference sibling
// properties (e.g., the IpProtocol transform references both IpProtocol and FromPort).
func ExtractPropertyContext(props map[string]interface{}, path string) map[string]interface{} {
	segments := strings.Split(path, "/")
	if len(segments) <= 1 {
		// Top-level property, context is the full props map
		return props
	}

	// Navigate to the parent context
	current := interface{}(props)
	for i := 0; i < len(segments)-1; i++ {
		seg := segments[i]

		switch v := current.(type) {
		case map[string]interface{}:
			next, ok := v[seg]
			if !ok {
				return nil
			}
			current = next
		case []interface{}:
			// Parse array index
			var idx int
			if _, err := fmt.Sscanf(seg, "%d", &idx); err != nil || idx < 0 || idx >= len(v) {
				glog.V(9).Infof("Failed to parse array index %q in path %s: invalid format or out of bounds (array len=%d)", seg, path, len(v))
				return nil
			}
			current = v[idx]
		default:
			return nil
		}
	}

	// The parent context should be a map
	if ctx, ok := current.(map[string]interface{}); ok {
		return ctx
	}
	return nil
}
