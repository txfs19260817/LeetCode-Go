class Part2Notifier:
    """
    Part 2:
    - Includes Part 1 behavior
    - Adds subscription_changes with new_plan
    - On change day, send one extra [Updated]
    - All later emails use the latest plan
    """

    def __init__(self, input_data):
        # Backward compatible keys: support both old and new field names.
        self.user_accounts = input_data.get("user_accounts", input_data.get("users", []))
        self.schedule = input_data.get("schedule", input_data.get("send_schedule", {}))
        self.subscription_changes = input_data.get("subscription_changes", [])
        self.changes_by_user = self._group_changes_by_user()

    def _group_changes_by_user(self):
        grouped = {}
        for idx, change in enumerate(self.subscription_changes):
            user_name = change.get("name", change.get("user"))
            day = change.get("day", change.get("date", change.get("change_date")))
            # Pre-normalize changes to simplify downstream logic.
            grouped.setdefault(user_name, []).append(
                {
                    "idx": idx,
                    "day": int(day),
                    "new_plan": change.get("new_plan"),
                }
            )

        for user_name in grouped:
            grouped[user_name].sort(key=lambda item: (item["day"], item["idx"]))
        return grouped

    def _extract_plan_changes(self, user_changes):
        # Part 2 only cares about new_plan changes.
        return [change for change in user_changes if change["new_plan"] is not None]

    def _plan_on_day(self, initial_plan, plan_changes, day):
        # Last change on/before `day` wins.
        current_plan = initial_plan
        for change in plan_changes:
            if change["day"] > day:
                break
            current_plan = change["new_plan"]
        return current_plan

    def generate_notifications(self):
        records = []

        for user in self.user_accounts:
            name = user["name"]
            initial_plan = user["plan"]
            begin_date = int(user["begin_date"])
            duration = int(user["duration"])
            end_date = begin_date + duration

            user_changes = self.changes_by_user.get(name, [])
            plan_changes = self._extract_plan_changes(user_changes)

            # Extra email on each plan-change day.
            for change in plan_changes:
                current_plan = self._plan_on_day(initial_plan, plan_changes, change["day"])
                records.append(
                    (change["day"], f"[Updated] {name}, subscribe in plan {current_plan}")
                )

            # Regular schedule emails, but plan text reflects latest effective plan.
            for key, label in self.schedule.items():
                if key == "start":
                    day = begin_date
                elif key == "end":
                    day = end_date
                else:
                    day = end_date + int(key)

                plan = self._plan_on_day(initial_plan, plan_changes, day)
                records.append((day, f"[{label}] {name}, subscribe in plan {plan}"))

        # Output must be chronological.
        records.sort(key=lambda item: item[0])
        return [f"{day}: {message}" for day, message in records]


def generate_notifications(input_data):
    return Part2Notifier(input_data).generate_notifications()


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
        "subscription_changes": [
            {"name": "B", "day": 5, "new_plan": "Z"},
            {"name": "A", "day": 20, "new_plan": "Gold"},
        ],
    }

    actual = generate_notifications(input_data)
    expected = [
        "0: [Welcome] A, subscribe in plan X",
        "1: [Welcome] B, subscribe in plan Y",
        "1: [Upcoming expiration] B, subscribe in plan Y",
        "5: [Updated] B, subscribe in plan Z",
        "15: [Upcoming expiration] A, subscribe in plan X",
        "16: [Expired] B, subscribe in plan Z",
        "20: [Updated] A, subscribe in plan Gold",
        "30: [Expired] A, subscribe in plan Gold",
    ]
    assert actual == expected, f"Unexpected output:\n{actual}"
