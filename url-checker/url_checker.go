package url_checker

import (
	"github.com/asaskevich/govalidator"
	"strings"
)

func ValidateUrl(url string) bool {
	return govalidator.IsURL(url)
}

func StartsWith(firstUrl, secondUrl string) bool {
	return strings.HasPrefix(secondUrl, firstUrl)
}
