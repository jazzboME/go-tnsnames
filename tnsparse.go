package tnsparse

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"tnsparse/parser"
)

// Tnsnames holds the list of found aliases in the tnsnames file
var Tnsnames []string;

// TnsnamesListener is local tnsnames type
type tnsnamesListener struct{
	*parser.BasetnsnamesParserListener
}

// NewTnsnames constructes a new tnsnames
func newtnsnamesListener() *tnsnamesListener {
	return new(tnsnamesListener)
}

// EnterAlias overrride to parse each Alias in the tnsnames
func(s *tnsnamesListener) EnterAlias(ctx *parser.AliasContext) {
	Tnsnames = append(Tnsnames, ctx.GetText())
}

// Parse processes the passed string (which presumably is a tnsnames.ora file)
func Parse(inputStream string) {
	input := antlr.NewInputStream(inputStream)
	lexer := parser.NewtnsnamesLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewtnsnamesParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Tnsnames()
	antlr.ParseTreeWalkerDefault.Walk(newtnsnamesListener(), tree)
}
