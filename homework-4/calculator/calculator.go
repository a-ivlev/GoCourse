package calculator

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"
	"strings"
)

const help = `
Справка по работе с калькулятором  
Калькулятор может производить следующие операции
(сложение, вычитание, умножение, деление),
проводить операции со следующими константами:
pi		- число Пи  = 3.141592653589793
e		- число E   = 2.718281828459045
phi		- число Phi = 1.618033988749895
получено из последовательности Фибоначчи
его ещё называют Золотым сечением
а также проводить операции с функциями:
sqrt(x)		- Квадратный корень из х
abs(x)		- модуль числа x
log(x y) 	- логарифм x  y
ln(x) 		- десятичный логарифм x
sin(x)		- синус x
cos(x) 		- косинус x
tan(x) 		- тангенс x
arcsin(x)	- арксинус x
arccos(x)	- арккосинус x
arctan(x)	- арктангенс x
max (x1 ... xn) - вычисляет максимальное значение среди введённых значений
min (x1 ... xn) - вычисляет минимальное значение среди введённых значений
exit 		- Наберите exit чтобы выйти из калькулятора.
`

// Help выводит на экран справку по калькулятору.
func Help() {
	fmt.Println(help)
}

// Calculate рассчитает данные выражения, переданного в строке
func Calculate(expr string) (float64, error) {
	// Поиск управляющих операторов в строке и построение дерева
	// Вернет корневой элемент в переменную root все математические действия, построенные
	// в виде дерева
	// Node будет содержать в себе действие для вычисления
	// exprNode() даст доступ к ветвям дерева, исходящим из корня/ноды
	//
	// type Expr interface {
	//  Node
	//  exprNode()
	// }
	root, err := parser.ParseExpr(expr)

	if err != nil {
		return -1, err
	} else {
		return eval(root)
	}
}

type Func struct {
	Name string
	Args int
	Func func(args ...float64) float64
}

var funcMap map[string]Func

// Создание массива функций для обработки данных
func init() {
	funcMap = make(map[string]Func)
	funcMap["sqrt"] = Func{
		Name: "sqrt",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sqrt(args[0])
		},
	}
	funcMap["abs"] = Func{
		Name: "abs",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Abs(args[0])
		},
	}
	funcMap["log"] = Func{
		Name: "log",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Log(args[0]) / math.Log(args[1])
		},
	}
	funcMap["ln"] = Func{
		Name: "ln",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Log(args[0])
		},
	}
	funcMap["sin"] = Func{
		Name: "sin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sin(args[0])
		},
	}
	funcMap["cos"] = Func{
		Name: "cos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Cos(args[0])
		},
	}
	funcMap["tan"] = Func{
		Name: "tan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Tan(args[0])
		},
	}
	funcMap["arcsin"] = Func{
		Name: "arcsin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Asin(args[0])
		},
	}
	funcMap["arccos"] = Func{
		Name: "arccos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Acos(args[0])
		},
	}
	funcMap["arctan"] = Func{
		Name: "arctan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Atan(args[0])
		},
	}
	funcMap["max"] = Func{
		Name: "max",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Max(args[0], args[1])
		},
	}
	funcMap["min"] = Func{
		Name: "min",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Min(args[0], args[1])
		},
	}
}

// Разбор полученных данных
func eval(expr ast.Expr) (float64, error) {
	switch expr.(type) {
	case *ast.BasicLit:
		return basic(expr.(*ast.BasicLit))
	// Бинарные выражения
	case *ast.BinaryExpr:
		return binary(expr.(*ast.BinaryExpr))
	// Вложенные вычисления
	case *ast.ParenExpr:
		return eval(expr.(*ast.ParenExpr).X)
	// Обработка посредством математических функций
	case *ast.CallExpr:
		return call(expr.(*ast.CallExpr))
	// Случай для обработки констант
	case *ast.Ident:
		return ident(expr.(*ast.Ident))
	default:
		return -1, errors.New("Не удалось распознать оператор")
	}
}

// Разбор чисел
func basic(lit *ast.BasicLit) (float64, error) {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)

		if err != nil {
			return -1, err
		} else {
			return float64(i), nil
		}
	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)

		if err != nil {
			return -1, err
		} else {
			return i, nil
		}
	default:
		return -1, errors.New("Неизвестный аргумент")
	}
}

func binary(expr *ast.BinaryExpr) (ret float64, err error) {
	x, err1 := eval(expr.X)
	y, err2 := eval(expr.Y)
	ret = -1

	if (err1 == nil) && (err2 == nil) {

		switch expr.Op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = float64(int64(x) % int64(y))
		case token.AND:
			ret = float64(int64(x) & int64(y))
		case token.OR:
			ret = float64(int64(x) | int64(y))
		case token.XOR:
			ret = math.Pow(x, y)
		default:
			err = errors.New("Неизвестный бинарный оператор")
		}
	} else {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
	}

	return
}

// Обработка констант
func ident(id *ast.Ident) (float64, error) {
	switch n := strings.ToLower(id.Name); n {
	case "pi":
		return math.Pi, nil
	case "e":
		return math.E, nil
	case "phi":
		return math.Phi, nil
	default:
		return -1, errors.New("Неизвестная константа " + n)
	}
}

// Обработка функциональных операторов с помощью созданного массива функций для обработки
func call(c *ast.CallExpr) (float64, error) {
	switch t := c.Fun.(type) {
	case *ast.Ident:
	default:
		_ = t
		return -1, errors.New("Неизвестный тип функции")
	}

	ident := c.Fun.(*ast.Ident)

	args := make([]float64, len(c.Args))
	for i, expr := range c.Args {
		var err error
		args[i], err = eval(expr)
		if err != nil {
			return -1, err
		}
	}

	name := strings.ToLower(ident.Name)

	if val, ok := funcMap[name]; ok {
		if len(args) == val.Args {
			return val.Func(args...), nil
		} else {
			return -1, errors.New("Слишком много аргументов для " + name)
		}
	} else {
		return -1, errors.New("Неизвестная функция " + name)
	}
}
