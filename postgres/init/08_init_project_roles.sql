CREATE TABLE IF NOT EXISTS project_roles (
  project_id integer REFERENCES project(id) ON DELETE CASCADE,
  role_id integer REFERENCES role(id) ON DELETE CASCADE,
  PRIMARY KEY (project_id, role_id)
);
