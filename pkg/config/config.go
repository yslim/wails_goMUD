package config

import (
    "flag"
    "fmt"
    "os"
    "regexp"
    "strings"

    "github.com/magiconair/properties"
    util "github.com/yslim/go-util"
    "goMUD/pkg/log"
)

type profile struct {
    Active string `properties:"active"`
}

type tSocket struct {
    BufferSizeStr string `properties:"buffer-size,default=1MB"`
    BufferSize    int    `properties:"-"`
}

type tMud struct {
    Server string `properties:"server,default=toox.co.kr"`
    Port   int    `properties:"port,default=5800"`
}

type defaults struct {
    UsingDB     bool   `properties:"using-db,default=false"`
    Application string `properties:"application"`
}

type flags struct {
    CfgFile string
}

type config struct {
    Profile  profile    `properties:"profile"`
    Log      log.Config `properties:"log"`
    Socket   tSocket    `properties:"socket"`
    Mud      tMud       `properties:"mud"`
    Defaults defaults   `properties:"defaults"`
    Flags    flags      `properties:"-"`
}

var (
    cfg      = config{}
    Profile  = &cfg.Profile
    Log      = &cfg.Log
    Socket   = &cfg.Socket
    Mud      = &cfg.Mud
    Defaults = &cfg.Defaults
    Flags    = &cfg.Flags
    IsDev    = false
)

func init() {
    if !cfg.parse() {
        showUsage()
        os.Exit(-1)
    }
}

func showUsage() {
    fmt.Printf("Usage:%s {params}\n", os.Args[0])
    fmt.Println("      -c {config file}")
    fmt.Println("      -h (show help info)")
}

// parse command line option and load config files
func (c *config) parse() bool {
    flag.StringVar(&c.Flags.CfgFile, "c", "${APP_HOME}/etc/goMUD/application.properties", "config file")
    help := flag.Bool("h", false, "help info")

    flag.Parse()

    if !c.load() {
        return false
    }

    if *help {
        showUsage()
        return false
    }

    return true
}

// load config files and unmarshal
func (c *config) load() bool {
    c.Flags.CfgFile = replaceShellVariables(c.Flags.CfgFile)

    // check profile : dev or prod
    p := properties.MustLoadFile(c.Flags.CfgFile, properties.UTF8)
    activeProfile := p.MustGetString("profile.active")

    // read default & profile specific
    p = properties.MustLoadFiles([]string{
        c.Flags.CfgFile,
        getProfileConfigPath(c.Flags.CfgFile, activeProfile),
    }, properties.UTF8, true)

    p.MustFlag(flag.CommandLine)

    err := p.Decode(c)
    if err != nil {
        fmt.Printf("Error: config file %s loading failed. %v\n", c.Flags.CfgFile, err)
        return false
    }

    fmt.Printf("config %s load ok!\n", c.Flags.CfgFile)

    if strings.EqualFold(c.Profile.Active, "dev") {
        IsDev = true
    } else {
        IsDev = false
    }

    // size string to int value
    bs, err := util.UnmarshalText([]byte(c.Socket.BufferSizeStr))
    if err != nil {
        fmt.Printf("[ Config ] socket.buffer-size, error=%v\n", err.Error())
        os.Exit(1)
    }
    c.Socket.BufferSize = bs

    return true
}

// replace ${APP_HOME} to real path
func replaceShellVariables(str string) string {
    regex := regexp.MustCompile(`\${(.*?)}`)
    vars := regex.FindAllStringSubmatch(str, -1)

    for i := 0; i < len(vars); i++ {
        shellVar := vars[i][1]
        shellValue := os.Getenv(shellVar)
        str = strings.ReplaceAll(str, vars[i][0], shellValue)
    }

    return str
}

// ${APP_HOME}/etc/app/app-{profile}.{toml,properties,yml}
func getProfileConfigPath(cfgPath, profile string) string {
    // cfg path 에서 마지막 "/"의 위치를 찾는다. 없으면 string 의 시작점
    lastSlashIdx := strings.LastIndex(cfgPath, "/")
    if lastSlashIdx == -1 {
        lastSlashIdx = 0
    } else {
        lastSlashIdx = lastSlashIdx + 1
    }

    // .toml, .yml 등의 "." 의 index 를 찾는다
    dotIdx := strings.Index(cfgPath[lastSlashIdx:], ".") + lastSlashIdx

    cfgName := cfgPath[lastSlashIdx:dotIdx]

    return cfgPath[:lastSlashIdx] + cfgName + "-" + profile + cfgPath[dotIdx:]
}
