CREATE TABLE IF NOT EXISTS role_skills (
  role_id integer REFERENCES role(id) ON DELETE CASCADE,
  skill_id integer REFERENCES skill(id) ON DELETE CASCADE,
  PRIMARY KEY (project_id, role_id)
);
