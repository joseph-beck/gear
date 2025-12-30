package gear

import "github.com/joseph-beck/gear/pkg/errs"

type NamedRule struct {
	Value string
}

func (n *NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n *NamedRule) Evaluate(context *Context, pos uint) (Result, error) {
	// Check if already memoized
	if r, err, ok := context.Packrat().Get(n, pos); ok {
		return r, err
	}

	// Check for left recursion
	isLeftRecursive := context.Packrat().Mark(n, pos)

	if isLeftRecursive {
		// Left recursion detected - return failure to break the cycle
		// This allows the base case (second alternative) to succeed
		return Result{}, errs.FailedToMatch
	}

	rule, ok := context.Grammar().Get(n.Value)
	if !ok {
		return Result{}, errs.RuleNotFound
	}

	// Try to parse - this will use the base case on first iteration
	r, err := rule.Expression.Evaluate(context, pos)

	if err != nil {
		context.Packrat().Put(n, pos, Result{}, err)
		return Result{}, err
	}

	// Successfully parsed - create result
	tree := NewCST(rule.Name)
	tree.Children = append(tree.Children, r.CST)
	tree.Value = n.Value

	result := Result{
		Next: r.Next,
		CST:  tree,
	}

	// Store the initial result
	context.Packrat().Update(n, pos, result, nil)
	context.Packrat().Put(n, pos, result, nil)

	// Now try to grow the result by re-evaluating
	// This handles cases like expr → expr '+' digit
	for {
		// Clear memoization at this position except for our named rule
		// This forces re-evaluation of the choice/sequence while keeping our seed
		lastResult := result
		context.Packrat().Clear(pos, n)
		context.Packrat().Put(n, pos, lastResult, nil)

		// Set growing mode to prevent caching during growth
		context.SetSeeding(true)

		// Try to parse again - now the memoized result allows progress
		nextR, nextErr := rule.Expression.Evaluate(context, pos)

		// Clear growing mode
		context.SetSeeding(false)

		// If parsing failed or didn't advance further, we're done growing
		if nextErr != nil || nextR.Next <= result.Next {
			context.Packrat().Put(n, pos, result, nil)
			break
		}

		// Update result with the longer match
		nextTree := NewCST(rule.Name)
		nextTree.Children = append(nextTree.Children, nextR.CST)
		nextTree.Value = n.Value

		result = Result{
			Next: nextR.Next,
			CST:  nextTree,
		}

		// Update the memo with the grown result
		context.Packrat().Put(n, pos, result, nil)
	}

	return result, nil
}
