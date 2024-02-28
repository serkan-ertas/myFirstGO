package main

/*****NOTES******

	func printSomething(value interface {})		//any value allowed
	func printSomething(value any)				//any value allowed

	func printSomething(value any) {
		switch value.(type) {					//can check the type
		case int:								//with this special type
			//..
		case string:
			//..
		case float64:
			//..
		default:
			//..
		}
	}

	typedVal, ok := value.(int)		// no matter if value is int or not typedVal type is int
									// ok = True if value is int


	func add[T int|float64|string](a,b T) T {	// generics feature
		return a + b							// T can be int/float/str and func will return T type
	}


******NOTES******/
