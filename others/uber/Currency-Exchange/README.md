# Currency Exchange

## Description

Given a list of currency relationships with exchange values (e.g., BGP -> USD: 109.0). Find the best exchange rate from `from` currency to `to` currency with the following assumptions:

- The rates are bi-directional. If the rate from A to B is `r`, the rate from B to A can be calculated as `1 / r`.
- The goal is to find the maximum achievable exchange rate, which is the highest value obtained by multiplying the rates along a valid exchange path.
- If the conversion is not possible, return -1.

Implement the `CurrencyConverter` class:

- `CurrencyConverter(String[] fromArr, String[] toArr, double[] rateArr)` Initialize the CurrencyConverter class with currency pairs and exchange rates.
- `getBestRate(String from, String to)` Return the best exchange rate that can be obtained from `from` currency to `to` currency.

### Constraints

- `1 <= fromArr.length, toArr.length, rateArr.length <= 1000`
- `0 < rateArr[i] <= 10^3`
- All currency names consist of English letters and have a length between 1 and 10

## Example

```
Input:
["CurrencyConverter", "getBestRate", "getBestRate", "getBestRate", "getBestRate"]
[[["GBP", "USD", "USD", "USD", "CNY"], ["JPY", "JPY", "GBP", "CAD", "EUR"], [155.0, 112.0, 0.9, 1.3, 0.14]], ["USD", "JPY"], ["JPY", "BGP"], ["XYZ", "GBP"], ["CNY", "CAD"]]

Output:
[null, 139.5, 0.00803, -1.0, -1.0]

Explanation:
CurrencyConverter currencyConverter = new CurrencyConverter(fromArr, toArr, rateArr);
currencyConverter.getBestRate("USD", "JPY"); // Return 139.5, the best exchange path is an indirect path USD -> BGP -> JPY = (1/0.9) * 155 = 139.5, better than the direct paths like USD -> JPY = 112.0
currencyConverter.getBestRate("JPY", "BGP"); // Return 0.00803, the best exchange path is JPY-> USD -> BGP.
currencyConverter.getBestRate("XYZ", "BGP"); // Returns -1.0, because "XYZ" doesn't exist in the currency array.
currencyConverter.getBestRate("CNY", "CAD"); // returns -1.0, because there's no exchange path exists from "CNY" to "CAD".
```
