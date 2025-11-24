# Camping
## Part 1
You are going on a camping trip, but before you leave you need to buy groceries. To optimize
your time spent in the store, instead of buying the items from your shopping list in order,
you plan to buy everything you need from one department before moving to the next.

Given an unsorted list of products with their departments and a shopping list, return the
time saved in terms of the number of department visits eliminated.

Example:

```
products = [
["Cheese", "Dairy"],
["Carrots", "Produce"],
["Potatoes", "Produce"],
["Canned Tuna", "Pantry"],
["Romaine Lettuce", "Produce"],
["Chocolate Milk", "Dairy"],
["Flour", "Pantry"],
["Iceberg Lettuce", "Produce"],
["Coffee", "Pantry"],
["Pasta", "Pantry"],
["Milk", "Dairy"],
["Blueberries", "Produce"],
["Pasta Sauce", "Pantry"]
]
list1 = ["Blueberries", "Milk", "Coffee", "Flour", "Cheese", "Carrots"]

For example, buying the items from list1 in order would take 5 department visits, whereas
your method would lead to only visiting 3 departments, a difference of 2 departments.
Produce(Blueberries)->Dairy(Milk)->Pantry(Coffee/Flour)->Dairy(Cheese)-
>Produce(Carrots) = 5 department visits
New: Produce(Blueberries/Carrots)->Pantry(Coffee/Flour)->Dairy(Milk/Cheese) = 3
department visits
list2 = ["Blueberries", "Carrots", "Coffee", "Milk", "Flour", "Cheese"] => 2
list3 = ["Blueberries", "Carrots", "Romaine Lettuce", "Iceberg Lettuce"] => 0
list4 = ["Milk", "Flour", "Chocolate Milk", "Pasta Sauce"] => 2
list5 = ["Cheese", "Potatoes", "Blueberries", "Canned Tuna"] => 0
All Test Cases:
shopping(products, list1) => 2
shopping(products, list2) => 2
shopping(products, list3) => 0
shopping(products, list4) => 2
shopping(products, list5) => 0
Complexity Variable:
n: number of products
```

## Part 2
You and your friends are driving to a Campground to go camping. Only 2 of you have cars, so you will be
carpooling.

Routes to the campground are linear, so each location will only lead to 1 location and there will be no loops
or detours. Both cars will leave from their starting locations at the same time. The first car to pass
someone's location will pick them up. If both cars arrive at the same time, the person can go in either car.
Roads are provided as a directed list of connected locations with the duration (in minutes) it takes to drive
between the locations. `[Origin, Destination, Duration it takes to drive]`

Given a list of roads, a list of starting locations and a list of people/where they live, return a collection of
who will be in each car upon arrival to the Campground.

```text
 Bridgewater--(30)-->Caledonia--(15)-->New Grafton--(5)-->Campground
                                       ^
 Liverpool---(10)---Milton-----(30)-----^
 roads1 = [
    ["Bridgewater", "Caledonia", "30"], <= The road from Bridgewater to Caledonia takes 30 m
    ["Caledonia", "New Grafton", "15"],
    ["New Grafton", "Campground", "5"],
    ["Liverpool", "Milton", "10"],
    ["Milton", "New Grafton", "30"]
 ]
 starts1 = ["Bridgewater", "Liverpool"]
 people1 = [
    ["Jessie", "Bridgewater"], ["Travis", "Caledonia"],
    ["Jeremy", "New Grafton"], ["Katie", "Liverpool"]
 ]
 Car1 path: (from Bridgewater): [Bridgewater(0, Jessie)->Caledonia(30, Travis)->New Grafton(4
 Car2 path: (from Liverpool): [Liverpool(0, Katie)->Milton(10)->New Grafton(40, Jeremy)->Camp
 Output (In any order/format):
    [Jessie, Travis], [Katie, Jeremy]
    
    
 Riverport->Chester->Campground
             ^
 Halifax------^
 roads2 = [["Riverport", "Chester", "40"], ["Chester", "Campground", "60"], ["Halifax", "Chester", "40"]
 starts2 = ["Riverport", "Halifax"]
 people2 = [["Colin", "Riverport"], ["Sam", "Chester"], ["Alyssa", "Halifax"]]
 Riverport->Bridgewater->Liverpool->Campground
 Output (In any order/format):
    [Colin, Sam], [Alyssa] OR [Colin], [Alyssa, Sam]
    
    
 Riverport->Bridgewater->Liverpool->Campground
 roads3 = [["Riverport", "Bridgewater", "1"], ["Bridgewater", "Liverpool", "1"], ["Liverpool", "Campground", "1"]]
 starts3_1 = ["Riverport", "Bridgewater"]
 starts3_2 = ["Bridgewater", "Riverport"]
 starts3_3 = ["Riverport", "Liverpool"]
 people3 = [["Colin", "Riverport"], ["Jessie", "Bridgewater"], ["Sam", "Liverpool"]]
 Output (starts3_1/starts3_2):  [Colin], [Jessie, Sam] - (Cars can be in any order)
 Output (starts3_3): [Jessie, Colin], [Sam]
  
All Test Cases: (Cars can be in either order)
 carpool(roads1, starts1, people1) => [Jessie, Travis], [Katie, Jeremy]
 carpool(roads2, starts2, people2) => [Colin, Sam], [Alyssa] OR [Colin], [Alyssa, Sam]
 carpool(roads3, starts3_1, people3) => [Colin], [Jessie, Sam]
carpool(roads3, starts3_2, people3) => [Jessie, Sam], [Colin]
 carpool(roads3, starts3_3, people3) => [Jessie, Colin], [Sam]
 
Complexity Variable:
 n = number of roads
```