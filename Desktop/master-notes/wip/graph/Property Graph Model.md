---
id: 67FE17EC-23E6-4C6D-91AB-00A71E333950
creation-date: 2022-11-10T14:30:03
type: concept
aliases:
  - Graph Model
  - Network Model
  - Graph Database
tags:
  - machine-learning/graph-model
  - database/graph-data
---

## 📦 What is a Property Graph Model?

To construct a <term>Property Graph Model</term>, there are 3 main components:

-   _Node_ (_Vertex_) - the object & core entity that makes up the whole graph
-   _Edge_ - describes the link/ relationship between 2 different Nodes with a *direction*
-   _Properties_ - which can the data/ attributes of the a Node/ Edge

At a higher level, a Node is a representation of some-thing (e.g., Location) and data provides more details about the “thing” (e.g., name, population, coordinate of the Location). Similarly, an Edge represents some relationship (e.g, someone lives in some-Location) and its properties provides more contextual details under a specific relationship.

> [!example]-
> In this example, there are:
> - `Person` Node with property: name “Tom Hanks”
> - `Movie` Node with property: title “Apollo 13”
> - The two Nodes contain a relationship where “Tom Hanks” `ACTED_IN` in “Apollo 13”. Since only “Tom Hanks” acted as “Jim Lovell” in “Apollo 13” but not other people and movies. This relationship specific information is stored as a Property within this relationship.

To help emerge network’s relationships and patterns, both Nodes and Edges can have their own type to catalog similarities.
-   Nodes can have _Labels_ to categorised into a common type of object (commonly a Noun)
-   Edges needs to have ******Types****** to describe common relationships (commonly a Verb)

> [!example]-
> Nodes can represent different entity of different Label (Blue: People, Orange: Movie, Green: Company)
> Relationships can be in different Types (Black: “cast in” somewhere, Red: “wrote” something, Purple: “property of” someone).

---
## 🆚 vs Relational Model

### **Similarity**

-   **Table representation**
    
    Same set of Node’s Labels can be think of as a relational tables and using the Edges to traverse to another Node is like having two different tables conducting a relational joins.
    

### **Differences**

-   **No Strict Schema**
    
    Graph Model doesn’t need to have a strict schema as it is commonly represented as key-value paired representation. This allows flexibility to maintain similar but different relationship even though one type of Node don’t have such attributes.
    
    <aside> 💭 Regions in different countries would have different structures, e.g., _states_ in US vs ********regions******** in Europe. Even though they share similar concepts, this would be difficult to link in a Relational Model but comes easy in Graph Model.
    
    </aside>
    
-   ****************************************************No Prior Knowledge about Relationship****************************************************
    
    Because Graph Model is based on Nodes traversing (walking along the paths), it doesn’t need to know all the intermediate patterns between the starting and ending Nodes. As long as the traversing matches the pattern in the query, Graph Model can easily return the data.
    
    Relational Model requires all the intermediate relations in all tables to complete the joins. In order words, to query the starting and ending Nodes in a Relational Table, we need to have all the relationships on how to join all the tables in between.
    
    [[Designing Data-Intensive Applications](https://www.notion.so/Designing-Data-Intensive-Applications-771fed2ccabc486ba59040b2a9cc663d) p.52]
    
    ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/82dd75b0-bfe6-404b-93b0-3122a196e9b0/Untitled.png)
    
    ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0bb7a57d-0306-41a4-af6f-f9f918a0acfe/Untitled.png)
    

---
## 🗣️ Common Use Cases

Most of the analyses we understand are in a structure of “table” or matrix but Graphs are useful to represent real world data that can be represented by different connections, e.g., road traffic, information flow, people connections, knowledge linkage, etc.



---
## ℹ️ Resources
-   [Graph Algorithms](https://www.notion.so/Graph-Algorithms-b886677a0ef04482887faf78668fbfb4) (chapter 1)
-   [Designing Data-Intensive Applications](https://www.notion.so/Designing-Data-Intensive-Applications-771fed2ccabc486ba59040b2a9cc663d) (p.50)
-   [Avon](https://www.notion.so/Avon-6d98177f83c1447791c474f320abedb0) 93