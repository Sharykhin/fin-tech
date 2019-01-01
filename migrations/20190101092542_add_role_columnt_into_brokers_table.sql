-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `brokers`
  ADD COLUMN `role` VARCHAR(45) NOT NULL AFTER `position`;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `brokers`
  DROP COLUMN `role`;