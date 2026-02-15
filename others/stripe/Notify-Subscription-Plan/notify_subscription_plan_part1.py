class Part1Notifier:
    """
    Part 1:
    - Input: user_accounts + schedule
    - Output: print emails in chronological order
    """

    def __init__(self, input_data):
        # Backward compatible keys: support both old and new field names.
        self.user_accounts = input_data.get("user_accounts", input_data.get("users", []))
        self.schedule = input_data.get("schedule", input_data.get("send_schedule", {}))

    def generate_notifications(self):
        records = []

        for user in self.user_accounts:
            # Base subscription window for this user.
            name = user["name"]
            plan = user["plan"]
            begin_date = int(user["begin_date"])
            duration = int(user["duration"])
            end_date = begin_date + duration

            for key, label in self.schedule.items():
                # "start" and "end" are fixed points; other keys are offsets to end_date.
                if key == "start":
                    day = begin_date
                elif key == "end":
                    day = end_date
                else:
                    day = end_date + int(key)

                records.append((day, f"[{label}] {name}, subscribe in plan {plan}"))

        # Output must be chronological.
        records.sort(key=lambda item: item[0])
        return [f"{day}: {message}" for day, message in records]


def generate_notifications(input_data):
    return Part1Notifier(input_data).generate_notifications()


def format_notifications(notifications):
    return notifications


if __name__ == "__main__":
    input_data = {
        "schedule": {
            "start": "Welcome",
            "-15": "Upcoming expiration",
            "end": "Expired",
        },
        "user_accounts": [
            {"name": "A", "plan": "X", "begin_date": 0, "duration": 30},
            {"name": "B", "plan": "Y", "begin_date": 1, "duration": 15},
        ],
    }

    actual = generate_notifications(input_data)
    expected = [
        "0: [Welcome] A, subscribe in plan X",
        "1: [Welcome] B, subscribe in plan Y",
        "1: [Upcoming expiration] B, subscribe in plan Y",
        "15: [Upcoming expiration] A, subscribe in plan X",
        "16: [Expired] B, subscribe in plan Y",
        "30: [Expired] A, subscribe in plan X",
    ]
    assert actual == expected, f"Unexpected output:\n{actual}"
