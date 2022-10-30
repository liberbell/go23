CREATE DATABASE 'company';

USE 'company';

CREATE TABLE IF NOT EXISTS 'employee' (
    'id' INT unsigned NOT NULL AUTO_INCREMENT,
    'name' VARCHAR(20),
    'city' VARCHAR(20), DEFAULT 'fredericton',