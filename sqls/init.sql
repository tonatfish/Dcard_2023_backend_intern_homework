CREATE TABLE pages (
  id VARCHAR(20) NOT NULL,
  data JSON NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE EXTENSION pg_cron;

SELECT cron.schedule('0-59/2 * * * *', 'DELETE FROM pages WHERE created_at < now() - interval ''1 mins''');