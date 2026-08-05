package role

import (
	"slices"
	"strings"
)

func normalizePermissionCodes(codes []string) []string {
	normalized := make([]string, 0, len(codes))

	for _, code := range codes {
		normalized = append(
			normalized,
			strings.ToLower(strings.TrimSpace(code)),
		)
	}

	return normalized
}

func hasDuplicatePermissionCodes(codes []string) bool {
	seen := make(map[string]struct{}, len(codes))

	for _, code := range codes {
		if _, ok := seen[code]; ok {
			return true
		}
		seen[code] = struct{}{}
	}

	return false
}

func validatePermissionCodes(codes []string) ([]string, error) {
	normalized := normalizePermissionCodes(codes)
	if slices.Contains(normalized, "") {
		return nil, ErrPermissionCodeRequired
	}
	if hasDuplicatePermissionCodes(normalized) {
		return nil, ErrDuplicatePermissionCodes
	}
	return normalized, nil
}
