# Incidentix

A lightweight Go agent that runs on your Linux server, watches CPU, memory, and disk in real time, and automatically investigates production incidents using AI — without ever touching your code or modifying anything on the server.

---

## The problem it solves

When something goes wrong at 2am — memory climbing, disk full, a service crash-looping — someone has to SSH in, dig through logs, correlate it with a recent deploy, and guess at a fix. For small teams there's no dedicated SRE. That's what Incidentix is for.

---

## How it works

1. A small Go agent runs on your server as a systemd service
2. It samples CPU, memory, and disk every 5 seconds
3. If something looks wrong (e.g. memory above 90% for 25 seconds), it fires an incident
4. The evidence — metrics, logs, process info — gets bundled and sent to the control plane
5. An AI investigator reads the evidence and writes a root-cause report with a suggested fix
6. You review the report, open a PR, and ship the fix through your normal workflow

The agent never modifies production. Everything is read-only. Fixes are suggestions.

---

## Stack

| Part | Tech |
|---|---|
| Agent | Go |
| API / control plane | Go |
| Investigator | Python + FastAPI |
| Database | PostgreSQL |
| Dashboard | React |


---

## Project layout

```
cmd/
  agent/      the agent binary
  api/        the control plane API
internal/
  agent/
    collector/    reads CPU, memory, disk from /proc
    detector/     evaluates rules against the rolling window
    transport/    sends heartbeats and incidents to the API
    application/  ties everything together, owns the run loop
```

---

## Running locally (development)

Start the test server:
```bash
go run ./cmd/api/main.go
```

Start the agent:
```bash
go run ./cmd/agent/main.go
```

The agent will POST metrics to `localhost:8000` every few seconds. You'll see them printed in the server terminal.

---

## Status

Early development. The agent collects metrics and detects anomalies locally. The API, investigator, and dashboard are next.

- [x] Collector — CPU, memory, disk
- [x] Detector — memory pressure, CPU saturation, disk full rules
- [x] Agent loop — rolling window, 5s tick
- [x] Transport — POST metrics to API
- [ ] Real API with PostgreSQL
- [ ] AI investigator
- [ ] React dashboard
- [ ] Install script
