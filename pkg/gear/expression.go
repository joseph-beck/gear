package gear

type ExpressionType int

const (
	EmptyExpression ExpressionType = iota
	CharExpression
	LiteralExpression
	ChoiceExpression
	NotExpression
	OptionalExpression
	SequenceExpression
	ZeroOrMoreExpression
	OneOrMoreExpression
	NamedRuleExpression
)

type Expression interface {
	Type() ExpressionType
	Evaluate(*Context) (Result, error)
}

type Result struct {
	CST CST
}

var (
	// Letter is a predefined expression that matches any single letter character (a-z, A-Z).
	Letter = &Choice{
		Value: []Expression{
			&Char{
				Value: 'a',
			},
			&Char{
				Value: 'b',
			},
			&Char{
				Value: 'c',
			},
			&Char{
				Value: 'd',
			},
			&Char{
				Value: 'e',
			},
			&Char{
				Value: 'f',
			},
			&Char{
				Value: 'g',
			},
			&Char{
				Value: 'h',
			},
			&Char{
				Value: 'i',
			},
			&Char{
				Value: 'j',
			},
			&Char{
				Value: 'k',
			},
			&Char{
				Value: 'l',
			},
			&Char{
				Value: 'm',
			},
			&Char{
				Value: 'n',
			},
			&Char{
				Value: 'o',
			},
			&Char{
				Value: 'p',
			},
			&Char{
				Value: 'q',
			},
			&Char{
				Value: 'r',
			},
			&Char{
				Value: 's',
			},
			&Char{
				Value: 't',
			},
			&Char{
				Value: 'u',
			},
			&Char{
				Value: 'v',
			},
			&Char{
				Value: 'w',
			},
			&Char{
				Value: 'x',
			},
			&Char{
				Value: 'y',
			},
			&Char{
				Value: 'z',
			},
			&Char{
				Value: 'A',
			},
			&Char{
				Value: 'B',
			},
			&Char{
				Value: 'C',
			},
			&Char{
				Value: 'D',
			},
			&Char{
				Value: 'E',
			},
			&Char{
				Value: 'F',
			},
			&Char{
				Value: 'G',
			},
			&Char{
				Value: 'H',
			},
			&Char{
				Value: 'I',
			},
			&Char{
				Value: 'J',
			},
			&Char{
				Value: 'K',
			},
			&Char{
				Value: 'L',
			},
			&Char{
				Value: 'M',
			},
			&Char{
				Value: 'N',
			},
			&Char{
				Value: 'O',
			},
			&Char{
				Value: 'P',
			},
			&Char{
				Value: 'Q',
			},
			&Char{
				Value: 'R',
			},
			&Char{
				Value: 'S',
			},
			&Char{
				Value: 'T',
			},
			&Char{
				Value: 'U',
			},
			&Char{
				Value: 'V',
			},
			&Char{
				Value: 'W',
			},
			&Char{
				Value: 'X',
			},
			&Char{
				Value: 'Y',
			},
			&Char{
				Value: 'Z',
			},
		},
	}

	// Digit is a predefined expression that matches any single digit character (0-9).
	Digit = &Choice{
		Value: []Expression{
			&Char{
				Value: '0',
			},
			&Char{
				Value: '1',
			},
			&Char{
				Value: '2',
			},
			&Char{
				Value: '3',
			},
			&Char{
				Value: '4',
			},
			&Char{
				Value: '5',
			},
			&Char{
				Value: '6',
			},
			&Char{
				Value: '7',
			},
			&Char{
				Value: '8',
			},
			&Char{
				Value: '9',
			},
		},
	}
)
