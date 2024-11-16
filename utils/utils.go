package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GenerateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName,ext)

	return fmt.Sprintf("%s-%d%s",name,time.Now().UnixNano(),ext)
}