package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/sung1011/meepo/service"
)

// SetupHandler _
type SetupHandler interface {
	// setup前 安装依赖
	BeforeSetup()
	// setup
	DoSetup()
	// setup后 启动, 环境变量
	AfterSetup()
}

// SetupHandlers 接口组
type SetupHandlers []SetupHandler

// Core 核心引擎 状态维护
type Core struct {
	Hs SetupHandlers
}

// NewCore 初始化
func NewCore() *Core {
	return &Core{
		Hs: []SetupHandler{
			service.NewNginx(),
			// docker
			// oh-my-zsh
		},
	}
}

var wg sync.WaitGroup

// Fire _
func (c *Core) Fire() {
	wg.Add(len(c.Hs))
	for _, h := range c.Hs {
		go func(h SetupHandler) {
			defer wg.Done()
			h.BeforeSetup()
			h.DoSetup()
			h.AfterSetup()
		}(h)
	}
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup all software",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		core := NewCore()
		core.Fire()
		wg.Wait()
		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
