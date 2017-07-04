package parser

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

// New - method to create a parser
// Returns parser object
func New(fixMe, todo, optimize bool, custom string) *Parser {
	if (fixMe || todo || optimize) == false && custom == "" {
		fixMe = true
		todo = true
		optimize = true
	}
	return &Parser{fixMe, todo, optimize, custom, make(notesPerFile)}
}

type notesPerFile map[string][]string

// Parser - parses files and collects information about annotations
type Parser struct {
	fixMe, todo, optimize bool
	custom                string
	notesPerFile
}

var fset *token.FileSet

func init() {
	fset = token.NewFileSet() // positions are relative to fset
}

// Parse - file by file path
func (p *Parser) Parse(file string) error {
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	if _, exists := p.notesPerFile[file]; !exists {
		p.notesPerFile[file] = []string{}
	}
	for _, comment := range f.Comments {
		if p.isNote(comment.Text()) {
			p.notesPerFile[file] = append(p.notesPerFile[file], comment.Text())
		}
	}
	return nil
}
func (p *Parser) isNote(node string) bool {
	return p.todo && strings.HasPrefix(node, "TODO:") ||
		p.fixMe && strings.HasPrefix(node, "FIXME:") ||
		p.optimize && strings.HasPrefix(node, "OPTIMIZE:") ||
		p.custom != "" && strings.HasPrefix(node, fmt.Sprintf("%s:", p.custom))
}

// Aggregate - parsed files information
// Returns preformated string
func (p *Parser) Aggregate() string {
	var result string
	buf := bytes.NewBufferString(result)

	for file, notes := range p.notesPerFile {
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
