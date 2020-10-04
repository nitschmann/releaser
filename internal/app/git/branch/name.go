package branch

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"

	"github.com/nitschmann/releaser/pkg/util/list"
)

// Name struct is for branch names
type Name struct {
	Delimiter string
	Prefix    string
	Suffix    string
	Title     string
}

// NewName returns a new pointer instance of Name
func NewName(delimiter string) *Name {
	return &Name{Delimiter: delimiter}
}

// FormatStringWithRegexAndDelimiter formates a String using a regex with valid chars for branch names
func (n Name) FormatStringWithRegexAndDelimiter(str string) string {
	if str == "" {
		return str
	}

	return n.ValidCharsRegex().ReplaceAllString(str, n.Delimiter)
}

func (n *Name) formatTitleWithRegexAndDelimiter() {
	n.Title = strings.ToLower(n.validTitleCharsRegex().ReplaceAllString(n.Title, n.Delimiter))
}

func (n *Name) formatStringFirstAndLastCharAlphanumeric(str string) string {
	var out string

	r := []rune(str)
	if !n.validAlphanumericCharsRegex().MatchString(string(r[0])) {
		r = append(r[:0], r[0+1:]...)
	}

	out = string(r)

	r = []rune(out)
	if !n.validAlphanumericCharsRegex().MatchString(string(r[len(out)-1])) {
		r = append(r[:len(out)-1], r[(len(out)-1)+1:]...)
	}

	out = string(r)
	return out
}

// Join joins n.Prefix, n.Title and n.Suffix together usign the given delimiter.
// Empty strings are ignored.
func (n Name) Join() string {
	parts := list.CleanEmptyStrings([]string{n.Prefix, n.Title, n.Suffix})
	name := n.formatStringFirstAndLastCharAlphanumeric(strings.Join(parts, n.Delimiter))
	return name
}

func (n Name) parseTemplateString(templatePattern string, templateValues map[string]string) (string, error) {
	var result string

	strTemplate, err := template.New("").Parse(templatePattern)
	if err != nil {
		return result, err
	}

	buf := &bytes.Buffer{}
	err = strTemplate.Execute(buf, templateValues)
	if err != nil {
		return result, err
	}

	result = buf.String()

	return result, nil
}

// SetPrefixWithTemplate sets the Prefix attribute with a Go template pattern string
func (n *Name) SetPrefixWithTemplate(prefixTemplatePattern string, templateValues map[string]string) error {
	prefix, err := n.parseTemplateString(prefixTemplatePattern, templateValues)
	if err != nil {
		return err
	}

	n.Prefix = n.FormatStringWithRegexAndDelimiter(prefix)

	return nil
}

// SetSuffixWithTemplate sets the Suffix attribute with a Go template pattern string
func (n *Name) SetSuffixWithTemplate(suffixTemplatePattern string, templateValues map[string]string) error {
	suffix, err := n.parseTemplateString(suffixTemplatePattern, templateValues)
	if err != nil {
		return err
	}

	n.Suffix = n.FormatStringWithRegexAndDelimiter(suffix)

	return nil
}

// SetTitleWithTemplate sets the Title attribute with a Go template pattern string
func (n *Name) SetTitleWithTemplate(titleTemplatePattern string, templateValues map[string]string) error {
	title, err := n.parseTemplateString(titleTemplatePattern, templateValues)
	if err != nil {
		return err
	}

	n.Title = title
	n.formatTitleWithRegexAndDelimiter()

	return nil
}

// ValidCharsRegex gives the regex which chars are allowed for branch names
func (n Name) ValidCharsRegex() *regexp.Regexp {
	reg, _ := regexp.Compile("[^a-zA-Z0-9-_+/]+")
	return reg
}

func (n Name) validTitleCharsRegex() *regexp.Regexp {
	reg, _ := regexp.Compile("[^a-zA-Z0-9-_+]+")
	return reg
}

func (n *Name) validAlphanumericCharsRegex() *regexp.Regexp {
	reg, _ := regexp.Compile("[a-zA-Z0-9]")
	return reg
}
