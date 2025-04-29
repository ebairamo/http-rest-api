func validateCard(cardNumber string) bool {
	var odd []int
	var even []int
	var card []int
	var resultodd int
	var resulteven int
	var resultfinal int

	for _, char := range strings.TrimSpace(cardNumber) {
		str := string(char)
		num, err := strconv.Atoi(str)
		fmt.Println(cardNumber)
		if err != nil {
			fmt.Println("Ошибка преобразования символа:", err)
			fmt.Println(str)
			return false
		}
		card = append(card, num)
	}

	if len(card) < 13 {
		fmt.Println("Длина карты меньше 13 цифр")
		return false

	}

	for i := 0; i < len(card); i++ {
		if i%2 == 0 {
			odd = append(odd, card[i])
		} else {
			even = append(even, card[i])
		}
	}

	for i := 0; i < len(odd); i++ {
		multiply := odd[i] * 2
		if multiply > 9 {
			multiply -= 9
		}
		resultodd += multiply
	}

	for i := 0; i < len(even); i++ {
		resulteven += even[i]
	}

	resultfinal = resultodd + resulteven

	if resultfinal%10 == 0 {

	}
	return true
}

func validateCardNumber(cardNumber string) bool {
	// Удаляем пробелы, если они есть
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	// Проверяем, что номер карты содержит только цифры и длина номера от 13 до 16 цифр
	if len(cardNumber) < 13 || len(cardNumber) > 16 {
		return false
	}

	// Реализуем алгоритм Луна
	var end []int
	var notend []int
	var card []int
	var resultend int
	var resultnotend int

	for _, char := range strings.TrimSpace(cardNumber) {
		num := int(char - '0') // Преобразуем byte в int
		card = append(card, num)
	}
	card = reverse(card)
	fmt.Println(card)
	for i := 0; i < len(card); i++ {
		if i%2 == 0 {
			end = append(end, card[i])
		} else {
			notend = append(notend, card[i])
		}
	}
	for i := 0; i < len(end); i++ {
		multiply := end[i] * 2
		if multiply > 9 {
			multiply -= 9
		}
		resultend += multiply
	}
	for i := 0; i < len(notend); i++ {
		resultnotend += notend[i]
	}
	fmt.Println("notresultend", resultnotend)
	fmt.Println("resultend", resultend)
	fmt.Println("end", end)
	fmt.Println("notend", notend)

	return true
}