package cmd

import (
	

	"flag"
	"fmt"
  
	"sort"
	"strconv"
	"string"
  	"os"
	"github.com/openebs/mayaserver/lib/config"
//using maya configuration can print the bind addr and nodes name
	"github.com/mitchellh/cli"
//cli imports all required cli-ui commands
)
  
  type Ports struct {
	HTTP int `mapstructure:"http"`
}
//this structure helps in managing the ports
//also supports default ports to server

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
// status structure which includes the cofig nodes

func usage() {
    fmt.Fprintf(os.Stderr, "usage: test-local-port [5656]\n")
    flag.PrintDefaults()
    os.Exit(2)
} //this function returns the actual port usage


func (f *ports)  readMayaConfig() *config.MayaConfig {
	var configPath []string
	cmdConfig := &config.MayaConfig{
		Ports: &config.Ports{},

 flag.Usage = usage
  flag.Parse()

  args := flag.Args()
}

	
//check with the port details to know if the port is used or not	
func (f *ports) check() *Statuscommand {

if len(args) < 1 {
      fmt.Fprintf(os.Stderr, "Input port is missing.")
      os.Exit(1)
}//checks the input port 

if  port := args[0]
  _, err := strconv.ParseUint(port, 10, 16)
	//this coverts the string to uint and compares with the actual expected port
  if err != nil {
    fmt.Fprintf(os.Stderr, "Invalid port %q: %s\n", port, err)
    os.Exit(1)
    }
  }
}

//using merge function in configuration to merge actual port and the expected port
func (mc *MayaConfig) Merge(b *MayaConfig) *MayaConfig {
if result.Ports == nil && b.Ports != nil {
		ports := *b.Ports
		result.Ports = &ports
	} else if b.Ports != nil {
		result.Ports = result.Ports.Merge(b.Ports)
}

//merge function 
func (a *Ports) Merge(b *Ports) *Ports {
	result := *a

	if b.HTTP != 0 {
		result.HTTP = b.HTTP
	}
	return &result
}

	//prints the cofiguration of status 
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
  
	//the main status command
	//here when the vagrant machine is running
	//then the status runs
  func (c *StatusCommand) Run(args []string) int {
   //check for vagrant up
	  //because when vagrant up is not done then the m-apiserver could not be responded
  if config.vmbox_status == 'running'
 {
	 //assiging the up variable the value of up command to check
 up := {func() (cli.Command, error) {
			return &cmd.UpCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                meta.Ui,
				ShutdownCh:        make(chan struct{}),
}, nil
}
//CMD UP is mot nil then the status command can be executed 
	//up variable either stores the nil | up terms 
	//when it is not nil the status command can be executed
if up != nil {

  mconfig := c.readMayaConfig() 

   expec_port= check () *Statuscommand 
	//the merged ports can be verfied when merged the ports may be sam as expected port and the port running
	//by this check the code can be runned 
	//when it is not running on the actual/expected port he status command cannot be executed
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
  
