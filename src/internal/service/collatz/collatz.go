package collatz

func Collatz(factor int)(steps int){

	switch factor {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		case 4:
			return 3
		default:
	}	
	
	for {
		switch factor % 2 {
			case 0:
				factor = even(factor)
			default:
				factor = odd(factor)
		}

		steps = steps + 1

		if factor == 1 {
			break
		}
	}

	return steps + 1
}

func even(factor int)(newFactor int){
	return factor/2
}

func odd(factor int)(newFactor int){	
	return (factor*3)+1
}