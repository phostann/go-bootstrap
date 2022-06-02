CREATE TYPE "gender_enum" AS ENUM (
  'male',
  'female'
);

CREATE TYPE "role_enum" AS ENUM (
  'admin',
  'employee',
  'user'
);

CREATE TABLE "users"
(
    "id"         bigserial PRIMARY KEY,
    "username"   varchar UNIQUE NOT NULL,
    "password"   varchar        NOT NULL,
    "avatar"     varchar        NOT NULL,
    "email"      varchar UNIQUE NOT NULL,
    "gender"     gender_enum    NOT NULL DEFAULT ('male'::gender_enum),
    "role"       role_enum      NOT NULL DEFAULT ('user'::role_enum),
    "created_at" timestamptz    NOT NULL DEFAULT (now()),
    "updated_at" timestamptz    NOT NULL DEFAULT (now()),
    "delete_at"  timestamptz
);
