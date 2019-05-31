package mod

import (
	"flag"
	"fmt"
	"github.com/beego/bee/cmd/commands"
	"github.com/beego/bee/cmd/commands/version"
	"github.com/beego/bee/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var m = `
module {{.appname}}

go 1.12

`

var mod = &commands.Command{
	CustomFlags: true,
	UsageLine:   "mod",
	Short:       "create mod go application",
	Long: ` bee mod -app application_name"
`,
	PreRun: func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() },
	Run:    modApp,
}

var appname string

func init() {
	fs := flag.NewFlagSet("mod", flag.ContinueOnError)
	fs.StringVar(&appname, "app", "", "Connection string used by the driver to connect to a database instance.")
	mod.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, mod)
}

func modApp(cmd *commands.Command, args []string) int {

	cmd.Flag.Parse(args)

	fmt.Println(args)
	app, _ := os.Getwd()
	app = filepath.Join(app, appname)

	os.MkdirAll(app, 0755)

	utils.WriteToFile(path.Join(app, "go.mod"),
		strings.Replace(m, "{{.appname}}", appname, -1))

	return 0
}
