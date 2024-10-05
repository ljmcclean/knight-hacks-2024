CREATE TABLE IF NOT EXISTS session (
  session_id text PRIMARY KEY,
  profile_id uuid NOT NULL,
  auth_level integer NOT NULL,
  last_access timestamp NOT NULL,
  valid boolean NOT NULL DEFAULT true
);

