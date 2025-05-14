-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TYPE action AS ENUM ('click', 'doubleclick', 'text', 'goto');

CREATE TYPE assertion AS ENUM ('text');

CREATE TABLE steps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    is_unique BOOLEAN DEFAULT FALSE,
    selector VARCHAR(255) NOT NULL,
    action VARCHAR(255) NULL,
    assertion VARCHAR(255) NULL,
    assertion_target VARCHAR(255) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE components (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    head_step_id UUID NOT NULL references steps(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE suites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    head_component_id UUID NULL references components(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE suite_components_nodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    suite_id UUID NOT NULL references suites(id),
    component_id UUID NOT NULL references components(id),
    next_component_id UUID NULL references components(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
