# Parking Lot

Design a parking lot with the ability to park and remove vehicles.

## Assumptions

- The parking lot has motorcycle, compact, and large spots.
- Vehicles include motorcycles, cars, vans (and the follow-up bus).
- A motorcycle can park in any spot.
- A car can park in a compact spot or a large spot.
- A van can park in a large spot or use 3 compact spots.
- A bus can park in a large spot or use 5 compact spots.
- Compact spots used by vans or buses do not need to be adjacent.
- When multiple spot types are available, use the smallest suitable spot.

## Required Operations

- Park a vehicle in the parking lot.
- Remove a previously parked vehicle.

## Example (One Possible Flow)

- Create a lot with 1 motorcycle, 4 compact, 1 large.
- Park motorcycle m1, car c1, van v1.
- Try to park bus b1 (fails because only 3 compact left and no large).
- Remove v1, then park b1 (uses the large spot).
