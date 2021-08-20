 # Problem Solving
## Scenario
    Ainun has 20 cakes and 25 apples. She wants to bundle those cakes and apples into boxes and give them to her friends. How many boxes that Ainun can make? And how many cakes and apples every box have?

## Terms and Conditions
    Apples and cakes divided evenly every box

## Solution
    To solve the problem we use Euler algorithm to divided apples and cakes evenly
<br>
Let's jump to the app ;)
<br><br>

## Requirements

 - Golang version 1.12+
 - Basic knowledge about go mod https://github.com/golang/go/wiki/Modules
 - With new feature from Golang `1.11+` you can work on this project outside `GOPATH`

## Run
- **command:**
    ```zsh
    go run main.go -name="Ainun" -apple=25 -cake=20
    ```
- **result:**
    ```zsh
    Ainun has 20 cakes and 25 apples. She wants to bundle those cakes and apples into boxes and give them to her friends. 
    - How many boxes that Ainun can make? 5 boxes
    - How many cakes and apples every box have? 4 apples and 5 cakes
    ```