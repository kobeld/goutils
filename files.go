package goutils

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	EXT_DOT      = "."
	DEF_KIND     = ""
	B            = "B"
	KB           = "KB"
	MB           = "MB"
	GB           = "GB"
	LKB          = 1024
	LMB          = 1024 * 1024
	LGB          = 1024 * 1024 * 1024
	FILE_PREVIEW = "s1"
	IMG          = "image"
	PRESENT      = "presentation"
	PDF          = "PDF"
	VIDEO        = "video"
	AUDIO        = "audio"
	TXT          = "txt"
	DOC          = "doc"
	DOCUMENT     = "document"
)

var EXT_MAP = map[string]string{
	"bmp":      IMG,
	"png":      IMG,
	"jpg":      IMG,
	"jpeg":     IMG,
	"gif":      IMG,
	"svg":      IMG,
	"raw":      IMG,
	"idraw":    IMG,
	"eps":      IMG,
	"key":      DOC,
	"ppt":      PRESENT,
	"pptx":     PRESENT,
	"pps":      PRESENT,
	"pdf":      PDF,
	"mov":      VIDEO,
	"m4v":      VIDEO,
	"mp4":      VIDEO,
	"mpg":      VIDEO,
	"avi":      VIDEO,
	"wmv":      VIDEO,
	"flv":      VIDEO,
	"wav":      AUDIO,
	"mp3":      AUDIO,
	"m4a":      AUDIO,
	"flac":     AUDIO,
	"aac":      AUDIO,
	"swf":      VIDEO,
	"txt":      DOC,
	"doc":      DOCUMENT,
	"docx":     DOCUMENT,
	"log":      DOC,
	"numbers":  DOC,
	"pages":    DOC,
	"rtf":      DOC,
	"sql":      DOC,
	"xls":      DOCUMENT,
	"xlsx":     DOCUMENT,
	"psd":      "",
	"tiff":     "",
	"tif":      "",
	"ai":       "",
	"zip":      "",
	"rb":       "",
	"aiff":     "",
	"aif":      "",
	"bz2":      "",
	"css":      "",
	"csv":      "",
	"dmg":      "",
	"graffle":  "",
	"gz":       "",
	"html":     "",
	"img":      "",
	"lha":      "",
	"markdown": "",
	"md":       "",
	"midi":     "",
	"mkv":      "",
	"mod":      "",
	"py":       "",
	"rar":      "",
	"tar":      "",
	"tga":      "",
	"tgz":      "",
}

func GetFileSize(contentLength int64) string {

	switch {
	case contentLength == 0:
		return "0" + B
	case contentLength < LKB:
		return strconv.FormatInt(contentLength, 10) + B
	case contentLength < LMB:
		return fmt.Sprintf("%.02f", float64(contentLength)/float64(LKB)) + KB
	case contentLength < LGB:
		return fmt.Sprintf("%.02f", float64(contentLength)/float64(LMB)) + MB
	case contentLength >= LGB:
		return fmt.Sprintf("%.02f", float64(contentLength)/float64(LGB)) + GB
	}
	return ""
}

func IsImageType(contentType string) bool {
	ct := strings.TrimSpace(contentType)
	if ct == "image/png" || ct == "image/jpeg" || ct == "image/gif" || ct == "image/x-png" || ct == "image/pjpeg" || ct == "image/bmp" {
		return true
	}

	return false
}
