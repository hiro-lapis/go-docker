## Graphql implementation

1. edit `schema.graphqls`

In context of graphql, "define graphql schema" means this step.  
Regularly, you define,,,  

|   keyword   |                                                               mean                                                                |
|:-----------:|:---------------------------------------------------------------------------------------------------------------------------------:
|   `type`    |                                                  used to define a Graphql object                                                  |
|   `input`   | used to define an Input Object Type. represents a complex input value that can be passed as an argument to a field or a mutation. |
| `interface` |               define common attributes in types. query can refer to interface(differ types, common attribute data)                |
|   `query`   |                                                                 b                                                                 |
|             |                                                                 b                                                                 |
|             |                                                                 b                                                                 |
|             |                                                                 b                                                                 |
|             |                                                                 b                                                                 |


```schema.graphql
// ! means not null
 
type User {
  id: ID!
  name: String!
  email: String!
}

input UserInput {
  name: String!
  email: String!
}
```

2. Generate graphql model

```
go 
```
