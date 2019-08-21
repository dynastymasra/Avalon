CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY NOT NULL, -- IDENTIFIER
    shop_id uuid NOT NULL,
    customer_id uuid NOT NULL,
    order_status varchar(25) NOT NULL,
    items JSONB NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone NULL DEFAULT NULL
);