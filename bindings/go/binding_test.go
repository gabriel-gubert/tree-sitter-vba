package tree_sitter_vba_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_vba "github.com/gabriel-gubert/tree-sitter-vba/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_vba.Language())
	if language == nil {
		t.Errorf("Error loading no grammar")
	}
}
