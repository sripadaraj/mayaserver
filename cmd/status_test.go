Package cmd 
import (
 "bytes"
"strings"
"os/signal"
"github.com/mitchellh/cli"
"github.com/openebs/mayaserver/lib/config"
"github.com/openebs/mayaserver/lib/server"

)


type StatusCommand struct {

Revision		string
Status			string
Ui 			cli.Ui

}
// a cli implementation that prints the status 
// the status of server is tried to print here

func (c *StatusCommand) Help() string {
return ""
}

func (c *StatusCommand) Run(_ []string) int {
var statusString bytes.Buffer

fmt.Fprintf(&statusString, "Mayaserver s%s", c.Status)

c.Ui.Output(versionString.String())
return 0
}

//this function prints the status
func (c *StatusCommand) Synopsis() string {

return "Prints the Mayaserver status"
}
//this is to print the ip address using the hostloop

func (c *StatusCommand) Run(_ []string) int {

name, err := os.Hostname()
//checking eith any errors using the error handeling exception 
if err != nil {
     fmt.Printf("Oops: %v\n", err)
     return
}

addrs, err := net.LookupHost(name)
if err != nil {
    fmt.Printf("Oops: %v\n", err)
    return
}
//if no error the adrress is printed
for _, a := range addrs {
    fmt.Println(a)
} 
