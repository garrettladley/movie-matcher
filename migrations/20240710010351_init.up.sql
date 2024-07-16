CREATE DOMAIN nuid_domain AS varchar(9) CHECK (value ~ '^[0-9]{9}$');

CREATE DOMAIN applicant_name_domain AS varchar(256) CHECK (value !~ '[/()"<>\\{}]');

CREATE TABLE IF NOT EXISTS applicants (
    nuid nuid_domain PRIMARY KEY,
    applicant_name applicant_name_domain NOT NULL,
    registration_time timestamp WITH time zone NOT NULL DEFAULT NOW(),
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