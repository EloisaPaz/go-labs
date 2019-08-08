# Calculator microservice using Golang

#### Simple calculator microservice using Golang. The calculator supports the basics operations.

#### Open the terminal and type the following command to run the calculator:

    go run calculator.go

#### The calculator will be available in the following url

For instructions of use:
- http://localhost:8080/calc/ 

For additions:
- http://localhost:8080/calc/sum/{firstNum}/{secondNum}

For subtractions: 
- http://localhost:8080/calc/sub/{firstNum}/{secondNum}

For multiplications:
- http://localhost:8080/calc/mul/{firstNum}/{secondNum}

For divisions: 
- http://localhost:8080/calc/div/{firstNum}/{secondNum}

You can also check the history through the following url:
- http://localhost:8080/calc/history