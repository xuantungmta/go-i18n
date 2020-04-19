// This file is generated by i18n/plural/codegen/generate.sh; DO NOT EDIT

package plural

func DefaultRules() Rules {
	rules := Rules{}
	addPluralRules(rules, []string{"vi, en, cn"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0
			if ops.NEqualsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n = 2
			if ops.NEqualsAny(2) {
				return Two
			}
			// n = 3
			if ops.NEqualsAny(3) {
				return Few
			}
			// n = 6
			if ops.NEqualsAny(4) {
				return Many
			}
			return Other
		},
	})

	return rules
}
