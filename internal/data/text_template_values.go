package data

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// TextTemplateValues has data fields which could be used in text templates
// (https://golangforall.com/en/post/templates.html)
type TextTemplateValues struct {
	// BranchName is the specified name of a new git branch
	BranchName string
	// CommitMessage  is the given message of a commit
	CommitMessage string
	// Datetime specific values
	DTYear   string
	DTMonth  string
	DTDay    string
	DTHour   string
	DTMinute string
	DTSecond string
	// UserHomeDir is the home direcotry path of the current user
	UserHomeDir string
	// Release specific values
	ReleaseNewVersion     string
	ReleaseCurrentVersion string
	// Type is the specified type flag
	Type string
	// Flags custom defined
	Flags map[string]string
}

// NewTextTemplateValues returns a new pointer instance of TextTemplateValues with default values
func NewTextTemplateValues() *TextTemplateValues {
	dateTime := time.Now()
	userHomeDir, _ := os.UserHomeDir()

	return &TextTemplateValues{
		DTYear:      strconv.Itoa(dateTime.Year()),
		DTMonth:     strconv.Itoa(int(dateTime.Month())),
		DTDay:       strconv.Itoa(dateTime.Day()),
		DTHour:      strconv.Itoa(dateTime.Hour()),
		DTMinute:    strconv.Itoa(dateTime.Minute()),
		DTSecond:    strconv.Itoa(dateTime.Second()),
		Flags:       make(map[string]string),
		UserHomeDir: userHomeDir,
	}
}

// AddFlag addes or overwrites an new entry into the Flag field under the given key
func (ttv *TextTemplateValues) AddFlag(key, value string) {
	ttv.Flags[key] = value
}

func (ttv *TextTemplateValues) helperFunctions() template.FuncMap {
	return template.FuncMap{
		"ToLower": strings.ToLower,
		"ToUpper": strings.ToUpper,
		"ToTitle": strings.ToTitle,
	}
}

// ParseTemplateString uses a given template string, parses it with the specified fields in the struct
// and returns the result as string. An error is returned in cases of any error during the process.
func (ttv *TextTemplateValues) ParseTemplateString(templateStr string) (string, error) {
	t, err := template.New("test").Funcs(ttv.helperFunctions()).Parse(templateStr)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, ttv)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
