DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users
(
    chat_id INT         NOT NULL,
    state   VARCHAR(20) NOT NULL,
    PRIMARY KEY (chat_id),
    UNIQUE (chat_id)
);

CREATE TABLE IF NOT EXISTS services
(
    id           SERIAL       NOT NULL,
    name         VARCHAR(255) NOT NULL,
    login        VARCHAR(255) NOT NULL,
    password     VARCHAR(255) NOT NULL,
    user_chat_id INT          NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_chat_id) REFERENCES users (chat_id)
);
