package options

import (
	"go-command/example"
	cliflag "k8s.io/component-base/cli/flag"
)


//add other config
type Options struct {

	Example   *example.Config `json:"example"`
}




func NewOptions()(*Options,error){

	return &Options{Example:example.NewExampleConfig()},nil
}



func (o *Options) Flags() (nfs cliflag.NamedFlagSets) {

	o.Example.AddFlags(nfs.FlagSet("demo config"))

	return nfs
}