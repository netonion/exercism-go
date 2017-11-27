package erratum

// Use is an error handling exercise
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	for res, err = o(); err != nil; res, err = o() {
		if _, ok := err.(TransientError); !ok {
			return
		}
	}

	defer func() {
		if rec := recover(); rec != nil {
			if ferr, ok := rec.(FrobError); ok {
				res.Defrob(ferr.defrobTag)
			}
			err = rec.(error)
		}
		res.Close()
	}()
	res.Frob(input)

	return
}
