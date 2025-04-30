package dice

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

var DiceFaces = [6]string{":oo:", ":chi:", ":nn:", ":uu:", ":maa:", ":ko:"}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func PlayDice(text string) (string, error){

	count := [6]int{0, 0, 0, 0, 0, 0}
	quantity := 5

	if text != ""{
		num, err := strconv.Atoi(text)
		if err != nil{
			return "", errors.New("invalid number format")
		}
		if 5 < num && num <= 8{
			quantity = num
		} else if 8 < num{
			quantity = 8
		}
	}

	message := ""

	for i := 0; i < quantity; i++{
		message += rollDice(&count)
	}

	message += detectCombo(count)

	return message, nil

}

func rollDice(count *[6]int) string{
	num := rand.Intn(6)
	count[num]++
	return DiceFaces[num]
}

func detectCombo(count [6]int) string{
	result := ""

	//chinchin
	if count[1] >= 2 && count[2] >= 2 {
		if count[0] >= 1 {	//o
			result += "\n![ochinchin](https://i.imgur.com/mNZqpBZ.gif)"
		} else {
			result += "\n![chinchin](https://i.imgur.com/kaeNkDe.gif)"
		}
	}
	//chinko
	if count[1] >= 1 && count[2] >= 1 && count[5] >= 1 {
		if count[0] >= 1 {	//o
			result += "\n![ochinko](https://i.imgur.com/06fVZiy.gif)"
		} else {
			result += "\n![chinko](https://i.imgur.com/TjWTeaZ.gif)"
		}
	}
	//manko
	if count[4] >= 1 && count[2] >= 1 && count[5] >= 1 {
		if count[0] >= 1 {	//o
			result += "\n![omanko](https://i.imgur.com/rkdEcoz.gif)"
		} else {
			result += "\n![manko](https://i.imgur.com/UIJli6O.gif)"
		}
	}
	//un
	if count[2] >= 1 && count[3] >= 1 {
		if count[5] >= 1 {	//ko
			result += "\n![unko](https://i.imgur.com/xgvtkVK.gif)"
		}
		if count[1] >= 1 {	//chi
			result += "\n![unchi](https://i.imgur.com/DlyEYwJ.gif)"
		}
	}

	return result
}