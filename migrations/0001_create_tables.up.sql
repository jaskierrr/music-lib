BEGIN;

CREATE TABLE IF NOT EXISTS songs (
    id BIGINT PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    link TEXT,
    lyrics TEXT,
    release_date VARCHAR(255),
    song_name VARCHAR(255) NOT NULL
);

COMMIT;
