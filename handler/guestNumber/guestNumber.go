package guestNumber

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var number string

func GenerateNumber(c *gin.Context) {
	num := ""

	for i := 0; i < 4; i++ {
		res, _ := rand.Int(rand.Reader, big.NewInt(9))

		if strings.Contains(num, res.String()) {
			i--
		} else {
			num += res.String()
		}
	}

	fmt.Println(num)

	number = num

	c.JSON(http.StatusOK, gin.H{
		"number": num,
	})
}

func Guest(c *gin.Context) {
	if number == "" {
		return
	}

	var data GuestRequest

	if err := c.Bind(&data); err != nil {
		log.Fatal(err)
		return
	}

	serverNumber := strings.Split(number, "")
	sp := strings.Split(data.Number, "")

	a := 0
	for i := range number {
		if serverNumber[i] == sp[i] {
			a += 1
		}
	}

	b := 0
	for i := range number {
		if strings.Contains(number, sp[i]) && sp[i] != serverNumber[i] {
			b += 1
		}
	}

	var msg string

	if a == 4 {
		msg = "猜中！"
		number = ""
	} else {
		msg = strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
	}

	fmt.Println(msg)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
