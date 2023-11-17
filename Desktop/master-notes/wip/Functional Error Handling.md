---
creation-date: 2023-01-24T21:26:27
type: concept
aliases:
  - Option & Some
tags:
  - programming/functional-programming
---

Because [[0_pages/05/2023-03-05-17-29-07-77300|Pure Functions]] takes in Inputs for the algorithm, this generally expects ==meaningful and valid inputs== that a function can use and generate some outputs. 

However when functions are composed and chained, there often contains scenarios which some functions cannot generate valid outputs for the next function. 


> [!Note]- Why Default Values are not always an option?
> A good pure function only acts on the inputs without knowledge of other parts of the code. By setting up a default value, we are implicitly requiring other functions to have knowledge about the external default values which is irrelevant to the function's algorithm ([[0_pages/05/2023-03-05-17-29-07-77300#No Side-effects|side effects]]). 
> ![[2023-01-24-21-36#^default-value]]
> For example, in this diagram. If a default value is defined within `func 1`, `func 2` would need to know what value `func1` defines in order to do exception. So whenever, `func1` changes this value, `func2` needs to allow after.






---
## ℹ️  Resources
- 