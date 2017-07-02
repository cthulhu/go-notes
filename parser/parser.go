package parser

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

// Parser constructor
// Returns parser object
func New(fixMe, todo, optimize bool) *Parser {
	if (fixMe || todo || optimize) == false {
		fixMe = true
		todo = true
		optimize = true
	}
	return &Parser{fixMe, todo, optimize, make(NotesPerFile)}
}

type NotesPerFile map[string][]string

type Parser struct {
	fixMe, todo, optimize bool
	NotesPerFile
}

var fset *token.FileSet

func init() {
	fset = token.NewFileSet() // positions are relative to fset
}

// Parses file by file path
func (p *Parser) Parse(file string) error {
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	if _, exists := p.NotesPerFile[file]; !exists {
		p.NotesPerFile[file] = []string{}
	}
	for _, comment := range f.Comments {
		if p.isNote(comment.Text()) {
			p.NotesPerFile[file] = append(p.NotesPerFile[file], comment.Text())
		}
	}
	return nil
}
func (p *Parser) isNote(node string) bool {
	return strings.HasPrefix(node, "TODO:") && p.todo ||
		strings.HasPrefix(node, "FIXME:") && p.fixMe ||
		strings.HasPrefix(node, "OPTIMIZE:") && p.optimize
}

// Aggregates parsed files information
// Returns preformated string
func (p *Parser) Aggregate() string {
	var result string
	buf := bytes.NewBufferString(result)

	for file, notes := range p.NotesPerFile {
		if len(notes) == 0 {
			continue
		}
		fmt.Fprintf(buf, "%s\n", file)
		for _, note := range notes {
			fmt.Fprintf(buf, "\t* // %s", note)
		}
		fmt.Fprintln(buf, "")
	}
	return buf.String()
}
