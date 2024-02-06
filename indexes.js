use("FindProMatesV2");

db.getCollection("users").createIndex({ email: 1 }, { unique: true, name: "unique_email" });

db.getCollection("users").createIndex({ username: 1 }, { unique: true, name: "unique_username" });

db.getCollection("projects").createIndex(
  { name: 1, owner: 1 },
  { unique: true, name: "unique_project_name_for_user" }
);
