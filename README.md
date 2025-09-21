# minihttpx

minihttpx is a **lightweight, low-resource HTTP probing tool** inspired by `httpx`.  
It’s designed for reconnaissance on low-spec machines (small VMs, laptops, Raspberry Pi), focusing on fast, simple HTTP checks: status codes, content length, and page titles.  

> ⚡ Minimal dependencies — built using Go standard library. Fast and easy to run on constrained systems.

---

## Features

- Input:
  - `-u` : specify one or more target URLs/hosts (repeatable)
  - `-l` : file containing targets (one per line)
  - `-p` / `-port` : comma-separated ports to probe (e.g., `80,443,8080`)
- Probes:
  - `-sc` : show HTTP status code
  - `-cl` : show content length
  - `-title` : extract HTML `<title>` tag
- Config:
  - `--timeout` : request timeout in seconds (default: 10s)
- Output:
  - Plain text, easy to pipe into other tools

---

## Install & Build

### Quick install using `go install`

If you have Go installed (1.20+), you can install `minihttpx` directly:

```bash
go install github.com/cristophercervantes/minihttpx/cmd/minihttpx@latest
```

This will place the `minihttpx` binary in your `$GOBIN` or `$GOPATH/bin`.

Make sure it’s in your `PATH`, for example:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

Now you can run it from anywhere:

```bash
minihttpx -u example.com -sc -title
```

### Build from source

Clone and build locally:

```bash
git clone https://github.com/cristophercervantes/minihttpx
cd minihttpx
go build -o minihttpx cmd/minihttpx/main.go
```

---

## Usage

### Basic single target
```bash
./minihttpx -u example.com -sc -title -cl
```

### Multiple targets (repeatable `-u`)
```bash
./minihttpx -u example.com -u example.org -sc
```

### Use a targets file
```bash
./minihttpx -l targets.txt -sc -title
```

### Probe specific ports
```bash
./minihttpx -u example.com -p 80,443,8080 -sc -title
```

### Set timeout
```bash
./minihttpx -u example.com --timeout 5
```

---

## Examples (bug-hunting workflow)

1. **Find live hosts quickly**
```bash
./minihttpx -l all_hosts.txt -sc -cl > live.txt
```

2. **Then feed into a scanner**
```bash
cat live.txt | awk '{print $1}' > urls.txt
nuclei -l urls.txt -t vulnerabilities/
```

3. **Parallel scanning on low-RAM machines**
Use GNU `parallel` to distribute load:
```bash
cat targets.txt | parallel -j 8 ./minihttpx -u {} -sc -title
```

---

## Contributing

Contributions are welcome. If you want features like JSON output, TLS info, or concurrency controls, open an issue or submit a PR. Keep the project lightweight if adding features — prefer optional flags that can be disabled.

---

## License

This project is provided as-is under the MIT License. See `LICENSE` for details.

---

## Safety & Legal

Only scan systems you own or where you have explicit permission. Unauthorized scanning can be illegal and unethical.
