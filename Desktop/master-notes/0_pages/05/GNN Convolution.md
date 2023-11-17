---
id: D33BC00D-5A12-4FAE-8E75-62A37233D026
creation-date: 2023-02-11T20:46:29 
type: concept
alias: ['GCN']
tags: [ machine-learning, machine-learning/graph-model ]
---

## Convolution Learning in GNN

*Convolution Learning* is about having a set of "filters" to scan over a set of neighbours and **learn the importance/ weightings of the neighbours**. 

In the context of *Graph Neural Networks* ([[0_pages/02/2023-02-26-17-06-30-81300|GNN]]), the goal of Convolution Learning is: 
- traverse over Neighbours Nodes for $d$ times
- conducted [[0_pages/05/2023-02-26-17-06-58-65100|localised convolution]] on Nodes $h_v$ over Neighbours embeddings $h_u$
- learn a set of weights $w_d$  of Neighbours embeddings at each $d$-th traversal

> [!Tip]- CNN vs GNN Convolution
> - CNN Conv - an element in a matrix ($x_{i,j}$) always contains some neighbours elements within a $d$-sized *convolution window*
> - GNN Conv - a single Node ($h_v$) doesn't always contain a fixed number of neighbours, i.e, some Nodes may have more connectivity and some have less

---
## Pre-Neural Feature Engineerings

![[3_hidden/_images/Pasted image 20230211205711.png|500]]

Traditional graph ML models are based feature engineering on the number of *traversal* ($d$) on the *degree of Nodes* ($D$). One way to do this is to apply *polynomial filters* $p(L)$. 

For each iteration of a Node $v$ ($v' = p_w(L) \times v$): 
- use a Graph Traversal Map (usually a *Laplacian Matrix* $L$)
- define a the number of traversals $d$ to expand outwards
- apply $d$ number of filters ($w$) on neighbours [[0_pages/05/2023-02-26-23-19-27-95200|Node Embedding]] for each traversal
- summation into a new embedding $v'$

Mathematically, each convolution has a computation of: $v' = (w_0I + w_1 L + w_2L^2 + ...) \times v$

---
## Neural Network Convolutions

Convolution by traversing Nodes is essentially a [[0_pages/02/2023-02-26-23-19-48-97800#Nodes Neighbours Messages Passing|Node-wise Message Passing]], which **aggregates information of neighbouring Nodes**. 



---
## ℹ️  Resources
- 