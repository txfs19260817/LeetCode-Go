MOTORCYCLE = 0
COMPACT = 1
LARGE = 2


class ParkingLot:
    def __init__(self, motorcycle, compact, large):
        # Available counts by spot type index.
        self.spots = [motorcycle, compact, large]
        # Vehicle id -> (spot_type, count) for easy removal.
        self.parked = {}
        # Preferred allocations by vehicle type, in priority order.
        self.rules = {
            "motorcycle": [(MOTORCYCLE, 1), (COMPACT, 1), (LARGE, 1)],
            "car": [(COMPACT, 1), (LARGE, 1)],
            "van": [(LARGE, 1), (COMPACT, 3)],
            "bus": [(LARGE, 1), (COMPACT, 5)],
        }

    def park(self, vid, vtype):
        # Allocate the first rule that fits the remaining capacity.
        if not vid or vid in self.parked or vtype not in self.rules:
            return False
        for spot, count in self.rules[vtype]:
            if self.spots[spot] >= count:
                self.spots[spot] -= count
                self.parked[vid] = (spot, count)
                return True
        return False

    def remove(self, vid):
        # Free the previously allocated spots for this vehicle id.
        alloc = self.parked.pop(vid, None)
        if alloc is None:
            return False
        spot, count = alloc
        self.spots[spot] += count
        return True

    def available(self, spot_type):
        # Remaining count for a given spot type.
        if 0 <= spot_type <= LARGE:
            return self.spots[spot_type]
        return 0


if __name__ == "__main__":
    lot = ParkingLot(1, 4, 1)
    print(lot.park("m1", "motorcycle"))
    print(lot.park("c1", "car"))
    print(lot.park("v1", "van"))
    print(lot.park("b1", "bus"))
    print(lot.remove("v1"))
    print(lot.park("b1", "bus"))
