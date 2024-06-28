# Small CRUD API

This is a simple CRUD API built with Go and the Gin web framework.

## Endpoints

### Get All Articles
- **URL**: `http://127.0.0.1:8080/articles`
- **Method**: `GET`
- **Description**: Retrieves all articles from the database.

### Get a Single Article
- **URL**: `http://127.0.0.1:8080/article/{title}`
- **Method**: `GET`
- **Description**: Retrieves a single article based on its title.
- **Example**: `http://127.0.0.1:8080/article/my-first-article`

### Create a New Article
- **URL**: `http://127.0.0.1:8080/article`
- **Method**: `POST`
- **Description**: Creates a new article in the database.
- **Request Body**:
  ```json
  {
    "title": "My New Article",
    "content": "This is the content of my new article."
  }
  ```

### Update an Article
- **URL**: `/article/:id`
- **Method**: `PATCH`
- **Description**: Updates an existing article based on its ID.
- **URL Parameter**:
  - `id`: The ID of the article to update.
- **Request Body**:
  - `title` (optional): The new title of the article.
  - `content` (optional): The new content of the article.

### Delete an Article
- **URL**: `/article/:id`
- **Method**: `DELETE`
- **Description**: Deletes an article based on its ID.
- **URL Parameter**:
  - `id`: The ID of the article to delete.

## Project Structure

```plaintext
small-crud/
│
├── controller/
│   └── controller.go
├── models/
│   └── model.go
│   └── setup.go
└── cmd/
    └── main.go
