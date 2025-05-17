# Gabriel Guimaraes - Personal Website

A simple personal website.

## Local Setup

1. Clone Repo

    ```bash
    git clone git@github.com:gabrielg2020/personal-website.git && cd personal-website
    ```

2. Run Server

    ### On local machine (for development)

    ```bash
    go run main.go
    ```

    ### On docker contianer

    ```bash
    docker build -f Dockerfile -t gabriels-personal-website .
    ```

    ```bash
    docker run -p 8080:8080 gabriels-personal-website:latest
    ```

### Create new blog post

1. Use cli tool

    ```bash
    go run cli/main.go <name_of_post>
    ```

2. Edit post

    ```bash
    vim content/<name_of_post>.html
    ```

3. Refresh webpage

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

Built with 💻 by Gabriel Guimaraes
