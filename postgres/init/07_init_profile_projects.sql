CREATE TABLE IF NOT EXISTS profile_projects (
  profile_id UUID REFERENCES profile(id) ON DELETE CASCADE,
  project_id integer REFERENCES project(id) ON DELETE CASCADE,
  PRIMARY KEY (profile_id, project_id)
);
