# Go Architecture
Demo common Architecture for web servers golang
- This Project demo API create/get animals

## Sample 3 Layered Architecture
Diagram:
![alt text](https://github.com/awesome-academy/Go-Architecture/blob/layered-architecture/images/layered-architectures.png?raw=true)

- Using [godotenv](https://github.com/joho/godotenv) to loads environment variables from .env files
- Using [gorilla/mux](https://github.com/gorilla/mux) to router

Source Code:

**Branch:** [layered-architecture](https://github.com/awesome-academy/Go-Architecture/tree/layered-architecture)

### How to run this project?
We can run this project with Docker. Here, I am providing both ways to run this project.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone https://github.com/awesome-academy/Go-Architecture.git

# Move to the project root directory
cd Go-Architecture
```

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:3000

### Example API Request and Response
- Create animal
  - request
    ```
    curl --location --request POST 'http://localhost:3000/animal' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'Name=Name Test' \
    --data-urlencode 'Age=10'
    ```
  
  - response
  
  ```json
    {
      "ID": 5,
      "Name": "Name Test",
      "Age": 10
    }
  ```

### Contributing to Go Architecture

All pull requests are welcome.
