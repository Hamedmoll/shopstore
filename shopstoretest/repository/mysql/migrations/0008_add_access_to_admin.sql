-- +migrate Up;
INSERT INTO `access_controls` (`actor_type`, `actor_id`, `permission_id`) VALUES('admin', 1, 1);

-- +migrate Down
DELETE FROM `access_controls` WHERE id = 1;