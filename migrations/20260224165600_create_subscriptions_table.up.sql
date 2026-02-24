CREATE TABLE subscriptions (
    id UUID PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
