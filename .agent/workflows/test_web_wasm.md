---
description: Build and serve the web application for testing
---
This workflow guides you through building and serving a GoCompose application for web (WASM) testing using `go run`.

## Prerequisites
- Go installed.
- Access to the `go-compose` repository.

## Steps

1. **Navigate to Repository Root**
   - Ensure you are in the root of the `go-compose` repository where `go.mod` is located.
   ```bash
   # Execute from the project root
   ```

2. **Serve Application**
   - Use `go run` to execute the serve command directly.
   - Replace `./cmd/demo/kitchen` with the path to the application package you want to serve (relative to the repo root).
   ```bash
   go run ./cmd/go-compose serve -http :8080 ./cmd/demo/kitchen
   ```
   - **Note:** This command blocks execution.
   - If you are an agent executing this:
     - Use `run_command` with `WaitMsBeforeAsync` set to a small value (e.g. 2000 ms) to ensure it starts, and capture the Command ID.
     - Ensure `SafeToAutoRun` is false if you are unsure, but usually `serve` is safe.
   - **Important:** Do *not* run `go install`.

3. **Access in Browser**
   - Open `http://localhost:8080` in your web browser.
   - If using the browser tool, navigate to this URL to perform UI tests.

4. **Cleanup**
   - When finished testing, terminate the background process using the `send_command_input` tool with `Terminate: true` and the Command ID from step 2.
