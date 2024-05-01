CREATE TABLE player
(
    id       INT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    balance  INT                NOT NULL DEFAULT 0
);

CREATE TABLE player_history
(
    id         INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    player_id  INT          NOT NULL,
    game_id    VARCHAR(255) NOT NULL,
    risk       VARCHAR(50)  NOT NULL,
    bet_amount INT          NOT NULL DEFAULT 0,
    win_amount INT          NOT NULL DEFAULT 0,
    created_at DATE         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player_id) REFERENCES player (id)
);

INSERT INTO player (id, username, balance)
VALUES (1, 'vinicius', 100000);