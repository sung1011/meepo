package cmd

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"github.com/sung1011/meepo/service"
)

// SetupHandler _
type SetupHandler interface {
	// setup前 安装依赖
	BeginSetup() error
	// setup
	DoSetup() error
	// setup后 启动, 环境变量
	EndSetup() error
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

			err := h.BeginSetup()
			if err != nil {
				log.Panic(err)
			}

			err = h.DoSetup()
			if err != nil {
				log.Panic(err)
			}

			err = h.EndSetup()
			if err != nil {
				log.Panic(err)
			}
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
