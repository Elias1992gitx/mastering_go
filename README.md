# Mastering Go - Learning Project

A comprehensive collection of Go programming examples and exercises designed to demonstrate fundamental concepts, best practices, and common patterns in Go development.

## 📁 Project Structure

This project is organized into focused modules, each demonstrating specific Go programming concepts:

### Core Modules

#### 🔢 **cla/** - Command Line Arguments
- **Purpose**: Demonstrates command-line argument processing
- **Key Concepts**: 
  - Parsing command-line arguments using `os.Args`
  - Type conversion with `strconv.ParseFloat()`
  - Finding minimum and maximum values from input
  - Error handling patterns (with warnings about ignoring errors)

#### ⚠️ **newError/** - Error Handling
- **Purpose**: Shows proper error handling patterns in Go
- **Key Concepts**:
  - Creating custom errors with `errors.New()`
  - Nil error checking
  - Error comparison and handling
  - Production-ready error handling practices

#### 🖨️ **printing/** - Output Formatting
- **Purpose**: Demonstrates various printing and formatting methods
- **Key Concepts**:
  - `fmt.Print()` vs `fmt.Println()`
  - `fmt.Printf()` with format specifiers
  - String and integer formatting
  - Output formatting best practices

#### 📤 **stdOut/** - Standard Output
- **Purpose**: Shows how to write to standard output
- **Key Concepts**:
  - Using `io.WriteString()` for output
  - Command-line argument processing
  - Standard output redirection

#### 📥 **stdIn/** - Standard Input
- **Purpose**: Demonstrates reading from standard input
- **Key Concepts**:
  - Using `bufio.Scanner` for input reading
  - File pointer operations with `os.Stdin`
  - Defer statements for resource cleanup
  - Continuous input processing

#### 📝 **stdErr/** - Standard Error
- **Purpose**: Shows how to write to standard error stream
- **Key Concepts**:
  - Error output using `os.Stderr`
  - Separating normal output from error output
  - Command-line argument validation

#### 📦 **package/** - Package Management
- **Purpose**: Demonstrates Go module and package usage
- **Key Concepts**:
  - Go modules with `go.mod`
  - External package dependencies
  - GitHub package imports
  - Module versioning

### Additional Examples

#### 🔧 **other/** - Miscellaneous Examples
- **curly.go**: Demonstrates Go's strict curly brace placement rules
- **packageNotUsed.go**: Shows unused import handling with blank identifier
- **aSourceFile.go**: Basic Go program structure example

## 🚀 Getting Started

### Prerequisites
- Go 1.24.4 or higher
- Git (for package dependencies)

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd mastering_go
   ```

2. Navigate to any module directory and run the examples:
   ```bash
   cd cla
   go run cla.go 1.5 2.3 0.8 4.1
   ```

## 📚 Learning Path

This project is designed to be explored in the following order:

1. **Start with `printing/`** - Basic output concepts
2. **Move to `cla/`** - Command-line argument processing
3. **Explore `stdIn/`, `stdOut/`, `stdErr/`** - I/O operations
4. **Study `newError/`** - Error handling patterns
5. **Examine `package/`** - Module and dependency management
6. **Review `other/`** - Additional Go syntax and rules

## 🛠️ Running Examples

Each module can be run independently. Here are some example commands:

```bash
# Command line arguments example
cd cla
go run cla.go 10.5 2.3 8.9 1.1

# Error handling example
cd newError
go run newError.go

# Input/Output examples
cd stdIn
go run stdIn.go

cd stdOut
go run stdOut.go "Hello, World!"

cd stdErr
go run stdErr.go "Error message"

# Package example
cd package
go run getPackage.go
```

## 📖 Key Learning Objectives

- **Command-Line Processing**: Understanding `os.Args` and argument parsing
- **Error Handling**: Proper error creation, checking, and handling patterns
- **I/O Operations**: Working with standard input, output, and error streams
- **Package Management**: Using Go modules and external dependencies
- **Go Syntax**: Understanding curly brace placement and import rules
- **Type Conversion**: Converting between string and numeric types
- **Resource Management**: Using defer statements for cleanup

## ⚠️ Important Notes

- The `cla/` module includes warnings about ignoring error values in production code
- Some examples in `other/` are commented out to demonstrate syntax rules
- The `package/` module requires internet access to download dependencies

## 🤝 Contributing

This is a learning project. Feel free to:
- Add new examples
- Improve existing code
- Add more comprehensive comments
- Suggest additional learning modules

## 📄 License

This project is for educational purposes. Feel free to use and modify the code for learning Go programming.

---

**Happy Go Programming! 🐹** 