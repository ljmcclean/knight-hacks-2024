CREATE TABLE IF NOT EXISTS profile_skills (
  profile_id UUID REFERENCES profile(id) ON DELETE CASCADE,
  skill_id integer REFERENCES skill(id) ON DELETE CASCADE,
  PRIMARY KEY (profile_id, skill_id)
);
