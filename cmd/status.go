package cmd

import (
	

	"flag"
	"fmt"
  
	"sort"
	"strconv"
  "strings"
  "os"
	"github.com/openebs/mayaserver/lib/config"

	"github.com/mitchellh/cli"
)
  
  type Ports struct {
	HTTP int `mapstructure:"http"`
}

type StatusCommand struct {
	Meta
	Ui   cli.Ui
	args []string
  BindAddr   string
	Region     string
	DataDir    string
	Datacenter string
	LogLevel   string
  }
  

func usage() {
    fmt.Fprintf(os.Stderr, "usage: test-local-port [5656\n")
    flag.PrintDefaults()
    os.Exit(2)
}


func (f *ports)  readMayaConfig() *config.MayaConfig {
	var configPath []string
	cmdConfig := &config.MayaConfig{
		Ports: &config.Ports{},

 flag.Usage = usage
  flag.Parse()

  args := flag.Args()
}

func (f *ports) check() *Statuscommand {

if len(args) < 1 {
      fmt.Fprintf(os.Stderr, "Input port is missing.")
      os.Exit(1)
}

if  port := args[0]
  _, err := strconv.ParseUint(port, 10, 16)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Invalid port %q: %s\n", port, err)
    os.Exit(1)
  {
      func (mc *MayaConfig) Listener(proto, addr string, port int) (net.Listener, error) {
if 0 > port || port > 65535 {
		return nil, &net.OpError{
			Op:  "listen",
			Net: proto,
			Err: &net.AddrError{Err: "invalid port", Addr: fmt.Sprint(port)},
	  	  }
	    }
	return net.Listen(proto, net.JoinHostPort(addr, strconv.Itoa(port)))
    }
  }
}


func (mc *MayaConfig) Merge(b *MayaConfig) *MayaConfig {
if result.Ports == nil && b.Ports != nil {
		ports := *b.Ports
		result.Ports = &ports
	} else if b.Ports != nil {
		result.Ports = result.Ports.Merge(b.Ports)
}


func (a *Ports) Merge(b *Ports) *Ports {
	result := *a

	if b.HTTP != 0 {
		result.HTTP = b.HTTP
	}
	return &result
}

func (c *StatusCommand) readMayaConfig() *config.MayaConfig {

 {
 
	flags := flag.NewFlagSet("up", flag.ContinueOnError)
	flags.Usage = func() { c.Ui.Error(c.Help()) }


	flags.Var((*flaghelper.StringFlag)(&configPath), "config", "config")
  flags.StringVar(&cmdConfig.BindAddr, "bind", "", "")
	flags.StringVar(&cmdConfig.NodeName, "name", "m-apiserver", "")

	flags.StringVar(&cmdConfig.DataDir, "data-dir", "", "")
	flags.StringVar(&cmdConfig.LogLevel, "log-level", "", "")
  
  }
  
  func (c *StatusCommand) Run(args []string) int {
   
  if config.vmbox_status == 'running'
 {
 up := {func() (cli.Command, error) {
			return &cmd.UpCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                meta.Ui,
				ShutdownCh:        make(chan struct{}),
}, nil
}

if up != nil {

  mconfig := c.readMayaConfig() 

   expec_port= check () *Statuscommand 
    if expec_port=='5656'
    {
     c.Ui.Output(out)
     out :=     fmt.Sprintf("Name          IP            Ports       Status    \n%-16s%-16s%d\t%-16s",
			mconfig.NodeName,
			mconfig.BindAddr,
			mconfig.Ports,
			Status)
    }
    else fmt.Fprintf("you are running on different port  : check with your configuration") 
    
    }
    
  }
  
