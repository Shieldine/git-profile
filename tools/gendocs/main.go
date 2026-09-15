// Command gendocs regenerates the Markdown command reference in docs/
// from the current cobra command tree. Run it after changing any
// command's flags, Short/Long text, or examples:
//
//	go run ./tools/gendocs
package main

import (
	"log"
	"os"

	"github.com/Shieldine/git-profile/cmd"
	"github.com/spf13/cobra/doc"
)

const outDir = "docs"

func main() {
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create %s: %v", outDir, err)
	}

	root := cmd.RootCmd()
	root.DisableAutoGenTag = true

	if err := doc.GenMarkdownTree(root, outDir); err != nil {
		log.Fatalf("failed to generate docs: %v", err)
	}

	log.Printf("generated command reference in %s/", outDir)
}
