CREATE TABLE IF NOT EXISTS project (
  id serial PRIMARY KEY,
  name varchar(200) NOT NULL,
  description varchar(500),
  is_remote integer NOT NULL,
  location varchar(150)
);
