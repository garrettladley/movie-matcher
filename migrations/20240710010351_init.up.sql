CREATE DOMAIN nu_email_domain AS varchar(64) CHECK (value ~ '^[a-zA-Z]+\.[a-zA-Z]+@northeastern\.edu$');

CREATE DOMAIN applicant_name_domain AS varchar(256) CHECK (value !~ '[/()"<>\\{}]');

CREATE TABLE IF NOT EXISTS applicants (
    email nu_email_domain PRIMARY KEY,
    applicant_name applicant_name_domain NOT NULL,
    created_at timestamp WITH time zone NOT NULL DEFAULT NOW(),
    token uuid UNIQUE NOT NULL,
    prompt text NOT NULL,
    solution text NOT NULL
);

CREATE TABLE IF NOT EXISTS submissions (
    submission_id uuid PRIMARY KEY,
    token uuid NOT NULL REFERENCES applicants (token) ON DELETE CASCADE,
    score smallint NOT NULL,
    submission_time timestamp WITH time zone NOT NULL DEFAULT NOW()
);