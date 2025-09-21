package runner

import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strings"

    "github.com/cristophercervantes/minihttpx/internal/probe"
)

type Runner struct {
    options *Options
}

func New(options *Options) *Runner {
    return &Runner{options: options}
}

func (r *Runner) Run() error {
    ctx := context.Background()
    httpProbe := probe.NewHTTPProbe(r.options.Timeout)

    var targets []string
    targets = append(targets, r.options.TargetInput...)

    if r.options.TargetList != "" {
        file, err := os.Open(r.options.TargetList)
        if err != nil {
            return err
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            targets = append(targets, scanner.Text())
        }
    }

    if len(targets) == 0 {
        return fmt.Errorf("no targets provided")
    }

    for _, target := range targets {
        target = strings.TrimSpace(target)
        if target == "" {
            continue
        }
        if len(r.options.Ports) > 0 {
            for _, port := range r.options.Ports {
                finalTarget := ensureScheme(fmt.Sprintf("%s:%d", stripScheme(target), port))
                r.runProbe(ctx, httpProbe, finalTarget)
            }
        } else {
            finalTarget := ensureScheme(target)
            r.runProbe(ctx, httpProbe, finalTarget)
        }
    }
    return nil
}

func (r *Runner) runProbe(ctx context.Context, httpProbe *probe.HTTPProbe, target string) {
    result, err := httpProbe.Probe(ctx, target)
    if err != nil {
        fmt.Printf("%s | error: %v\n", target, err)
        return
    }

    out := []string{target}
    if r.options.StatusCode {
        out = append(out, fmt.Sprintf("SC=%d", result.StatusCode))
    }
    if r.options.ContentLength {
        out = append(out, fmt.Sprintf("CL=%d", result.ContentLength))
    }
    if r.options.Title && result.Title != "" {
        out = append(out, fmt.Sprintf("Title=%s", result.Title))
    }

    fmt.Println(strings.Join(out, " | "))
}

func ensureScheme(target string) string {
    if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
        return target
    }
    return "http://" + target
}

func stripScheme(target string) string {
    if strings.HasPrefix(target, "http://") {
        return strings.TrimPrefix(target, "http://")
    }
    if strings.HasPrefix(target, "https://") {
        return strings.TrimPrefix(target, "https://")
    }
    return target
}
