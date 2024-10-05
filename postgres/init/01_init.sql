CREATE TABLE IF NOT EXISTS profile (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(100) NOT NULL,
  email varchar(200) UNIQUE NOT NULL,
  description varchar(500),
  location varchar(150),
  password text NOT NULL
);
