package parser_test

import (
	"github.com/antlr4-go/antlr/v4"
	"os"
	"path"
	"strings"
	"testing"

	dmparser "github.com/413008319/dm-parser"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors int
}

func NewCustomErrorListener() *CustomErrorListener {
	return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func TestDmparser(t *testing.T) {
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	examplesSQLScript, err := os.ReadDir("examples-sql-script")
	require.NoError(t, err)

	files := append(examples, examplesSQLScript...)

	for i, file := range files {
		if strings.HasSuffix(file.Name(), ".sql.tree") {
			continue
		}
		var filePath string
		if i < len(examples) {
			filePath = path.Join("examples", file.Name())
		} else {
			filePath = path.Join("examples-sql-script", file.Name())
		}
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			input, err := antlr.NewFileStream(filePath)
			require.NoError(t, err)

			lexer := dmparser.NewDmSqlLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := dmparser.NewDmSqlParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = true
			_ = p.Dmprogram()

			require.Equal(t, 0, lexerErrors.errors)
			require.Equal(t, 0, parserErrors.errors)
		})
	}
}
func TestDmparser2(t *testing.T) {
	statement := `select * from aaa;`
	lexer := dmparser.NewDmSqlLexer(antlr.NewInputStream(statement))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := dmparser.NewDmSqlParser(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true
	tree := p.Dmprogram()

	for _, item := range tree.GetChildren() {
		if stmt, ok := item.(dmparser.ISql_clausesContext); ok {
			lastToken := stream.Get(stmt.GetStop().GetTokenIndex())
			text := stream.GetTextFromTokens(stmt.GetStart(), lastToken)
			print(text)
		}

	}
}
