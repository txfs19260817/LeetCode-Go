# Top N Similar Movies

You are given:

- a movie name `movie`
- a ratings map `ratings`, where `ratings[movie_name]` is that movie's rating
- a list of similarity pairs `similarities`, where each pair `(a, b)` means movie `a` is similar to movie `b`
- an integer `n`

Similarity is transitive. If `A` is similar to `B` and `B` is similar to `C`, then `A` is also considered similar to `C`.

Return the top `n` similar movies with the highest ratings.

Assumptions used by this implementation:

- Similarity is undirected: `(A, B)` also means `(B, A)`.
- The input movie itself is not included in the recommendations.
- If multiple similar movies have the same rating, sort them by movie name in ascending order.

## Example
**Input:**
```text
movie = "A"
ratings = {
    "A": 6,
    "B": 7,
    "C": 8,
    "D": 9,
}
similarities = [("A", "B"), ("B", "C")]
n = 1
```

**Output:**
```text
["C"]
```

Explanation:

- `A` is similar to `B`.
- `B` is similar to `C`.
- Therefore `A` is also similar to `C`.
- Among movies similar to `A`, movie `C` has the highest rating.
