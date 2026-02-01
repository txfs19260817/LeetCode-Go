# Prefix to Postfix

**Difficulty:** Easy  
**Topics:** Stack  
**Interview Stages:** OA  
**Frequency:** N/A  
**Asked By:** Affirm  
**Last Reported:** 6 days ago

Given a string prefix expression, convert it to the corresponding postfix expression.

A prefix expression is one where the operator precedes its operands. For example, `*+AB-CD` corresponds to the
infix expression `(A+B) * (C-D)`.
A postfix expression is one where the operator follows its operands. For example, `AB+CD-*` corresponds to the
same infix expression `(A+B) * (C-D)`.

Assume the input only contains single-digit numbers, single uppercase letters, and operators (`+`, `-`, `*`, `/`).
The input will always be a valid prefix expression.

## Constraints

- The length of the prefix expression will be between 1 and 100.
- The expression contains only single-digit numbers, uppercase letters, and the operators `+`, `-`, `*`, `/`.

## Examples

**Example 1**

Input: `+12`  
Output: `12+`

**Example 2**

Input: `-*345`  
Output: `34*5-`

**Example 3**

Input: `*+AB-CD`  
Output: `AB+CD-*`
