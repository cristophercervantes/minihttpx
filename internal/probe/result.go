package probe

import (
    "io"
    "net/http"
    "regexp"
    "strings"
)

type Result struct {
    Target        string
    StatusCode    int
    ContentLength int64
    ContentType   string
    Title         string
}

func NewResult(target string) *Result {
    return &Result{Target: target}
}

func (r *Result) ExtractTitle(resp *http.Response) {
    body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
    if err != nil {
        return
    }
    resp.Body.Close()
    resp.Body = io.NopCloser(strings.NewReader(string(body)))

    re := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
    match := re.FindStringSubmatch(string(body))
    if len(match) > 1 {
        r.Title = strings.TrimSpace(match[1])
    }
}
