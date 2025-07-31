package printer

import (
	"bufio"
	"fmt"
	"os"
	"server/models"
	"strings"
	"time"

	"github.com/kenshaw/escpos"
)

const MAX_CHARS_PER_LINE = 32

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func formatFrenchDate(raw string) string {
	t, err := time.Parse("2006-01-02 15:04:05.000000000 -0700 MST", raw)
	if err != nil {
		return "Date invalide"
	}

	months := [...]string{
		"janvier", "fevrier", "mars", "avril", "mai", "juin",
		"juillet", "aout", "septembre", "octobre", "novembre", "decembre",
	}

	day := t.Day()
	month := months[t.Month()-1]
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	return fmt.Sprintf("%d %s %d a %02d:%02d", day, month, year, hour, minute)
}

func printLine(left, right string) string {
  spaceCount := max(MAX_CHARS_PER_LINE-len(left)-len(right), 1)
	spaces := strings.Repeat(" ", spaceCount)
	return left + spaces + right + "\n"
}

func PrintOrder(orderData *models.OrderData) error {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(f)

	p.Init()
	p.SetEmphasize(3)
	p.SetFontSize(1, 1)
	p.SetAlign("center")
	p.Write("Les Delices de Marie")
	p.Formfeed()
	p.SetFontSize(3, 4)
	p.Write("TACO\n")
	p.SetFontSize(1, 1)
	p.Write("Tel: +22891541906 / +22879806420\n")
	p.Write(fmt.Sprintf("%s\n", formatFrenchDate(orderData.IssuedAt)))
	p.Write(fmt.Sprintf("Commande no %s\n", orderData.ID[:8]))
	p.FormfeedN(1)
	p.SetAlign("left")

	for _, item := range orderData.OrderItems {
		total := item.Price * item.Quantity

		p.Write(fmt.Sprintf("%s (%s)\n", item.ProductName, item.ProductVariant))
		productData := printLine(
			fmt.Sprintf("Qte: %d  PU: %d  ", item.Quantity, item.Price),
			fmt.Sprintf("Total: %d", total),
		)
		p.Write(productData)
		p.Write("--------------------------------\n")
	}
	totalLine := printLine("Total", fmt.Sprintf("%d", orderData.Total))
	p.Write(totalLine)

	discountLine := printLine(
		"Remise: ",
		fmt.Sprintf("%d", orderData.Discount),
	)
	p.Write(discountLine)
	p.Write("--------------------------------\n")
	totalLine = printLine("Total a payer", fmt.Sprintf("%d", orderData.SubTotal))
	p.Write(totalLine)

	p.FormfeedN(2)
	p.SetAlign("center")
	p.SetFontSize(2, 3)
	p.Write("MERCI\n")
	p.FormfeedN(3)
	p.End()

	w.Flush()
	return nil
}
