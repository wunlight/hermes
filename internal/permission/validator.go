package permission

import "regexp"

var permissionCodePattern = regexp.MustCompile(
	`^[a-z]+(\.[a-z]+)+$`,
)

func isValidPermissionCode(code string) bool {
	return permissionCodePattern.MatchString(code)
}
