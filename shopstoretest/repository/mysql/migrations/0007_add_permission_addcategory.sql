-- +migrate Up
INSERT INTO `permissions` (`id`, `title`) VALUES(1, 'add_category');

-- +migrate Down
DELETE FROM `permissions` WHERE id = 1;