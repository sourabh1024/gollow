CREATE DATABASE IF NOT EXISTS test;

USE test;

CREATE TABLE dummy_data (
  id BIGINT NOT NULL AUTO_INCREMENT,
  pid INT NOT NULL ,
  first_name TEXT,
  last_name TEXT,
  balance DECIMAL(20,2),
  max_credit DECIMAL(20,2),
  max_debit DECIMAL(20,2),
  score DECIMAL(20,2),
  is_active BOOLEAN NOT NULL ,
  created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);