
CREATE TABLE users (
    id           BIGSERIAL PRIMARY KEY,
    username     TEXT    NOT NULL UNIQUE,
    password_hash TEXT   NOT NULL DEFAULT '',
    salt          TEXT   NOT NULL DEFAULT '',
    is_superuser  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at    BIGINT  NOT NULL,
    updated_at    BIGINT  NOT NULL
);
