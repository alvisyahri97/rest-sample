package users

type User struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Phone   string `json:"phone" xml:"phone"`
}
