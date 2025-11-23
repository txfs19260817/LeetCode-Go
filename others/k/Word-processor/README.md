## Word wrap

We are building a word processor and we would like to implement a "word-wrap" functionality.

Given a list of words followed by a maximum number of characters in a line, return a collection of strings
where each string element represents a line that contains as many words as possible, with the words in
each line being concatenated with a single '-' (representing a space, but easier to see for testing). The
length of each string must not exceed the maximum character length per line.

Your function should take in the maximum characters per line and return a data structure representing all
lines in the indicated max length.

Examples:
```
words1 = [ "The", "day", "began", "as", "still", "as", "the",
"night", "abruptly", "lighted", "with", "brilliant",
"flame" ]
wrapLines(words1, 13) "wrap words1 to line length 13" =>
[ "The-day-began",
"as-still-as",
"the-night",
"abruptly",
"lighted-with",
"brilliant",
"flame" ]
wrapLines(words1, 20) "wrap words1 to line length 20" =>
[ "The-day-began-as",
"still-as-the-night",
"abruptly-lighted",
"with-brilliant-flame" ]

words2 = [ "Hello" ]
wrapLines(words2, 5) "wrap words2 to line length 5" =>
[ "Hello" ]

words3 = [ "Hello", "world" ]
wrapLines(words3, 5) "wrap words3 to line length 5" =>
[ "Hello",
"world" ]

words4 = ["Well", "Hello", "world" ]
wrapLines(words4, 5) "wrap words4 to line length 5" =>
[ "Well",
  "Hello", 
  "world" ]
  
words5 = ["Hello", "HelloWorld", "Hello", "Hello"]
wrapLines(words5, 20) "wrap words 5 to line length 20 =>
[ "Hello-HelloWorld",
    "Hello-Hello" ]

All Test Cases: 
words,  max line length wrapLines(words1, 13)
wrapLines(words1, 20)
wrapLines(words2, 5)
wrapLines(words3, 5)
wrapLines(words4, 5)
wrapLines(words5, 20)
n = number of words OR total characters
```

## Reflow

We are building a word processor and we would like to implement a "reflow" functionality that also applies full
justification to the text.

Given an array containing lines of text and a new maximum width, re-flow the text to fit the new width.
Each line should have the exact specified width. If any line is too short, insert '-' 
(as stand-ins for spaces) between words as equally as possible until it fits.

Note: we are using '-' instead of spaces between words to make testing and visual verification of the results easier.

```javascript
lines = ["The day began as still as the",
        "night abruptly lighted with",
        "brilliant flame"]

reflowAndJustify(lines, 24) => // "reflow lines and justify to length 24"

["The--day--began-as-still",
"as--the--night--abruptly",
"lighted--with--brilliant",
"flame"] // <--- a single word on a line is not padded with spaces

reflowAndJustify(lines, 25) "reflow lines and justify to length 25" =>
[ "The-day-began-as-still-as",
    "the-----night----abruptly",
    "lighted---with--brilliant",
    "flame" ]

reflowAndJustify(lines, 26) "reflow lines and justify to length 26" =>
[ "The--day-began-as-still-as",
    "the-night-abruptly-lighted",
    "with----brilliant----flame" ]

reflowAndJustify(lines, 40) "reflow lines and justify to length 40" =>
[ "The--day--began--as--still--as-the-night",
    "abruptly--lighted--with--brilliant-flame" ]

reflowAndJustify(lines, 14) "reflow lines and justify to length 14" =>
['The--day-began',
    'as---still--as',
    'the------night',
    'abruptly',
    'lighted---with',
    'brilliant',
    'flame']
```

## Word Processor

We are building a word processor and we would like to implement word wrapping functionality that is a bit smarter.

Some users want to have the resulting lines to be as balanced as possible for higher readability. Therefore, our application should divide a given text into a sequence of lines so that every line spans at most some fixed width and the difference between different lines are as little as possible.

We define the balance score of a collection of strings as follows: find the longest line. For each line, sum the square of
the length difference between this line and the longest line.

Given a string of text and a maximum number of characters in a line, return a list of strings which result in the lowest balance score (on a tie, return any of the valid solutions).

Note: we are using '-' instead of spaces between words to make testing and visual verification of the results easier.

Examples:

```
text0 = "one two add a three four"

balancedWrapLines(text0, 7) "balance the lines of text0 wrapping to a max length of 7"=>

[ "one",  // 2*2
  "two",  // 2*2
  "add-a",// 0
  "three",// 0
  "four"] // 1*1 = score of 9 

Wrong answer:

[ "one-two", // 0
  "add-a"  // 2*2
  "three", // 2*2
  "four"]  // 3*3 = score of 17
 
text1 = "Seven six eight thirteen four five sixteen"

balancedWrapLines(text1, 15) "balance the lines of text1 wrapping to a max 15 length"=>

[ "Seven-six-eight",  // 0
"thirteen-four",    // 2*2
"five-sixteen" ]    // 3*3 = score of 13


text2 = "XXXXXX X XXXX X X XXXXX";

balancedWrapLines(text2, 9) "balance the lines of text2 wrapping to a max 9 length"=>

Correct answer (each line wrapped to a different length <= 9):

[ "XXXXXX-X",   // 0*0
"XXXX-X",     // 2*2
"X-XXXXX" ]   // 1*1 = score of 5

Or

[ "XXXXXX",     // 2*2
"X-XXXX-X",   // 0*0
"X-XXXXX" ]   // 1*1 = score of 5

Wrong answer (wrapping to the same length for each line):

[ 'XXXXXX-X',   // 0
'XXXX-X-X',   // 0
'XXXXX' ]     // 3*3 = score of 9

All Test Cases:
text,  max line length
balancedWrapLines(text1, 15)
balancedWrapLines(text2, 9)

n = number of words OR total characters
m = length of longest line
```