# Mini API Demo

A streamlined API management example with automatic code generation from OpenAPI specifications. This project demonstrates how to effectively manage multiple API endpoints using code generation to minimize boilerplate.

## Features

- OpenAPI-driven development
- Automatic code generation
- Clean separation between API definitions and implementation
- Simple, in-memory implementation for quick testing

## Project Structure

```
mini-api-demo/
├── api-specs/            # OpenAPI specifications
│   ├── todos-api.yaml    # Todo API definition
│   └── categories-api.yaml # Categories API definition
├── generated/            # Generated code from OpenAPI
│   ├── todos/
│   └── categories/
├── handlers/             # API implementations
│   ├── todos.go
│   └── categories.go
├── cmd/
│   └── server/           # API server
│       └── main.go
├── gen.go                # Code generator
├── go.mod                # Go module definition
└── Makefile              # Build commands
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- oapi-codegen (`go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest`)

### Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/mini-api-demo.git
   cd mini-api-demo
   ```

2. Update the module name in `go.mod` and import paths in all files:
   ```
   # Replace 'yourusername' with your GitHub username or organization
   sed -i 's|github.com/yourusername/mini-api-demo|github.com/your-actual-username/mini-api-demo|g' $(find . -name "*.go")
   ```

3. Generate code from OpenAPI specs:
   ```
   make generate
   ```

4. Build the server:
   ```
   make build
   ```

5. Run the server:
   ```
   make run
   ```

## API Endpoints

The server exposes the following endpoints:

### Todo API

- `GET /todos` - List all todos
- `POST /todos` - Create a new todo
- `GET /todos/{todoId}` - Get a specific todo
- `PUT /todos/{todoId}` - Update a todo
- `DELETE /todos/{todoId}` - Delete a todo

### Categories API

- `GET /categories` - List all categories
- `POST /categories` - Create a new category
- `GET /categories/{categoryId}` - Get a specific category
- `PUT /categories/{categoryId}` - Update a category
- `DELETE /categories/{categoryId}` - Delete a category

### Documentation

- `GET /api-docs/todos` - Todo API specification
- `GET /api-docs/categories` - Categories API specification

## Adding New APIs

To add a new API:

1. Create a new OpenAPI specification in the `api-specs` directory
2. Run `make generate` to generate the code
3. Create a handler implementation in the `handlers` directory
4. Register the handler in `cmd/server/main.go`

## How to Scale This Approach

This approach scales well to handle many APIs by:

1. **Code Generation**: Eliminates boilerplate and ensures consistency
2. **Clear Separation**: API definitions are decoupled from implementation
3. **Standard Patterns**: All handlers follow the same patterns
4. **Versioning**: APIs can be versioned independently

For production, you might want to add:

- Authentication and authorization
- Database persistence
- Request validation
- Rate limiting
- Monitoring and logging
- API gateway integration

## License

MIT