package runner

import (
    "flag"
    "fmt"
    "strconv"
    "strings"
    "time"
)

type Options struct {
    TargetList   string
    TargetInput  []string
    Ports        []int

    StatusCode   bool
    ContentLength bool
    Title        bool

    Timeout      time.Duration
}

func ParseOptions() *Options {
    options := &Options{}
    flag.StringVar(&options.TargetList, "l", "", "Input file containing list of hosts to process")
    flag.Var((*StringSlice)(&options.TargetInput), "u", "Input target host(s) to probe")

    ports := flag.String("p", "", "Ports to probe (comma separated, e.g. 80,443,8080)")

    flag.BoolVar(&options.StatusCode, "sc", false, "Display response status-code")
    flag.BoolVar(&options.ContentLength, "cl", false, "Display response content-length")
    flag.BoolVar(&options.Title, "title", false, "Display page title")

    timeout := flag.Int("timeout", 10, "Timeout in seconds")

    flag.Parse()
    options.Timeout = time.Duration(*timeout) * time.Second

    if *ports != "" {
        for _, ps := range strings.Split(*ports, ",") {
            p, err := strconv.Atoi(strings.TrimSpace(ps))
            if err == nil {
                options.Ports = append(options.Ports, p)
            } else {
                fmt.Printf("Invalid port: %s\n", ps)
            }
        }
    }

    return options
}

type StringSlice []string

func (s *StringSlice) String() string {
    return fmt.Sprintf("%v", *s)
}

func (s *StringSlice) Set(value string) error {
    *s = append(*s, value)
    return nil
}
