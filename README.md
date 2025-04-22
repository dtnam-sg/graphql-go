Generate : ```chmod +x ./generate && ./generate```

Run: ``` go run server.go ```

#Query Example 

```graphql
query getUsers {
  users {
    users {
      id
      name
      email
      baseInfo {
        status
        createdBy
        createdTime
        updatedBy
      }
      roles {
        id
        name
        description
      }
      functions {
        id
        name
        description
      }
    }
    total
  }
}

query userFiltered {
  users(filter: { name: "User Test 4" }, pagination: { page: 1, size: 1 }) {
    users {
      id
      name
      email
      baseInfo {
        status
        createdBy
        createdTime
        updatedBy
      }
      roles {
        id
        name
        description
      }
      functions {
        id
        name
        description
      }
    }
    total
  }
}

query user {
  user(id: "T57") {
    id
    name
    email
    roles {
      id
      name
      description
    }
    functions {
      id
      name
      description
    }
  }
}

query roles {
  roles {
    id
    name
    description
  }
}

query functions {
  functions {
    id
    name
    description
  }
}

query function {
  function(id: "T11") {
    id
    name
    description
  }
}

mutation createUser {
  createUser(
    input: {
      name: "User Test 4"
      email: "loo1oooo7272727@gmail.com"
      roleIds: ["T61"]
      functionIds: ["T11", "T12"]
    }
  ) {
    user {
      name
      email
      id
      roles {
        id
        name
        description
      }
      functions {
        id
        name
        description
      }
    }
    success
    message
  }
}

mutation createRole {
  createRole(input: { name: "Namdt-2", description: "ahuahauahau" }) {
    success
    role {
      id
      name
      description
      __typename
    }
  }
}
```