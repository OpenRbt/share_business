# wash-bonus

## Using

1. Build your service

    > **HINT**: If your auth-srv is private repository, yor must [connecting to GitHub with SSH](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/connecting-to-github-with-ssh)


    ```bash
    bash build --full
    ```

2. Run PostgreSQL in Docker container

    ```bash
    make db_start
    ```

3. Run your service

    ```bash
    ./bin/wash-bonus
    ```

4. Shutdown

    - stop the service itself
    - if you don't need data from your database, stop it:

    ```bash
    make db_stop
    ```
