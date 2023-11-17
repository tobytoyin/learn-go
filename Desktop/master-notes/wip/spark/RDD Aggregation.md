---
id: FFB842BE-A03E-48CF-99C0-EF20EFBBE8D8
creation-date: 2023-02-19T21:56:13 
type: concept
alias: ['aggregation', 'groupBy', 'groupByKey']
tags: [ technology/spark/pyspark, distributed-computing ]
---

*RDD aggregation* is a low level aggregations implementation. Similar to [[0_pages/01/2023-02-26-17-16-38-81100|groupBy in MapReduce]], an RDD aggregation takes roughly 3 steps: 

> [!Tip]- 3 Steps in RDD Aggregation
> 1. Grouping - extract group key then repartition data of the same group together
> 2. Transformation - optional Transformation done on each grouped partition
> 3. Reduce/ Action - compute a single value output for each grouped partition
> 
> Note - this is equivalently to the [[0_pages/03/2023-02-26-17-19-01-95400#General Rule for Key-Value Operations|general operations rules of key-value RDDs]]

---
## Grouping 

![[2023-02-19-22-15#^groupbykey-vs]]

*Grouping* can be done using two methods: 
- `.groupBy` - "generates" key on RDD Row using a key-generation function
- `.groupByKey` - "uses" the extracted key of [[0_pages/03/2023-02-26-17-19-01-95400|Key-Value RDD]]

> [!Example]- groupBy vs groupByKey
> ![[2023-02-19-22-15#^groupbykey-vs]]
> `groupByKey` is a follow up step from `keyBy`. This assumes that the key generated for a key-value RDD is already suitable to proceed on Grouping. 
> ---
> ![[2023-02-19-22-15#^groupBy]]
> `groupBy` operates directly on RDD's Row and generate a group by key. There's no assumption on what key already exists and suitable as an initial Grouping. 

---
## Transformations

Functions can directly `map` on grouped partitions as a [[0_pages/02/2023-03-11-15-37-42-19800|partition transformation]]. One major difference of doing `map(transform)` before/ after group by is: 
- before grouped - `map` transforms on a Row level (i.e., `Row`)
- after grouped - `map` transforms on a Partition level (i.e., `Iterable[Row]`)

![[2023-02-19-22-15#^map]]

Because repartition has already taken place, this can particularly improve performance when we need **transformations that rely on grouped aggregates** (e.g., row data - row mean, etc.,)

---
## Reduce/ Action







---
## ℹ️  Resources
- 