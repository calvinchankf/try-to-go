CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "username" TEXT,
    "password" TEXT,
    "created_timestamp" TIMESTAMP WITH TIME ZONE DEFAULT now(),
    "last_login_timestamp" TIMESTAMP WITH TIME ZONE DEFAULT now()
  );