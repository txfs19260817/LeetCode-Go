# Design Worker Management System

Introduce operations for adding workers, registering their entering or leaving the office
and retrieving information about the amount of time that they have worked.

`boolean addWorker(String workerId, String position, int compensation)` — should add the
workerId to the system and save additional information about them: their position and
compensation.

If the workerId already exists, nothing happens and this operation should return "false".
If the workerId was successfully added, return "true".
workerId and position are guaranteed to contain only English letters and spaces.

`String registerWorker(String workerId, int timestamp)` — should register the time when
the workerId entered or left the office. The time is represented by the timestamp. Note
that registerWorker operation calls are given in the increasing order of the timestamp
parameter.

If the workerId doesn't exist within the system, nothing happens and this operation should
return "invalid_request".
If the workerId is not in the office, this operation registers the time when the workerId
entered the office.
If the workerId is already in the office, this operation registers the time when the
workerId left the office.
If the workerId's entering or leaving time was successfully registered, return "registered".

`int get(String workerId)` — should return a number representing the total calculated
amount of time that the workerId spent in the office.

The amount of time is calculated using finished working sessions only. It means that if
the worker has entered the office but hasn't left yet, this visit is not considered in the
calculation.
If the workerId doesn't exist within the system, return -1.

## Example

**Input:**

```
["OfficeManager", "addWorker", "addWorker", "registerWorker", "registerWorker", "get", "registerWorker", "registerWorker", "registerWorker", "get", "get", "registerWorker"]
[[], ["Ashley", "Middle Developer", 150], ["Ashley", "Junior Developer", 100], ["Ashley", 10], ["Ashley", 25], ["Ashley"], ["Ashley", 40], ["Ashley", 67], ["Ashley", 100], ["Ashley"], ["Walter"], ["Walter", 120]]
```

**Output:**

```
[null, true, false, "registered", "registered", 15, "registered", "registered", "registered", 42, -1, "invalid_request"]
```

OfficeManager officeManager = new OfficeManager();
officeManager.addWorker("Ashley", "Middle Developer", 150); // Returns true.
officeManager.addWorker("Ashley", "Junior Developer", 100); // Returns false. The same worker ID already exists within the system.
officeManager.registerWorker("Ashley", 10); // Returns "registered". "Ashley" entered the office at timestamp 10.
officeManager.registerWorker("Ashley", 25); // Returns "registered". "Ashley" left the office at timestamp 25.
officeManager.get("Ashley"); // Returns 15. "Ashley" spent 25 - 10 = 15 time units in the office.
officeManager.registerWorker("Ashley", 40); // Returns "registered".
officeManager.registerWorker("Ashley", 67); // Returns "registered".
officeManager.registerWorker("Ashley", 100); // Returns "registered".
officeManager.get("Ashley"); // Returns 42. "Ashley" spent (25 - 10) + (67 - 40) = 42 time units in the office.
officeManager.get("Walter"); // Returns -1. id "Walter" was never added to the system.
officeManager.registerWorker("Walter", 120); // Returns "invalid_request". "Walter" was never added to the system.
