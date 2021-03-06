package app

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/omecodes/common/env/web/app"
	templates2 "github.com/omecodes/common/env/web/templates"
	"github.com/omecodes/common/futils"
	"github.com/omecodes/common/utils/jcon"
	"github.com/omecodes/common/utils/lang"
	log2 "github.com/omecodes/common/utils/log"
	"github.com/shibukawa/configdir"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type App struct {
	sync.Mutex
	vendor string
	name   string

	options     *options
	initialized bool

	cmd          *cobra.Command
	configureCMD *cobra.Command
	startCMD     *cobra.Command
	versionCMD   *cobra.Command

	translationsDir string
	dataDir         string
	cacheDir        string
	wwwDir          string
	webAppsDir      string
	templatesDir    string
	Resources       *Resources
	configs         jcon.Map
}

func (a *App) init() {
	if a.initialized {
		return
	}
	a.initialized = true

	a.Lock()
	defer a.Unlock()

	execName := a.name

	a.cmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: fmt.Sprintf("Run %s help command", execName),
		Run: func(cmd *cobra.Command, args []string) {
			if a.options.startCMDFunc != nil {
				err := a.initDirs()
				if err != nil {
					log.Fatalln(err)
				}

				log2.File = filepath.Join(a.dataDir, "configure.log")

				err = a.initResources()
				if err != nil {
					log2.Fatal("resources init", log2.Err(err))
				}

				if len(a.options.configItems) > 0 {
					cfgFilename := filepath.Join(a.dataDir, "configs.json")
					err = jcon.Load(cfgFilename, &a.configs)
					if err != nil {
						log2.Error("configs loading", log2.Err(err))
					}
				}

				a.options.startCMDFunc()
			} else {
				if err := cmd.Help(); err != nil {
					log2.Fatal("cmd", log2.Err(err))
				}
			}
		},
	}

	flags := a.cmd.PersistentFlags()
	flags.StringVar(&a.wwwDir, "www", "", "Web resources dir")
	flags.StringVar(&a.templatesDir, "tmpl", "", "Templates resources dir")

	// add configure command
	if len(a.options.configItems) > 0 {
		a.configureCMD = &cobra.Command{
			Use:   "configure",
			Short: fmt.Sprintf("Configure %s", execName),
			Run: func(cmd *cobra.Command, args []string) {
				err := a.initDirs()
				if err != nil {
					log.Fatalln(err)
				}

				log2.File = filepath.Join(a.dataDir, "configure.log")

				configFilename := filepath.Join(a.dataDir, "configs.json")
				oldConf := jcon.Map{}
				_ = jcon.Load(configFilename, &oldConf)

				err = a.configure(configFilename, os.ModePerm, a.options.configItems...)
				if err != nil {
					log2.Fatal("configure failed", log2.Err(err))
				}

				if a.options.afterConfigure != nil {
					err = a.options.afterConfigure(a.configs)
					if err != nil {
						log2.Fatal("post configure failed", log2.Err(err))
					}

					err = a.configs.Save(configFilename, os.ModePerm)
					if err != nil {
						log2.Fatal("save configs file", log2.Err(err))
					}
				}
			},
		}
		a.cmd.AddCommand(a.configureCMD)
	}

	// add run command
	if a.options.startCMDFunc != nil {
		a.startCMD = &cobra.Command{
			Use:   "start",
			Short: fmt.Sprintf("Start %s", execName),
			Run: func(cmd *cobra.Command, args []string) {
				err := a.initDirs()
				if err != nil {
					log.Fatalln(err)
				}
				log2.File = filepath.Join(a.dataDir, "run.log")

				err = a.initResources()
				if err != nil {
					log2.Fatal("resources init", log2.Err(err))
				}

				cfgFilename := filepath.Join(a.dataDir, "configs.json")
				if futils.FileExists(cfgFilename) {
					err = jcon.Load(cfgFilename, &a.configs)
					if err != nil {
						log2.Error("configs loading", log2.Err(err))
					}
				}

				a.options.startCMDFunc()
			},
		}
		a.cmd.AddCommand(a.startCMD)
	}

	// add version command
	if a.options.version != "" {
		a.versionCMD = &cobra.Command{
			Use:   "version",
			Short: "Displays app name and version",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(a.options.version)
			},
		}
		a.cmd.AddCommand(a.versionCMD)
	}
}

