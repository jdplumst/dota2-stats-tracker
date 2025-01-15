-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE session (
    id TEXT PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE hero (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    img TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE faction (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE position (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE match (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_id INTEGER NOT NULL,
    hero_id INTEGER NOT NULL,
    match_id INTEGER NOT NULL,
    date DATE NOT NULL,
    faction_id INTEGER NOT NULL,
    position_id INTEGER NOT NULL,
    result BOOLEAN NOT NULL,
    lane_result BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (hero_id) REFERENCES hero(id),
    FOREIGN KEY (faction_id) REFERENCES faction(id),
    FOREIGN KEY (position_id) REFERENCES position(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE session;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE hero;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE match;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE faction;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE position;
-- +goose StatementEnd
