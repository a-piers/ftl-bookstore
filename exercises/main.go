package exercises

import (
	"fmt"
)

// Book : represent all relevant information related to a book.
type Book struct {
	// Author          Author
	Author          []string
	Title           string
	Description     string
	Copies          int
	Edition         int
	PriceCents      int
	DiscountPercent int
	ISBN11          uint
	ISBN13          uint
	Featured        bool
	IsSerie         bool
	Serie           map[int]string
}

// SoldCopy : lowers the amount of copies with -1.
// Can be used to keep track of amount of copies in store.
func (b Book) SoldCopy() {
	b.Copies--
}

// ChangeTitle : changes the title of the book
func (b Book) ChangeTitle(title string) error {
	if title != "" {
		b.Title = title
	} else {
		return fmt.Errorf("title can't be empty! Provided: %s", title)
	}
	return nil
}

// NetPrice : returns the net price of the book with discount applied.
func (b Book) NetPrice() int {
	return (b.PriceCents) * (1 - (b.DiscountPercent / 100))
}

// Author : represents writers of books
type Author struct {
	Firstname, Lastname string
}

func printTitle(title string) {
	fmt.Println("Title: ", title)
}

// AddToCatalog : adds a book to slice of books and returns the slice
func AddToCatalog(b Book, s []Book) []Book {
	return append(s, b)
}

// Featured : returns slice of books that are Featured
func Featured(s []Book) (f []Book) {
	// var f []Book
	for _, b := range s {
		if b.Featured {
			f = append(f, b)
		} else {
			continue
		}
	}
	return f
}

func main() {
	// var title string
	// var copies int
	// var author string
	// var edition int
	// var inStock bool
	// var specialEdition bool
	// var discountPercentage float64
	// title = "The Happiness Project"
	// copies = 99
	// author = "Paul van Loon"
	// edition = 1
	// inStock = true
	// specialEdition = true
	// discountPercentage = 10.0
	// // printTitle(copies)
	// fmt.Println(title, copies, author, edition, inStock, specialEdition, discountPercentage)

	var b = Book{
		Title: "Nickolas Chuckleby",
		// Author:          Author{Firstname: "Charles", Lastname: "Dickens"},
		Author:          []string{"Charles Dickens"},
		Description:     "Book for childs which is telling nice stories",
		Copies:          99,
		PriceCents:      1199,
		DiscountPercent: 10,
		ISBN11:          10101010101,
		ISBN13:          101010101010101,
		Featured:        true,
		IsSerie:         true,
		Serie: map[int]string{
			1: "Nickolas",
			2: "Chuckleby",
		},
	}
	fmt.Printf(`Book:
	Title: %s
	Author: %s
	Description: %s
	Copies: %d
	Edition: %d
	PriceCents: %d
	NetPrice: %d
	Discount: %d
	ISBN11: %d
`, b.Title, b.Author, b.Description, b.Copies, b.Edition, b.PriceCents, b.NetPrice(), b.DiscountPercent, b.ISBN11)

	var catalog []Book
	catalog = []Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express", Featured: true},
		{Title: "One Hundred Years of Good Company"},
	}
	fmt.Println(catalog[0])
	fmt.Println(catalog[1])
	fmt.Println(len(catalog))
	// fmt.Println(catalog[2]) // Index out of range
	catalog[0] = Book{Title: "Heart of Kindness"}
	// OR
	catalog[0].Title = "Heart of Kindness"
	fmt.Println(catalog[0])
	c := Book{Title: "The Grapes of Mild Irritation", Featured: true}
	catalog = append(catalog, c)
	fmt.Println(catalog)
	catalog = AddToCatalog(Book{Title: "Breaking my heart"}, catalog)
	fmt.Println(catalog)
	for i := range catalog {
		fmt.Println(i, ": ", catalog[i].Title)
	}
	for i, b := range catalog {
		fmt.Printf("Index: %d\tCopies: %d\tAuthor: %s\t\tTitle: %s\n", i, b.Copies, b.Author, b.Title)
		// fmt.Println(i, ": ", b.Title)
	}
	fmt.Printf("Number of book in catalog: %d\n", len(catalog))
	var f []Book = Featured(catalog)
	fmt.Println("Featured books:")
	for i, b := range f {
		fmt.Printf("Index: %d\tCopies: %d\tAuthor: %s\t\tTitle: %s\n", i, b.Copies, b.Author, b.Title)
		// fmt.Println(i, ": ", b.Title)
	}
}
