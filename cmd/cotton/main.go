package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/cotton-go/cotton/internal/generate"
)

const usage = `USAGE
  weaver generate                 // weaver code generator
  weaver version                  // show weaver version
`

func main() {
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "generate":
		generateFlags := flag.NewFlagSet("generate", flag.ExitOnError)
		tags := generateFlags.String("tags", "", "Optional tags for the generate command")
		generateFlags.Usage = func() {
			fmt.Fprintln(os.Stderr, generate.Usage)
		}
		generateFlags.Parse(flag.Args()[1:])
		buildTags := "ignoreWeaverGen"
		if *tags != "" { // tags flag was specified
			// TODO(rgrandl): we assume that the user specify the tags properly. I.e.,
			// a single tag, or a list of tags separated by comma. We may want to do
			// extra validation at some point.
			buildTags = buildTags + "," + *tags
		}
		if err := generate.Generate(".", generateFlags.Args(), generate.Options{BuildTags: buildTags}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return

	case "version":
	default:
	}
}
