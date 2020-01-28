package example

import (
	"fmt"
	"github.com/spf13/pflag"
)

type Config struct {
	Demo string `json:"demo"`
}

func NewExampleConfig()*Config{

	return &Config{}
}


func (c *Config)AddFlags(fs *pflag.FlagSet){

	fs.StringVar(&c.Demo,"demo",c.Demo,"this is a demo config")

}

func (c *Config)RunExample()error {
	fmt.Println("demo is",c.Demo)
	return nil
}
