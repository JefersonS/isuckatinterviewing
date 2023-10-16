DROP TABLE IF EXISTS questions;
CREATE TABLE questions (
  id         INT AUTO_INCREMENT NOT NULL,
  question      VARCHAR(128) NOT NULL,
  answer     VARCHAR(255) NOT NULL,
  itSucks      INT NOT NULL DEFAULT 0,
  youSuck   INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

INSERT INTO questions
  (question, answer)
VALUES
  ('What does the D stands for in SOLID?', 'DEZZ NUTS'),
  ('What is the engine behind PostgreSQL?', 'InnoDB'),
  ('What is the event loop in JS?', 'Events looping around'),
  ('Is React rective?', 'HtmX');
