
-- CREATE TABLE `users`
-- (
--     id   bigint auto_increment,
--     name varchar(255) NOT NULL,
--     PRIMARY KEY (`id`)
-- );

INSERT INTO `users` (`name`, `age`, `email`, `created_at`)
VALUES ('Solomon', 32, '1@me.com', '2020-01-08 00:00:00'),
       ('Menelik', 30, '2@me.com', '2020-01-08 00:00:00');


-- CREATE TABLE `todos`
-- (
--     text varchar(255) NOT NULL,
--     status varchar(255) NOT NULL,
--     deleted_at datetime,
--     created_at datetime,
-- );

INSERT INTO `todos` (`text`, `status`,`deleted_at`, `created_at`, `updated_at`)
VALUES ('Shopping', '0', '', '', ''),
       ('Laundry', '1', '', '', '');
