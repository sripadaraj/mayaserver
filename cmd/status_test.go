package main

import (
   
    "fmt"
    "github.com/mitchellh/cli"
    "os"
    "time"
    "encoding/json"
 
)

type Configuration struct {
    URL      []string
    Mem string
}


func loadConfig(path string) (Configuration, error) {

    file, _ := os.Open(path)
    defer file.Close()

    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)

    return configuration, err
}


func main() {
    type Result struct {
        url    string
        status Status
    }

    rc := make(chan Result)

    configuration, err := loadConfig("config.json")

    if err != nil {
        fmt.Println("Error :", err)
        return
    }

    sites := make(map[string]*Site, len(configuration.URLs))

    for _, url := range configuration.URLs {
        sites[url] = &Site{url, UNCHECKED}
    }

    mc := memcache.New(configuration.Memcached)

    sites_output := make(map[string]bool)

    for {
        for _, site := range sites {
            go func(site *Site, rc chan Result) {
                status, _ := site.Status()
                rc <- Result{site.url, status}
            }(site, rc)
        }

        for i := 0; i < len(sites); i++ {
            res := <-rc
            site := sites[res.url]
            if site.last_status != res.status {
                sites[res.url].last_status = res.status
            }
        }

        for k, v := range sites {
            sites_output[k] = v.last_status == 2
        }

        site_json, err := json.Marshal(sites_output)

        if err != nil {
            panic(err)
        }
        mc.Set(&memcache.Item{
            Key:   "mission_control.sites",
            Value: site_json,
        })
        time.Sleep(time.Second * 5)
    }
}
