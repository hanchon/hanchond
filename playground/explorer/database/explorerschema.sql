CREATE TABLE IF NOT EXISTS blocks(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    height BIGINT UNIQUE NOT NULL,
    txcount INTEGER NOT NULL,
    hash TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS blocksindex on blocks (height);

CREATE TABLE IF NOT EXISTS transactions(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    cosmoshash TEXT NOT NULL,
    ethhash TEXT NOT NULL,
    typeurl TEXT NOT NULL,
    sender TEXT NOT NULL,
    blockheight BIGINT NOT NULL REFERENCES blocks(height) ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS cosmoshash on transactions (cosmoshash);
CREATE INDEX IF NOT EXISTS ethhash on transactions (ethhash);
