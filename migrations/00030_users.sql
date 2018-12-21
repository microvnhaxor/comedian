-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `users` ADD `real_name` VARCHAR(255) DEFAULT '';
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `users` DROP `real_name`;