func (a *App) initDirs() error {
	// initializing directories
	dirs := configdir.New(a.vendor, a.name)
	globalFolder := dirs.QueryFolders(configdir.Global)[0]
	cacheFolder := dirs.QueryFolders(configdir.Cache)[0]

	a.dataDir = globalFolder.Path

	if a.options.version != "" {
		a.dataDir = filepath.Join(a.dataDir, fmt.Sprintf("v%s", a.options.version))
	}

	err := os.MkdirAll(a.dataDir, os.ModePerm)
	if err != nil {
		return err
	}

	a.cacheDir = cacheFolder.Path
	err = os.MkdirAll(a.cacheDir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initResources() error {
	// initializing resources manager is required
	if a.options.withResources {
		a.Resources = new(Resources)

		if a.templatesDir == "" {
			a.templatesDir = filepath.Join(a.dataDir, "res", "templates")
			err := os.MkdirAll(a.templatesDir, os.ModePerm)
			if err != nil {
				return err
			}
		}

		i18nDir := filepath.Join(a.dataDir, "res", "i18n")
		err := os.MkdirAll(i18nDir, os.ModePerm)
		if err != nil {
			return err
		}
		a.Resources.templates = templates2.NewFolder(a.templatesDir)

		a.Resources.i18n = lang.NewManager(i18nDir)
		err = a.Resources.i18n.Load()
		if err != nil {
			return err
		}

		if a.wwwDir == "" {
			a.wwwDir = filepath.Join(a.dataDir, "res", "www")
			err := os.MkdirAll(a.wwwDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		a.Resources.staticsDir = a.wwwDir

		if a.webAppsDir == "" {
			a.webAppsDir = filepath.Join(a.dataDir, "res", "webapps")
			err := os.MkdirAll(a.webAppsDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		a.Resources.web = app.NewFolder(a.webAppsDir, a.Resources.i18n)
	}
	return nil
}

func (a *App) configure(outputFilename string, mode os.FileMode, items ...configItem) error {
	oldValues := jcon.Map{}
	_ = jcon.Load(outputFilename, &oldValues)

	a.configs = jcon.Map{}
	for _, item := range items {
		key := item.configType.String()
		itemOldValues := oldValues.GetConf(key)

		values, err := item.create(item.description, itemOldValues)
		if err != nil {
			return err
		}
		a.configs.Set(key, values)
	}

	return a.configs.Save(outputFilename, mode)
}

func (a *App) GetConfig(item ConfigType) jcon.Map {
	return a.configs.GetConf(item.String())
}

func (a *App) GetCommand() *cobra.Command {
	return a.cmd
}

func (a *App) StartCommand() *cobra.Command {
	return a.startCMD
}

func (a *App) ConfigureCommand() *cobra.Command {
	return a.configureCMD
}

func (a *App) InitDirs() error {
	return a.initDirs()
}

func (a *App) LoadConfigs() error {
	cfgFilename := filepath.Join(a.dataDir, "configs.json")
	return jcon.Load(cfgFilename, &a.configs)
}

func (a *App) InitResources() error {
	return a.initResources()
}

func (a *App) DataDir() string {
	return a.dataDir
}

func (a *App) CacheDir() string {
	return a.cacheDir
}

func (a *App) Label() string {
	return strcase.ToDelimited(a.name, ' ')
}

func (a *App) Name() string {
	return a.name
}

func New(vendor string, name string, opts ...Option) *App {
	a := &App{
		vendor:  vendor,
		name:    name,
		options: new(options),
		configs: jcon.Map{},
	}
	for _, opt := range opts {
		opt(a.options)
	}
	a.init()
	return a
}
