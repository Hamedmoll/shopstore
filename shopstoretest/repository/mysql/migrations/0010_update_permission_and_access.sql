-- +migrate Up
INSERT INTO `permissions` (`id`, `title`) VALUES(2, 'add_product');
INSERT INTO `access_controls` (`actor_type`, `actor_id`, `permission_id`) VALUES('admin', 1, 2);

-- +migrate Down
DELETE FROM `permissions` WHERE id = 2;
DELETE FROM `access_controls` WHERE id = 2;