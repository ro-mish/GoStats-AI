# GoStats-AI

In this repo, we look at performance improvements on the Go for Statistics Module. 
We examine automated code generation, AI-assisted programming, and AI-generated code.

### Automated code generation
In automated code-generation, we see several speed-ups that a developer may realize.
The time taken to develop and build boilerplate code is now drastically reduced. AI is able to identify general use-case requirements and write the relevant code to let the developer spend less time setting up and more time implementing logic. 

For my study, I used Claude 3.5 Sonnet from Anthropic. This is renowned by many as a sophisticated model for code gen, and is my personal favorite since I have worked extensively with it. 

Building these files and code components now is a matter of identifying project requirements. 

### Examining AI-Assisted Programming
For this assignment I leveraged Cursor which is an open source alternative to VS Code Co-Pilot. Because Co-Pilot did not have a free tier I decided to go with Cursor and use the free trial. This is very similar in the fact that it allows for AI assisted programming with generated components inside of the code files, as well as suggestions and a chat window. 
When using the cursor chat window I simply asked it "what would you improve about the current implementation" in the Go code. 
It returned the following:

```
Based on the Go implementation shown in cmd/GoStatsTests/main.go, here are several improvements I would suggest:
Error Handling: The code currently ignores errors from stats.Mean(). We should handle these properly:
		xMean, _ := stats.Mean(dataSet.x)
		yMean, _ := stats.Mean(dataSet.y)
Here's how it should be handled:
xMean, err := stats.Mean(dataSet.x)
if err != nil {
    fmt.Printf("Error calculating x mean: %v\n", err)
    continue
}
yMean, err := stats.Mean(dataSet.y)
if err != nil {
    fmt.Printf("Error calculating y mean: %v\n", err)
    continue
}
Modular Design: The statistical calculations should be separated into their own function for better reusability and testing:
func calculateLinearRegression(x, y []float64) (slope, intercept float64, err error) {
    xMean, err := stats.Mean(x)
    if err != nil {
        return 0, 0, fmt.Errorf("error calculating x mean: %v", err)
    }
    
    yMean, err := stats.Mean(y)
    if err != nil {
        return 0, 0, fmt.Errorf("error calculating y mean: %v", err)
    }

    var numerator, denominator float64
    for i := range x {
        numerator += (x[i] - xMean) * (y[i] - yMean)
        denominator += (x[i] - xMean) * (x[i] - xMean)
    }
    
    if denominator == 0 {
        return 0, 0, fmt.Errorf("division by zero in slope calculation")
    }
    
    slope = numerator / denominator
    intercept = yMean - slope*xMean
    
    return slope, intercept, nil
}
Data Structure: Create a proper type for the dataset:
type Dataset struct {
    X []float64
    Y []float64
    Name string
}
Input Validation: Add validation for the input data:
func validateDataset(x, y []float64) error {
    if len(x) != len(y) {
        return fmt.Errorf("x and y arrays must have same length")
    }
    if len(x) == 0 {
        return fmt.Errorf("dataset cannot be empty")
    }
    return nil
}
Benchmarking: Add proper benchmarking tests in the test file:
func BenchmarkLinearRegression(b *testing.B) {
    x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
    y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        slope, intercept, _ := calculateLinearRegression(x, y)
        _ = slope
        _ = intercept
    }
}
Documentation: Add proper documentation comments for exported functions and types:
// Dataset represents a pair of x and y values for linear regression analysis
type Dataset struct {
    X []float64
    Y []float64
    Name string
}

// calculateLinearRegression performs simple linear regression on the provided x and y values
// and returns the slope and intercept of the regression line
func calculateLinearRegression(x, y []float64) (slope, intercept float64, err error)
These improvements would make the code more robust, maintainable, and easier to test. They also follow Go best practices for error handling and documentation.
The improved implementation would be more production-ready and would maintain the performance advantages shown in the README while adding necessary safety checks and better structure.
```
In evaluating the chat window, we see that there are several different considerations that the LLM has made. For instance, we've seen that it's able to look at error handling, modular design, data structures, input validation, and benchmarking, which are all key considerations that we want to have in our GO program. Because GO is a statically typo language this is a lot more robust in understanding the key functionalities of the program and actually considers important aspects of the programming language. Fundamentally this is really impressive and overall saves a lot of time for a developer when evaluating their code since it's able to identify key areas that affect the code's overall readability and maintainability while not straying away from the key functionality and logic in the code itself. 

Moreover, the chat window does a great job of mentioning documentation, which is important in any code project because it allows a reader to fully gain a deeper understanding of the code base and how key components in the program work. It suggests adding more comments and also making the code more readable which in turn improves overall developer productivity as well as end user alignment with the product. 

Another underrated feature about CodeGen in this case is the UI of the IDE. In any case most of the copilot and assistant based IDEs such as VS Code Copilot and Cursors Chat and Composer are mostly the same. But along with this it allows for quick iteration and in-context understanding of the code base, improving learning and developer speed. 


### AI Generated Code

Below is a directory structure:
```
anscombe/
├── go.mod         # Module definition (AI Generated)
└── main.go        # Main program file with all code (AI Generated)
go.mod             # AI Generated mod file


prompts.txt        # All prompts that are being used by Claude 3.5 Sonnet
```

In prompts.txt, we display the conversation that I had with Claude 3.5 Sonnet. In this specific instance we asked several different questions to Claude specifically around taking in the project requirements and being able to generate code that solves the requirements. In doing so the LLM was very quick in outputting the code required for the task while also considering the project directory structure and overall functionality of the program. One key feature that I noticed was the ability for the model to generate instructions that helps a developer run all of the code in a seamless manner. This seamlessness in the output is incredibly important moving forward because it allows for not only basic code generation, it creates an end-to-end instruction set that allows for anybody to pick up code and run it!

## Conclusion
In conclusion, we see that the use of AI assistance and co-pilots in programming significantly improves developer speeds. In short, it allows us to use text-to-code modeling that can not only interpret a developer's requests but also consider autonomously the best decisions for a project and code base. 
In this case my recommendation to management is to not necessarily reduce developer workforce sizes but rather enroll subscriptions to AI assistants and co-pilots to improve developer productivity. Fundamentally, I think that humans and AI can work synergistically towards developing much more robust and advanced software beyond what is the current industry standard. In any case the state of the art will always push boundaries and humans will currently still be in the loop for this. As a general rule for this specific case, I recommend getting everybody subscriptions to either VS Code co-pilotm, Cursor, or even Windsurf by Codeium. 
Any developer friendly tool that leverages AI to generate code will be paramount in a company's success because it not only improves speed, but also the readability and maintainability of code given that certain developers may or may not follow certain practices. Having an AI that can standardize all of these different heterogeneous sources of code into one homogenous and efficient implementation is incredibly valuable. 

## Links used for this assignment

- https://techhq.com/2025/01/ai-tools-for-code-completion/
- https://codesubmit.io/blog/ai-code-tools/
- https://www.forbes.com/sites/bernardmarr/2024/06/07/generative-ai-can-write-computer-codewill-we-still-need-software-developers/