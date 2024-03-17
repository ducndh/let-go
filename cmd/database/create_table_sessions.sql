CREATE TABLE sessions (
token CHAR(43) PRIMARY KEY,
data BYTEA NOT NULL,
expiry TIMESTAMP NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);