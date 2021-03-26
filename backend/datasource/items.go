/*
	Author: Kyle Ong
	Date: 10/25/2018

	datasource for readlinglist application
*/
package datasource

import (
	. "ongkyle.com/reading-list/common"
)

var testItems = []Item{
	Item{
		SessionID: "test",
		Title:     "The Man in the High Castle",
		Author:    "Philip K Dick",
		Completed: false},
	Item{
		SessionID: "test",
		Title:     "Inquisitorial Inquiries",
		Author:    "Richard L. Kagan & Abigail Dyer",
		Completed: false},
	Item{
		SessionID: "test",
		Title:     "The Price",
		Author:    "Arthur Miller",
		Completed: false},
	Item{
		SessionID: "test",
		Title:     "A Thousand Splendid Suns",
		Author:    "Khaled Hoesseini",
		Completed: false}}

var Items = map[string][]Item{
	"test": testItems}
