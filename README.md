# URL-Shortener
URL-Shortener application 

### Steps to run this app
1. Make sure you have go1.20 version in your system
2. Make sure you have docker desktop
3. clone or download this project 
4. To run the project run  `docker-compose up --build`
5. There are 4 Apis in this application.
    - Create Short Url API 
        - URL:  `http://localhost:8080/v1/create/short-url`
        - Method : **POST**
        - Request Body: `{"long_url": "https://www.hasnain.com/top"}`
        - Response: `{
                           "long_url": "https://www.hasnain.com/top",
                           "short_url": "localhost:8080/a38597eFAS"
                    } `
    - Redirect API
        - Method: **GET**
        - URL : `http://localhost:8080/{shortUrl}`

    - Health API
        - Method :  **GET**
        - URL : `http://localhost:8080/health`

    - Metrics API to fetch Top 3 domain
        - Method: **GET**
        - URL : `http://localhost:8081/v1/metrics/top?param=K`
        - Domain: Replace the K with any value it will return top k domains

    

