// +build dev

package templates

import "net/http"

var Templates http.FileSystem = http.Dir(".")
