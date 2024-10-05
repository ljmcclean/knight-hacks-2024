CREATE TABLE IF NOT EXISTS project (
  id serial PRIMARY KEY,
  name varchar(200) NOT NULL,
  description varchar(500) NOT NULL,
  is_remote integer NOT NULL,
  location varchar(150) NOT NULL,
  skills varchar(60)[] NOT NULL,
  user_id uuid
);
