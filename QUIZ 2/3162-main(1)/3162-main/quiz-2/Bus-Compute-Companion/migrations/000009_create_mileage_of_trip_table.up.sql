CREATE TABLE IF NOT EXISTS mileage_of_trip (
  id bigserial PRIMARY KEY,
  beginning_location_id BIGINT REFERENCES locations_of_stop(id),
  destination_location_id BIGINT REFERENCES locations_of_stop(id),
  total_miles BIGINT
);