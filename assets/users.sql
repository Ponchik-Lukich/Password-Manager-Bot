CREATE TABLE IF NOT EXISTS users
(
    chat_id    INT         NOT NULL,
    state      VARCHAR(20) NOT NULL,
    message_id integer NOT NULL DEFAULT 0,
    PRIMARY KEY (chat_id),
    UNIQUE (chat_id)
);