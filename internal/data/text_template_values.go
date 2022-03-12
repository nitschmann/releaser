package data

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/nitschmann/releaser/pkg/git"
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
	// Flags custom defined
	Flags map[string]string
	// Git specific values
	GitRemote      string
	GitRepoHttpURL string
	GitRepoName    string
	// GitCommitLogs with
	GitCommitLogs []git.CommitLog
	// UserHomeDir is the home direcotry path of the current user
	UserHomeDir string
	// Type is the specified type flag
	Type string
	// Release specific values
	ReleaseTag    string
	ReleaseTarget string
}

// NewTextTemplateValues returns a new pointer instance of TextTemplateValues with default values
func NewTextTemplateValues() *TextTemplateValues {
	dateTime := time.Now()
	userHomeDir, _ := os.UserHomeDir()

	obj := &TextTemplateValues{
		DTYear:      strconv.Itoa(dateTime.Year()),
		DTMonth:     strconv.Itoa(int(dateTime.Month())),
		Flags:       make(map[string]string),
		UserHomeDir: userHomeDir,
	}

	obj.DTMonth = obj.dateNumberStringWithPad(int(dateTime.Month()))
	obj.DTDay = obj.dateNumberStringWithPad(dateTime.Day())
	obj.DTHour = strconv.Itoa(dateTime.Hour())
	obj.DTMinute = obj.dateNumberStringWithPad(dateTime.Minute())
	obj.DTSecond = obj.dateNumberStringWithPad(dateTime.Second())

	return obj
}

// AddFlag addes or overwrites an new entry into the Flag field under the given key
func (ttv *TextTemplateValues) AddFlag(key, value string) {
	ttv.Flags[key] = value
}

func (ttv *TextTemplateValues) helperFunctions() template.FuncMap {
	return template.FuncMap{
		"ToLower": strings.ToLower,
		"ToUpper": strings.ToUpper,
		"ToTitle": strings.Title,
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

func (ttv *TextTemplateValues) dateNumberStringWithPad(num int) string {
	return fmt.Sprintf("%02d", num)
}
