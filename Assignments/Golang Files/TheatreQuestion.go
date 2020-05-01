//A1_8179721
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Play struct {
	name      string
	purchased []Ticket
	showStart time.Time
	showEnd   time.Time
}

type Comedy struct {
	laughs float32
	deaths int32
	Play
}

type Tragedy struct {
	laughs float32
	deaths int
	Play
}

type Category struct {
	category  string
	basePrice float32
}
type Seat struct {
	number int32
	row    int32
	cat    *Category
}

type Ticket struct {
	cusName string
	s       *Seat
	show    *Show
}

type Theatre struct {
	seats []Seat
	shows []Show
}
type Show interface {
	getName() string
	getShowStart() time.Time
	getShowEnd() time.Time
	addPurchase(*Ticket) bool
	isNotPurchased(*Ticket) bool
	checkSoldOut() bool
}

func NewSeat(nSeat int32, rowNum int32, c *Category) *Seat {
	return &Seat{
		number: nSeat,
		row:    rowNum,
		cat:    c,
	}
}

func NewTicket(cusString string, seat *Seat, sh *Show) *Ticket {
	return &Ticket{
		cusName: cusString,
		s:       seat,
		show:    sh,
	}
}
func NewTheatre(numSeats int, showsPlaying []Show) *Theatre {
	t := make([]Seat, numSeats)
	var index = 0
	for i := 0; i < numSeats/5; i++ {
		for j := 0; j < numSeats/5; j++ {
			index = i*5 + j
			if j == 0 {
				t[index] = Seat{int32(i + 1), int32(j + 1), &Category{"Prime", 25.0}}
			} else if j == 4 {
				t[index] = Seat{int32(i + 1), int32(j + 1), &Category{"Special", 35.0}}
			} else {
				t[index] = Seat{int32(i + 1), int32(j + 1), &Category{"Standard", 15.0}}
			}
		}
	}
	return &Theatre{
		seats: t,
		shows: showsPlaying,
	}
}

func (c *Comedy) addPurchase(t *Ticket) bool {
	var b bool
	if c.isNotPurchased(t) == true {
		c.Play.purchased = append(c.Play.purchased, *t)
		b = true
	} else {
		b = false
	}
	return b
}

func (c *Comedy) isNotPurchased(t *Ticket) bool {
	for _, tickets := range c.Play.purchased {
		if tickets.s.number == t.s.number && tickets.s.row == t.s.row {
			return false
		}
	}
	return true
}

func (t *Tragedy) addPurchase(ti *Ticket) bool {
	var b bool
	if t.isNotPurchased(ti) == true {
		t.Play.purchased = append(t.Play.purchased, *ti)
		b = true
	} else {
		b = false
	}
	return b
}

func (t *Tragedy) isNotPurchased(ti *Ticket) bool {
	for _, tickets := range t.Play.purchased {
		if tickets.s.number == ti.s.number && tickets.s.row == ti.s.row {
			return false
		}
	}
	return true
}

func (p Play) getName() string {
	return p.name
}

func (p Play) getShowStart() time.Time {
	return p.getShowStart()
}

func (p Play) getShowEnd() time.Time {
	return p.getShowEnd()
}
func (c *Comedy) checkSoldOut() bool {
	shows := len(c.purchased)
	if shows != 25 {
		return false
	}
	return true
}

func (t *Tragedy) checkSoldOut() bool {
	shows := len(t.purchased)
	if shows != 25 {
		return false
	}
	return true
}

func main() {
	//Declared variable
	var show int32
	var rowNum int32
	var seatNum int32
	var isNotSoldOut = true

	//Create shows
	shows := make([]Show, 2)
	shows[0] = &Comedy{0.2, 0.0, Play{"Tartuffe", make([]Ticket, 0), time.Date(2020, time.March, 03, 16, 0, 0, 0, time.UTC), time.Date(2020, time.March, 03, 17, 20, 0, 0, time.UTC)}}
	shows[1] = &Tragedy{0.0, 12.0, Play{"Macbeth", make([]Ticket, 0), time.Date(2020, time.April, 16, 9, 30, 0, 0, time.UTC), time.Date(2020, time.April, 16, 12, 30, 0, 0, time.UTC)}}

	//CreateTheatre
	comedy := Comedy{0.2, 0.0, Play{"Tartuffe", make([]Ticket, 0), time.Date(2020, time.March, 03, 16, 0, 0, 0, time.UTC), time.Date(2020, time.March, 03, 17, 20, 0, 0, time.UTC)}}
	tragedy := Tragedy{0.0, 12.0, Play{"Macbeth", make([]Ticket, 0), time.Date(2020, time.April, 16, 9, 30, 0, 0, time.UTC), time.Date(2020, time.April, 16, 12, 30, 0, 0, time.UTC)}}
	newTheatre := NewTheatre(25, []Show{&comedy, &tragedy})

	//Sell Tickets
	for isNotSoldOut {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your name: ")
		name, _ := reader.ReadString('\n')

		fmt.Print("Choose show (Macbeth or Tartuffe): ")
		mName, _ := reader.ReadString('\n')
		if strings.TrimRight(mName, "\n") == "Tartuffe" {
			show = 0
		} else if strings.TrimRight(mName, "\n") == "Macbeth" {
			show = 1
		}

		fmt.Print("What row # would you like? ")
		for {
			_, err := fmt.Scanf("%d", &rowNum)
			if err != nil {
				fmt.Print(err)
			} else if rowNum < 0 || rowNum > 5 {
				fmt.Println("Enter a valid number between 1 and 5")
			} else {
				break
			}
		}

		fmt.Print("What seat # would you like? ")
		for {
			_, err := fmt.Scanf("%d", &seatNum)
			if err != nil {
				fmt.Print(err)
			} else if seatNum < 0 || seatNum > 5 {
				fmt.Println("Enter a valid number between 1 and 5")
			} else {
				break
			}
		}
		//Make ticket then check if it is valid
		cusTicket := NewTicket(name, &Seat{seatNum, rowNum,
			newTheatre.seats[rowNum].cat}, &shows[show])

		if shows[show].isNotPurchased(cusTicket) {
			shows[show].addPurchase(cusTicket)
			fmt.Println("Ticket has been added! Thank you", name)
		} else {
			if shows[show].checkSoldOut() == true {
				fmt.Println("Oops! Looks like the show is sold out.")
			} else {
				fmt.Println("Error: Seat is taken. Please choose another seat")
			}
		}

		if shows[0].checkSoldOut() == true && shows[1].checkSoldOut() == true {
			fmt.Println("Tickets are sold out for both films. Sorry for the inconvenience")
			isNotSoldOut = false
		}
	}
}
