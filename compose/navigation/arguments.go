package navigation

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/zodimo/go-maybe"
)

// NavType represents the type of a navigation argument
type NavType int

const (
	NavTypeString NavType = iota
	NavTypeInt
	NavTypeBool
)

// NavArgumentSpec defines the schema for a navigation argument (like Kotlin's NavArgument)
// This is the metadata/definition, not the runtime value
type NavArgumentSpec struct {
	Type                  NavType
	IsNullable            bool
	DefaultValue          maybe.Maybe[any]
	IsDefaultValuePresent bool
}

// NavArgumentSpecOption configures a NavArgumentSpec
type NavArgumentSpecOption func(*NavArgumentSpec)

// WithNullable sets whether the argument is nullable
func WithNullable(nullable bool) NavArgumentSpecOption {
	return func(s *NavArgumentSpec) {
		s.IsNullable = nullable
	}
}

// WithDefaultValue sets the default value for the argument
func WithDefaultValue(value any) NavArgumentSpecOption {
	return func(s *NavArgumentSpec) {
		s.DefaultValue = maybe.Some[any](value)
		s.IsDefaultValuePresent = true
	}
}

// NewNavArgumentSpec creates a new argument spec with the given type and options
func NewNavArgumentSpec(navType NavType, opts ...NavArgumentSpecOption) NavArgumentSpec {
	spec := NavArgumentSpec{
		Type:                  navType,
		IsNullable:            false,
		DefaultValue:          maybe.None[any](),
		IsDefaultValuePresent: false,
	}
	for _, opt := range opts {
		opt(&spec)
	}
	return spec
}

// NavArguments holds the runtime argument values (like Kotlin's SavedState)
type NavArguments map[string]any

// GetString retrieves a string argument
func (a NavArguments) GetString(key string) (string, bool) {
	if a == nil {
		return "", false
	}
	val, ok := a[key]
	if !ok {
		return "", false
	}
	str, ok := val.(string)
	return str, ok
}

// GetInt retrieves an integer argument
func (a NavArguments) GetInt(key string) (int, bool) {
	if a == nil {
		return 0, false
	}
	val, ok := a[key]
	if !ok {
		return 0, false
	}
	switch v := val.(type) {
	case int:
		return v, true
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, false
		}
		return i, true
	default:
		return 0, false
	}
}

// GetBool retrieves a boolean argument
func (a NavArguments) GetBool(key string) (bool, bool) {
	if a == nil {
		return false, false
	}
	val, ok := a[key]
	if !ok {
		return false, false
	}
	switch v := val.(type) {
	case bool:
		return v, true
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return false, false
		}
		return b, true
	default:
		return false, false
	}
}

// placeholderPattern matches {argName} in route patterns
var placeholderPattern = regexp.MustCompile(`\{([^}]+)\}`)

// parseRoutePattern extracts path and query argument names from a route pattern
// e.g., "details/{id}?name={name}" → pathArgs: ["id"], queryArgs: ["name"]
func parseRoutePattern(pattern string) (pathArgs []string, queryArgs []string) {
	// Split by ? to separate path from query
	parts := strings.SplitN(pattern, "?", 2)
	pathPart := parts[0]

	// Extract path arguments
	matches := placeholderPattern.FindAllStringSubmatch(pathPart, -1)
	for _, match := range matches {
		if len(match) > 1 {
			pathArgs = append(pathArgs, match[1])
		}
	}

	// Extract query arguments if present
	if len(parts) > 1 {
		queryPart := parts[1]
		matches = placeholderPattern.FindAllStringSubmatch(queryPart, -1)
		for _, match := range matches {
			if len(match) > 1 {
				queryArgs = append(queryArgs, match[1])
			}
		}
	}

	return pathArgs, queryArgs
}

