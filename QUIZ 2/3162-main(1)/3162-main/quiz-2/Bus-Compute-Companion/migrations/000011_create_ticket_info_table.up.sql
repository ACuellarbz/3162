CREATE TABLE IF NOT EXISTS ticket_info (
  id bigserial PRIMARY KEY,
  route_id BIGINT REFERENCES bus_schedule(id),
  ticket_price BIGINT,
  number_of_tickets_available BIGINT
);