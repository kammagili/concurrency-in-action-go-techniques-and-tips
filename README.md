# concurrency-in-action-go-techniques-and-tips

This project contains all the code examples from the "Concurrency in Action: Go Techniques and Tips" talk from GopherCon Israel 2025.

You can find the PDF of the talk in this repository: `Concurrency in Action-Go Techniques and Tips - gili kamma.pdf`.

Each Go file demonstrates a different concurrency pattern or technique discussed in the presentation.

## How to Run the Examples

1. **Open a terminal and navigate to this directory:**
   ```sh
   cd concurrency-in-action-go-techniques-and-tips
   ```
2. **Run a specific example:**
   - Open `main.go` in your editor.
   - Uncomment the function call for the example you want to run (e.g., `GoroutineExample()`).
   - Only one example should be uncommented at a time for clarity.
   - Save the file.
   - Run the example:
     ```sh
     go run .
     ```

3. **To run benchmarks or tests:**
   ```sh
   go test -bench=.
   ```