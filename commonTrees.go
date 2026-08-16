package morsegode

// AlphabeticMorseTree is a Morse tree for the 26 letters of the Latin
// alphabet (A-Z) using International Morse code symbols ('.' for dot, '-'
// for dash).
var AlphabeticMorseTree = GeneralMorseTree{
	TreeRoot: &AlphabeticTree,

	LeftSymbol:  '.',
	RightSymbol: '-',
}

// AlphabeticTree is the root node of AlphabeticMorseTree. Each node stores
// the letter reached by following the sequence of dots (left) and dashes
// (right) from the root.
var AlphabeticTree = GeneralMorseTreeNode{
	letter: "",

	// .
	left: &GeneralMorseTreeNode{
		// E .
		letter: "e",

		left: &GeneralMorseTreeNode{
			// I ..
			letter: "i",

			left: &GeneralMorseTreeNode{
				// S ...
				letter: "s",

				left: &GeneralMorseTreeNode{
					// H ....
					letter: "h",
				},

				right: &GeneralMorseTreeNode{
					// V ...-
					letter: "v",
				},
			},

			right: &GeneralMorseTreeNode{
				// U ..-
				letter: "u",

				left: &GeneralMorseTreeNode{
					// F ..-.
					letter: "f",
				},

				right: &GeneralMorseTreeNode{
					// (..--)
					letter: "",
				},
			},
		},

		right: &GeneralMorseTreeNode{
			// A .-
			letter: "a",

			left: &GeneralMorseTreeNode{
				// R .-.
				letter: "r",

				left: &GeneralMorseTreeNode{
					// L .-..
					letter: "l",
				},

				right: &GeneralMorseTreeNode{
					// .-.-
					letter: "",
				},
			},

			right: &GeneralMorseTreeNode{
				// W .--
				letter: "w",

				left: &GeneralMorseTreeNode{
					// P .--.
					letter: "p",
				},

				right: &GeneralMorseTreeNode{
					// J .---
					letter: "j",
				},
			},
		},
	},

	// -
	right: &GeneralMorseTreeNode{
		// T -
		letter: "t",

		left: &GeneralMorseTreeNode{
			// N -.
			letter: "n",

			left: &GeneralMorseTreeNode{
				// D -..
				letter: "d",

				left: &GeneralMorseTreeNode{
					// B -...
					letter: "b",
				},

				right: &GeneralMorseTreeNode{
					// X -..-
					letter: "x",
				},
			},

			right: &GeneralMorseTreeNode{
				// K -.-
				letter: "k",

				left: &GeneralMorseTreeNode{
					// C -.-.
					letter: "c",
				},

				right: &GeneralMorseTreeNode{
					// Y -.--
					letter: "y",
				},
			},
		},

		right: &GeneralMorseTreeNode{
			// M --
			letter: "m",

			left: &GeneralMorseTreeNode{
				// G --.
				letter: "g",

				left: &GeneralMorseTreeNode{
					// Z --..
					letter: "z",
				},

				right: &GeneralMorseTreeNode{
					// Q --.-
					letter: "q",
				},
			},

			right: &GeneralMorseTreeNode{
				// O ---
				letter: "o",
			},
		},
	},
}
