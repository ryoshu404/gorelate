package store

// schema holds the initial database schema.
// Applied on first run if tables do not exist.
const schema = `
CREATE TABLE IF NOT EXISTS ioc_records (
    id              SERIAL PRIMARY KEY,
    indicator       TEXT NOT NULL UNIQUE,
    type            TEXT NOT NULL,
    sources         JSONB NOT NULL DEFAULT '[]',
    first_seen      TIMESTAMPTZ NOT NULL,
    last_seen       TIMESTAMPTZ NOT NULL,
    confidence      TEXT NOT NULL DEFAULT 'low',
    tags            JSONB NOT NULL DEFAULT '[]',
    threat_category TEXT,
    raw_feed_data   JSONB,
    vt_enriched     BOOLEAN NOT NULL DEFAULT FALSE,
    detection_stubs BOOLEAN NOT NULL DEFAULT FALSE,
    active          BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS cycle_summaries (
    id           SERIAL PRIMARY KEY,
    generated_at TIMESTAMPTZ NOT NULL,
    summary_type TEXT NOT NULL,
    content      TEXT NOT NULL,
    window_start TIMESTAMPTZ NOT NULL,
    window_end   TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS detection_stubs (
    id            SERIAL PRIMARY KEY,
    ioc_record_id INTEGER NOT NULL REFERENCES ioc_records(id),
    stub_type     TEXT NOT NULL DEFAULT 'yara',
    content       TEXT NOT NULL,
    generated_at  TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_ioc_confidence ON ioc_records(confidence);
CREATE INDEX IF NOT EXISTS idx_ioc_type       ON ioc_records(type);
CREATE INDEX IF NOT EXISTS idx_ioc_last_seen  ON ioc_records(last_seen);
CREATE INDEX IF NOT EXISTS idx_ioc_active     ON ioc_records(active);
`
