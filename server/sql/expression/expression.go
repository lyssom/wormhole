package expression

import "fmt"

type Expression interface {
	Evaluate(row Row) (interface{}, error)
	String() string
}

type Row map[string]interface{}

type ComparisonExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (c *ComparisonExpression) Evaluate(row Row) (interface{}, error) {
	leftVal, err := c.Left.Evaluate(row)
	if err != nil {
		return nil, err
	}
	rightVal, err := c.Right.Evaluate(row)
	if err != nil {
		return nil, err
	}

	switch c.Operator {
	case "=":
		return leftVal == rightVal, nil
	case "!=", "<>":
		return leftVal != rightVal, nil
	case "<":
		return compare(leftVal, rightVal) < 0, nil
	case "<=":
		return compare(leftVal, rightVal) <= 0, nil
	case ">":
		return compare(leftVal, rightVal) > 0, nil
	case ">=":
		return compare(leftVal, rightVal) >= 0, nil
	default:
		return nil, fmt.Errorf("unknown comparison operator: %s", c.Operator)
	}
}

func (c *ComparisonExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", c.Left.String(), c.Operator, c.Right.String())
}

type BooleanExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BooleanExpression) Evaluate(row Row) (interface{}, error) {
	leftVal, err := b.Left.Evaluate(row)
	if err != nil {
		return nil, err
	}
	rightVal, err := b.Right.Evaluate(row)
	if err != nil {
		return nil, err
	}

	leftBool, ok := leftVal.(bool)
	if !ok {
		return nil, fmt.Errorf("left side of AND/OR must be boolean, got %T", leftVal)
	}
	rightBool, ok := rightVal.(bool)
	if !ok {
		return nil, fmt.Errorf("right side of AND/OR must be boolean, got %T", rightVal)
	}

	switch b.Operator {
	case "AND":
		return leftBool && rightBool, nil
	case "OR":
		return leftBool || rightBool, nil
	default:
		return nil, fmt.Errorf("unknown boolean operator: %s", b.Operator)
	}
}

func (b *BooleanExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator, b.Right.String())
}

type LogicalNot struct {
	Operand Expression
}

func (l *LogicalNot) Evaluate(row Row) (interface{}, error) {
	val, err := l.Operand.Evaluate(row)
	if err != nil {
		return nil, err
	}
	boolVal, ok := val.(bool)
	if !ok {
		return nil, fmt.Errorf("NOT operand must be boolean, got %T", val)
	}
	return !boolVal, nil
}

func (l *LogicalNot) String() string {
	return fmt.Sprintf("NOT %s", l.Operand.String())
}

type ColumnRef struct {
	Table  string
	Column string
}

func (c *ColumnRef) Evaluate(row Row) (interface{}, error) {
	if val, ok := row[c.Column]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("column %s not found in row", c.Column)
}

func (c *ColumnRef) String() string {
	if c.Table != "" {
		return c.Table + "." + c.Column
	}
	return c.Column
}

type Literal struct {
	Value interface{}
}

func (l *Literal) Evaluate(row Row) (interface{}, error) {
	return l.Value, nil
}

func (l *Literal) String() string {
	return fmt.Sprintf("%v", l.Value)
}

type BinaryArithmetic struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryArithmetic) Evaluate(row Row) (interface{}, error) {
	leftVal, err := b.Left.Evaluate(row)
	if err != nil {
		return nil, err
	}
	rightVal, err := b.Right.Evaluate(row)
	if err != nil {
		return nil, err
	}

	leftNum, ok := toFloat64(leftVal)
	if !ok {
		return nil, fmt.Errorf("left side of arithmetic must be numeric, got %T", leftVal)
	}
	rightNum, ok := toFloat64(rightVal)
	if !ok {
		return nil, fmt.Errorf("right side of arithmetic must be numeric, got %T", rightVal)
	}

	switch b.Operator {
	case "+":
		return leftNum + rightNum, nil
	case "-":
		return leftNum - rightNum, nil
	case "*":
		return leftNum * rightNum, nil
	case "/":
		if rightNum == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return leftNum / rightNum, nil
	case "%":
		return float64(int64(leftNum) % int64(rightNum)), nil
	default:
		return nil, fmt.Errorf("unknown arithmetic operator: %s", b.Operator)
	}
}

func (b *BinaryArithmetic) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator, b.Right.String())
}

type AggregateFunction struct {
	FuncName string
	Arg      Expression
}

func (a *AggregateFunction) EvaluateGroup(rows []Row) (interface{}, error) {
	switch a.FuncName {
	case "COUNT":
		return float64(len(rows)), nil
	case "SUM":
		if a.Arg == nil {
			return nil, fmt.Errorf("SUM requires an argument")
		}
		var sum float64
		for _, row := range rows {
			val, err := a.Arg.Evaluate(row)
			if err != nil {
				return nil, err
			}
			f, ok := toFloat64(val)
			if !ok {
				return nil, fmt.Errorf("SUM requires numeric argument")
			}
			sum += f
		}
		return sum, nil
	case "AVG":
		if a.Arg == nil {
			return nil, fmt.Errorf("AVG requires an argument")
		}
		if len(rows) == 0 {
			return nil, nil
		}
		var sum float64
		for _, row := range rows {
			val, err := a.Arg.Evaluate(row)
			if err != nil {
				return nil, err
			}
			f, ok := toFloat64(val)
			if !ok {
				return nil, fmt.Errorf("AVG requires numeric argument")
			}
			sum += f
		}
		return sum / float64(len(rows)), nil
	case "MIN":
		if a.Arg == nil {
			return nil, fmt.Errorf("MIN requires an argument")
		}
		if len(rows) == 0 {
			return nil, nil
		}
		minVal, err := a.Arg.Evaluate(rows[0])
		if err != nil {
			return nil, err
		}
		for _, row := range rows[1:] {
			val, err := a.Arg.Evaluate(row)
			if err != nil {
				return nil, err
			}
			if compare(val, minVal) < 0 {
				minVal = val
			}
		}
		return minVal, nil
	case "MAX":
		if a.Arg == nil {
			return nil, fmt.Errorf("MAX requires an argument")
		}
		if len(rows) == 0 {
			return nil, nil
		}
		maxVal, err := a.Arg.Evaluate(rows[0])
		if err != nil {
			return nil, err
		}
		for _, row := range rows[1:] {
			val, err := a.Arg.Evaluate(row)
			if err != nil {
				return nil, err
			}
			if compare(val, maxVal) > 0 {
				maxVal = val
			}
		}
		return maxVal, nil
	default:
		return nil, fmt.Errorf("unknown aggregate function: %s", a.FuncName)
	}
}

func (a *AggregateFunction) String() string {
	if a.Arg == nil {
		return fmt.Sprintf("%s(*)", a.FuncName)
	}
	return fmt.Sprintf("%s(%s)", a.FuncName, a.Arg.String())
}

func compare(a, b interface{}) int {
	aNum, aOk := toFloat64(a)
	bNum, bOk := toFloat64(b)
	if aOk && bOk {
		if aNum < bNum {
			return -1
		} else if aNum > bNum {
			return 1
		}
		return 0
	}
	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)
	if aStr < bStr {
		return -1
	} else if aStr > bStr {
		return 1
	}
	return 0
}

func toFloat64(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case float32:
		return float64(n), true
	case float64:
		return n, true
	default:
		return 0, false
	}
}

// Compare compares two values and returns -1, 0, or 1
func Compare(a, b interface{}) int {
	return compare(a, b)
}