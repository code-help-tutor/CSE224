WeChat: cstutorcs
QQ: 749389476
Email: tutorcs@163.com
package tritonhttp

import (
	"mime"
	"net/textproto"
	"time"
)

// CanonicalHeaderKey returns the canonical format of the
// header key s. The canonicalization converts the first
// letter and any letter following a hyphen to upper case;
// the rest are converted to lowercase. For example, the
// canonical key for "content-type" is "Content-Type".
// You should store header keys in the canonical format
// in internal data structures.
func CanonicalHeaderKey(s string) string {
	return textproto.CanonicalMIMEHeaderKey(s)
}

// FormatTime formats time according to the HTTP spec.
// It is like time.RFC1123 but hard-codes GMT as the time zone.
// You should use this function for the "Date" and "Last-Modified"
// headers.
func FormatTime(t time.Time) string {
	s := t.UTC().Format(time.RFC1123)
	s = s[:len(s)-3] + "GMT"
	return s
}

// MIMETypeByExtension returns the MIME type associated with the
// file extension ext. The extension ext should begin with a
// leading dot, as in ".html". When ext has no associated type,
// MIMETypeByExtension returns "".
// You should use this function for the "Content-Type" header.
func MIMETypeByExtension(ext string) string {
	return mime.TypeByExtension(ext)
}
