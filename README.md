# Simple_REST_API
Simple REST API that allows for a handful of HTTP requests.
Utilizes gRPC for content delivery at https://github.com/HenryNgai/Simple_RPC_API


# Clone the repository
```bash
git clone git@github.com:HenryNgai/Simple_REST_API.git
```

# Run the application
```bash
cd Simple_REST_API
go run ./cmd/api/main.go
```

# Using Application
```
go to http://localhost:8080/
press register (says email but supports any type in "email" field) - Can be changed using binding in struct)
presss login to test if user was successfully registered
```

JSON token saved message should appear after logging in

Validate json token by clicking validate button

# Todo - My notes
- [x] Implement salting and hashing for password
- [x] Create dependency layer? Update db
- [x] Implement JWT generation
- [x] Implement JWT verification
- [ ] Make gRPC calls to Simple_RPC_API

# More Todo - After REST API
- [ ] RPC API (different repo maybe)
- [ ] Graphql API (different repo maybe)
- [ ] Additional NoSQL database for a service?
