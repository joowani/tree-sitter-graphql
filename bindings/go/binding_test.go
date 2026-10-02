package tree_sitter_graphql_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_graphql "github.com/joowani/tree-sitter-graphql/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_graphql.Language())
	if language == nil {
		t.Fatal("Error loading GraphQL grammar")
	}
	parser := tree_sitter.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(language); err != nil {
		t.Errorf("Error loading GraphQL grammar: %v", err)
	}
}
