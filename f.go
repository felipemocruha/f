package f

func Reduce[T any, S any](fn func(acc T, val S) (T, error), init T, coll []S) (T, error) {
	for _, val := range coll {
		res, err := fn(init, val)
		if err != nil {
			return init, err
		}

		init = res
	}

	return init, nil
}

func Map[T any, S any](fn func(val T) (S, error), coll []T) ([]S, error) {
	var init []S

	one := func(acc []S, val T) ([]S, error) {
		res, err := fn(val)
		if err != nil {
			return nil, err
		}
		acc = append(acc, res)

		return acc, nil
	}

	return Reduce(one, init, coll)
}

func Filter[T any](fn func(val T) (bool, error), coll []T) ([]T, error) {
	var init []T

	one := func(acc []T, val T) ([]T, error) {
		ok, err := fn(val)
		if err != nil {
			return nil, err
		}

		if ok {
			acc = append(acc, val)
		}

		return acc, nil
	}

	return Reduce(one, init, coll)
}
