-- +migrate Up
CREATE TABLE `game_result` (
    `id` VARCHAR(10) NOT NULL,
    `description` VARCHAR(50),
    PRIMARY KEY (`id`)
);
CREATE TABLE `game` (
    `id` VARCHAR(50) NOT NULL,
    `site` VARCHAR(100),
    `result` VARCHAR(100),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`result`) REFERENCES `game_result` (`id`)
);
CREATE TABLE `color` (
    `name` VARCHAR(5) NOT NULL,
    PRIMARY KEY (`name`)
);
CREATE TABLE `player` (
    `id` VARCHAR(50) NOT NULL,
    `game_id` VARCHAR(50) NOT NULL,
    `user_id` VARCHAR(50) NOT NULL,
    `color` VARCHAR(5) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
    FOREIGN KEY (`color`) REFERENCES `color` (`name`)
);
CREATE TABLE `ply` (
    `id` VARCHAR(50) NOT NULL,
    `turn` INTEGER NOT NULL,
    `move` VARCHAR(10) NOT NULL,
    `game_id` VARCHAR(50) NOT NULL,
    `player_id` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`turn`, `player_id`),
    FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
    FOREIGN KEY (`player_id`) REFERENCES `player` (`id`)
);
CREATE TABLE `comment` (
    `id` VARCHAR(50) NOT NULL,
    `text` TEXT NOT NULL,
    `public` BOOLEAN NOT NULL,
    `ply_id` VARCHAR(10) NOT NULL,
    `author_id` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`ply_id`) REFERENCES `ply` (`id`)
);
INSERT INTO `color`(`name`)
VALUES ('WHITE'),
    ('BLACK');
INSERT INTO `game_result`(`id`, `description`)
VALUES ('1-0', 'White won'),
    ('0-1', 'Black won'),
    ('1/2-1/2', 'Draw'),
    ('-', 'Game ongoing');
-- +migrate Down
DROP TABLE IF EXISTS `comment`;
DROP TABLE IF EXISTS `ply`;
DROP TABLE IF EXISTS `player`;
DROP TABLE IF EXISTS `color`;
DROP TABLE IF EXISTS `game`;
DROP TABLE IF EXISTS `game_result`;