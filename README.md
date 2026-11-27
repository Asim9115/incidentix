# Incidentix

A lightweight Go agent that runs on your Linux server, watches CPU, memory, and disk in real time, and automatically investigates production incidents using AI — without ever touching your code or modifying anything on the server.

---

## The problem it solves

When something goes wrong at 2am — memory climbing, disk full, a service crash-looping — someone has to SSH in, dig through logs, correlate it with a recent deploy, and guess at a fix. For small teams there's no dedicated SRE. That's what Incidentix is for.

---

## How it works

1. A small Go agent runs on your server as a systemd service
2. It samples CPU, memory, and disk on a short tick and keeps a rolling window
3. If a rule holds for the whole window (memory above 90%, CPU saturated, disk full), it opens an incident once
4. The control plane stores the incident, then asks the agent for specific evidence
5. The agent runs a read-only tool, sends the result back, and the investigator writes a root-cause report with a suggested fix
6. You review the report, open a PR, and ship the fix through your normal workflow

The agent never modifies production. Everything is read-only. Fixes are suggestions.

The agent only makes outbound calls. The control plane never connects in. Each step is a short request. The investigation state lives in the database, not on an open HTTP connection.

---

## Requests

Register a server and put the token in the agent environment. If the token is missing, the agent exits. It does not prompt.

```bash
POST /v1/agents/server
{"name": "api-1"}

→ {"server_id": "...", "token": "idx-..."}
```

The raw token is returned once. The API stores a hash. Later requests send `Authorization: Bearer <token>`. The server id comes from that token.

Heartbeats:

```bash
POST /v1/agents/metrics
Authorization: Bearer <token>

{"cpu": {"usage_percent": 91.2}, "memory": {...}, "disk": {...}, "timestamp": "..."}
```

When a rule fires, the agent opens an incident with the metric window. The handler inserts the row as `investigating` and returns immediately. It does not wait for the model.

```bash
POST /v1/incidents
Authorization: Bearer <token>

{"type": "CPUHigh", "severity": "High", "metrics": [/* window */]}

→ {"incident_id": "...", "status": "investigating"}
```

While the rule stays true, the agent does not open another incident. When the metric drops, it marks the incident resolved.

A worker then picks the incident and asks the model what evidence it needs. The model can only choose from a fixed tool list, for example `tail_logs(source, lines)` or `top_processes(limit)`. The worker stores that as a command with status `pending`.

The agent polls for work and posts the tool output:

```bash
GET /v1/commands
Authorization: Bearer <token>

→ {"id": "...", "incident_id": "...", "tool": "tail_logs", "args": {"source": "journald", "lines": 10}}

POST /v1/commands/:id/result
Authorization: Bearer <token>

{"output": "..."}
```

The worker passes the output back to the model. The model either queues one more tool or writes the report on the incident and sets status to `ready`. The report is the root cause and the suggested change. The agent does not apply it.

The agent refuses any tool outside the list, and it refuses anything that writes.

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
    transport/    outbound client: metrics, incidents, command results
    application/  ties everything together, owns the run loop
```

---

## Running locally (development)

Start the API:
```bash
go run ./cmd/api/main.go
```

Start the agent:
```bash
go run ./cmd/agent/main.go
```

The API listens on `:8080`. The agent POSTs metrics to `http://localhost:8080`. The bearer token in the request flow above is not wired up yet.

---

## Status

Early development. The agent collects metrics and detects anomalies locally. The API stores servers and metrics. Token auth, incidents, and the command loop are next.

- [x] Collector — CPU, memory, disk
- [x] Detector — memory pressure, CPU saturation, disk full rules
- [x] Agent loop — rolling window
- [x] Transport — POST metrics to the API
- [x] API with PostgreSQL — servers and metrics
- [ ] Bearer token on agent requests
- [ ] POST incident, poll commands, POST tool results
- [ ] Investigator — tool choice and report
- [ ] React dashboard
- [ ] Install script
