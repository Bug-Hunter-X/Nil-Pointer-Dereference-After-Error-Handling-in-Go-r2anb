func Calculate(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := Calculate(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Handle the error gracefully
        return // Exit to prevent further execution which might lead to a nil dereference
    }
    fmt.Println("Result:", result)
}
