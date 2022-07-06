package main

func main() {
	c := make(map[string]int)
	go func ()  {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}
}
