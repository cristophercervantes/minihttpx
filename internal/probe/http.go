package probe

import (
    "context"
    "net/http"
    "time"
)

type HTTPProbe struct {
    client *http.Client
}

func NewHTTPProbe(timeout time.Duration) *HTTPProbe {
    return &HTTPProbe{
        client: &http.Client{Timeout: timeout},
    }
}

func (p *HTTPProbe) Probe(ctx context.Context, target string) (*Result, error) {
    result := NewResult(target)

    req, err := http.NewRequestWithContext(ctx, "GET", target, nil)
    if err != nil {
        return result, err
    }

    resp, err := p.client.Do(req)
    if err != nil {
        return result, err
    }
    defer resp.Body.Close()

    result.StatusCode = resp.StatusCode
    result.ContentLength = resp.ContentLength
    result.ContentType = resp.Header.Get("Content-Type")
    result.ExtractTitle(resp)

    return result, nil
}