// matchRoute checks if a route matches a pattern and extracts arguments
// e.g., matchRoute("details/{id}?filter={filter}", "details/123?filter=active")
//
//	→ NavArguments{"id": "123", "filter": "active"}, true
func matchRoute(pattern, route string) (NavArguments, bool) {
	args := make(NavArguments)

	// Split pattern and route by ?
	patternParts := strings.SplitN(pattern, "?", 2)
	routeParts := strings.SplitN(route, "?", 2)

	patternPath := patternParts[0]
	routePath := routeParts[0]

	// Match path segments
	if !matchPath(patternPath, routePath, args) {
		return nil, false
	}

	// Match query parameters if pattern has them
	if len(patternParts) > 1 {
		patternQuery := patternParts[1]
		routeQuery := ""
		if len(routeParts) > 1 {
			routeQuery = routeParts[1]
		}
		matchQuery(patternQuery, routeQuery, args)
	}

	return args, true
}

// matchPath matches the path portion and extracts arguments
func matchPath(patternPath, routePath string, args NavArguments) bool {
	patternSegments := strings.Split(patternPath, "/")
	routeSegments := strings.Split(routePath, "/")

	if len(patternSegments) != len(routeSegments) {
		return false
	}

	for i, patternSeg := range patternSegments {
		routeSeg := routeSegments[i]

		if strings.HasPrefix(patternSeg, "{") && strings.HasSuffix(patternSeg, "}") {
			// This is a placeholder - extract the argument name and value
			argName := patternSeg[1 : len(patternSeg)-1]
			// URL-decode the value to handle special characters
			decodedValue, err := url.PathUnescape(routeSeg)
			if err != nil {
				decodedValue = routeSeg // Fallback to raw value if decoding fails
			}
			args[argName] = decodedValue
		} else if patternSeg != routeSeg {
			// Static segment doesn't match
			return false
		}
	}

	return true
}

// matchQuery matches query parameters and extracts arguments
func matchQuery(patternQuery, routeQuery string, args NavArguments) {
	// Parse route query parameters
	routeParams, _ := url.ParseQuery(routeQuery)

	// Parse pattern to find expected query args
	patternParams := strings.Split(patternQuery, "&")
	for _, param := range patternParams {
		kv := strings.SplitN(param, "=", 2)
		if len(kv) != 2 {
			continue
		}
		paramName := kv[0]
		paramValue := kv[1]

		// Check if the value is a placeholder
		if strings.HasPrefix(paramValue, "{") && strings.HasSuffix(paramValue, "}") {
			argName := paramValue[1 : len(paramValue)-1]
			// Get value from route query params
			if values, ok := routeParams[paramName]; ok && len(values) > 0 {
				args[argName] = values[0]
			}
			// Note: if not present, it's optional (query args are optional)
		}
	}
}

// buildRoute constructs a route from pattern and arguments
// Values are URL-encoded to handle special characters like /, spaces, &, etc.
// e.g., buildRoute("details/{id}", NavArguments{"id": "hello world"}) → "details/hello%20world"
func buildRoute(pattern string, args NavArguments) string {
	// Split by ? to handle path and query separately
	parts := strings.SplitN(pattern, "?", 2)
	pathPart := parts[0]

	// Replace placeholders in path with URL-encoded values
	pathResult := placeholderPattern.ReplaceAllStringFunc(pathPart, func(match string) string {
		argName := match[1 : len(match)-1]
		if val, ok := args[argName]; ok {
			// URL-encode path segment to handle special characters
			return url.PathEscape(fmt.Sprintf("%v", val))
		}
		return match
	})

	// Handle query part if present
	if len(parts) > 1 {
		queryPart := parts[1]
		// Replace placeholders in query with URL-encoded values
		queryResult := placeholderPattern.ReplaceAllStringFunc(queryPart, func(match string) string {
			argName := match[1 : len(match)-1]
			if val, ok := args[argName]; ok {
				// URL-encode query value
				return url.QueryEscape(fmt.Sprintf("%v", val))
			}
			return match
		})
		return pathResult + "?" + queryResult
	}

	return pathResult
}
