# gorelate

> **Note: This README is not final. The project is actively under development.**

A threat intelligence pipeline written in Go. Gorelate ingests public IOC feeds, normalizes and deduplicates indicators, correlates across sources for confidence scoring, generates YARA detection stubs from high-confidence IOCs, and provides LLM-powered triage summaries of emerging threat patterns.

~~Live at **[iocs.ryoshu404.com](https://iocs.ryoshu404.com)**~~ Upon completion

---

## What It Does

- Ingests IOC feeds from AlienVault OTX, URLHaus, and abuse.ch on an hourly schedule
- Normalizes indicators into a common schema and deduplicates by indicator value
- Scores confidence based on source correlation: 1 source → low, 2 → medium, 3 → high
- Enriches high-confidence IOCs with VirusTotal data (free tier)
- Generates YARA detection stubs for high-confidence indicators
- Produces LLM-powered analyst summaries on an 8-hour and daily schedule via Anthropic API
- Exposes a fully read-only REST API, publicly accessible with no authentication

---

## Architecture

```
Schedule trigger (Go ticker — hourly)
→ Run all fetchers (OTX, URLHaus, abuse.ch)
→ Normalize and dedup
→ Correlate and score
→ VT enrichment if HIGH confidence
→ Check if 8-hour or daily summary is due
    → Yes: generate cycle summary, cache to DB
    → No: skip, continue
→ Check if HIGH confidence IOCs crossed stub threshold
    → Yes: generate YARA stubs
→ Sleep until next cycle

API layer (independent, read-only from DB)
```

---

## API

Read-only. No authentication. No write endpoints.

```
GET /ioc/{indicator}                → clean summary response
GET /ioc/{indicator}/full           → complete record including raw feed data
GET /ioc/{indicator}/sources        → sources array and per-source metadata
GET /ioc/{indicator}/confidence     → confidence tier
GET /iocs?q=&type=&confidence=&tag= → filtered bulk query
GET /summaries                      → paginated list of cycle summaries
GET /summaries/latest               → most recent per type (emerging, daily)
GET /summaries/{id}                 → specific summary by ID
GET /health                         → service health check
```

**Example response (`GET /ioc/{indicator}`):**
```json
{
  "indicator": "evil.com",
  "type": "domain",
  "confidence": "high",
  "sources": ["otx", "abuse.ch", "urlhaus"],
  "tags": ["phishing", "c2"],
  "threat_category": "phishing",
  "first_seen": "2026-03-01T00:00:00Z",
  "last_seen": "2026-03-28T08:00:00Z",
  "vt_enriched": true,
  "detection_stubs": true,
  "llm_summary_available": false
}
```

---

## Project Structure

```
gorelate/
├── cmd/gorelate/         # entrypoint
├── internal/
│   ├── fetchers/         # one fetcher per feed (OTX, URLHaus, abuse.ch)
│   ├── pipeline/         # orchestrator — wires fetchers through the full cycle
│   ├── correlator/       # dedup and source merge across feeds
│   ├── scorer/           # confidence tier assignment
│   ├── enrichment/       # VirusTotal enrichment (HIGH confidence only)
│   ├── scheduler/        # ticker + wall clock summary triggers
│   ├── store/            # PostgreSQL — models, interface, schema
│   ├── api/              # read-only HTTP API
│   ├── stubs/            # YARA stub generation
│   └── summarizer/       # LLM triage summary via Anthropic API
├── frontend/             # single-file HTML/CSS/JS dashboard
├── deploy/               # Cloudflare Tunnel config example
├── scripts/              # retention cleanup
├── Dockerfile
├── docker-compose.yml
└── .env.example
```

---

## Data Sources

| Feed | Role |
|---|---|
| AlienVault OTX | Core — confidence scoring |
| URLHaus | Core — confidence scoring |
| abuse.ch | Core — confidence scoring |
| VirusTotal | Enrichment only — HIGH confidence IOCs |

URLHaus and abuse.ch share an operator and will sometimes report the same indicators — this validates correlation rather than inflating it.

---

## Confidence Scoring

| Sources reporting | Confidence |
|---|---|
| 1 | low |
| 2 | medium |
| 3 | high |

VirusTotal enriches the record at HIGH confidence but does not affect the tier. HIGH is the ceiling.

---

## Scheduler

```
Ingestion cycle  → every 1 hour
8-hour summary   → 06:00, 14:00, 22:00 UTC
Daily summary    → 12:00 UTC
```

If the daily summary fired within the last hour, the 8-hour call is skipped. If no HIGH confidence IOCs exist in the window, the LLM call is skipped and the cycle is logged.

---

## Running Locally

```bash
cp .env.example .env
# fill in API keys

docker compose up --build
```

API available at `http://localhost:8080`.

---

## Deployment

Dockerized, running on a home server, exposed via Cloudflare Tunnel. See `deploy/cloudflare-tunnel.yml.example` for tunnel config reference.

---

## License

Commons Clause + Apache 2.0. Free to use and modify. Not for commercial resale.

---

## Related Projects

- [statica](https://github.com/ryoshu404/statica) — static analysis pipeline
- [macollect](https://github.com/ryoshu404/macollect) — macOS DFIR forensic artifact collector
