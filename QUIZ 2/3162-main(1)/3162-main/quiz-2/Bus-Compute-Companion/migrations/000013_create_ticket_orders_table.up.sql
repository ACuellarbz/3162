CREATE TABLE IF NOT EXISTS ticket_orders (
  id bigserial PRIMARY KEY,
  ticket_id BIGINT REFERENCES ticket_info(id),
  user_id BIGINT REFERENCES users_info(id),
  seat_id BIGINT,
  purchased_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);