# Go JSON/YAML Formatter

This is a simple command-line tool built in Go that reads a messy or minified JSON or YAML file and prints a clean, pretty, and colorized version to the terminal.

It automatically detects the file type (`.json`, `.yml`, or `.yaml`) and formats it accordingly.

-----

## ✨ Features

  * Formats both **JSON** and **YAML** files.
  * Automatically detects the file type based on its extension.
  * Pretty-prints the output with standard indentation.
  * **Colorizes** the output for better readability:
      * **Cyan** for JSON
      * **Magenta** for YAML
  * Validates the file content and reports errors for invalid formats.

-----

## 🚀 How to Run

### Prerequisites

  * Go (Version 1.18 or newer)

### 1\. Set Up the Project

1.  In your project directory, initialize the Go module:

    ```bash
    go mod init go-formatter
    ```

2.  Install the required dependencies:

    ```bash
    go get gopkg.in/yaml.v3
    go get github.com/fatih/color
    ```

### 2\. Run the Formatter

Use `go run .` followed by the name of the file you want to format.

#### Example: Formatting JSON

1.  Create a file `messy.json`:

    ```json
    {"name":"John","age":30,"isStudent":false,"courses":[{"title":"History","credits":3}],"address":null}
    ```

2.  Run the tool:

    ```bash
    go run . messy.json
    ```

3.  **Output (in Cyan):**

    ```json
    {
      "address": null,
      "age": 30,
      "courses": [
        {
          "credits": 3,
          "title": "History"
        }
      ],
      "isStudent": false,
      "name": "John"
    }
    ```

#### Example: Formatting YAML

1.  Create a file `messy.yaml`:

    ```yaml
    name: Jane
    age: 28
    isStudent: true
    courses:
      - title: Math
        credits: 4
      - title: Physics
        credits: 5
    address:
      street: 123 Main St
      city: Anytown
    ```

2.  Run the tool:

    ```bash
    go run . messy.yaml
    ```

3.  **Output (in Magenta):**

    ```yaml
    name: Jane
    age: 28
    isStudent: true
    courses:
        - title: Math
          credits: 4
        - title: Physics
          credits: 5
    address:
        street: 123 Main St
        city: Anytown
    ```

-----

## 🛠️ (Bonus) Build the Executable

To use this like a real command-line tool without typing `go run` every time, you can build the binary:

```bash
# This creates an executable file (e.g., 'format' or 'format.exe')
go build -o format
```

Now you can run it directly from that folder:

```bash
./format messy.json
```

-----

## 🧠 Key Concepts

  * **`os.Args`**: Used to read command-line arguments (specifically, the filename provided by the user).
  * **`os.ReadFile`**: Reads the raw byte content of the specified file.
  * **`filepath.Ext`**: Used to detect the file type by checking its extension.
  * **`json.Unmarshal` / `yaml.Unmarshal`**: Parses the raw bytes into a generic `interface{}`, which can hold any data structure.
  * **`json.MarshalIndent` / `yaml.Marshal`**: Re-serializes the `interface{}` back into a "pretty" byte slice. `json.MarshalIndent` is used to add indentation, while the `yaml.v3` library formats nicely by default.
  * **`github.com/fatih/color`**: A third-party library used to wrap the output string in ANSI color codes for a better terminal experience.
