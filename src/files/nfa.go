//Author: Kieran O'Halloran
//Adapted from:
//https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

package files

type state struct {
	//store letter with state. Value is 0
	symbol rune
	//the two arrows that come from the state. pointers to other states
	edge1 *state
	edge2 *state
}

type nfa struct { //keep track of initial and accept states
	initial *state
	accept  *state
}

func Poregtonfa(pofix string) *nfa { //return pointer to nfa
	nfaStack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1] // : means give me everything off the statck up to the last element but not including it
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial //join together and push the conceited fragment back to the nfa stack

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept}) //pops 2 fragments off the stack, joins the accept state of frag1 to initial state of frag2, push the new fragment onto the nfa stack

		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1] //: means give me everything off the statck up to the last element but not including it
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '*':
			//only pop one frag off the nfa stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//new frag is the old frag with 2 extra states
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) //push new frag to nfa stack

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept} //label new accept state with symbol r

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfaStack[0]

}

//function to add states to array of pointers (current/next)
func addState(l []*state, s *state, a *state) []*state {
	//adding desired state to array
	l = append(l, s)
	//checking it isnt the accept state and its arrow label is e
	if s != a && s.symbol == 0 {
		//adding states meeting above condition
		l = addState(l, s.edge1, a)
		//if there's a 2nd edge add relevant state
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	//returning new current array
	return l
}

//function returning boolean on a regular expression matching a given string
func PostFixMatch(postFix string, str string) bool {
	//set returned boolean to false to start
	matched := false

	//convert postfix regex to an non-determistic finite automata
	pfixNfa := Poregtonfa(postFix)

	//array of state pointers containing all the current states of nfa
	current := []*state{}

	//after reading a character from the string this contains all next states
	next := []*state{}

	current = addState(current[:], pfixNfa.initial, pfixNfa.accept)
	//looping through given string one rune at a time
	for _, r := range str {
		//looping through array of current states
		for _, c := range current {
			//if current rune is the same as the arrow labels of the current state
			if c.symbol == r {
				next = addState(next[:], c.edge1, pfixNfa.accept)
			}
		}

		//setting current array to next(making move) for next rune being read in and resetting next array to null
		current, next = next, []*state{}
	}
	//loop through current states after nfa is finished
	for _, c := range current {
		//if one of the current states is an accept state its a match
		if c == pfixNfa.accept {
			matched = true
			break
		}
	}

	//returning whether it's a match or not
	return matched
}
