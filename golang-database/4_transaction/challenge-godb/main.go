package main

import (
	"enigma_laundy/library"
	"fmt"
	"os"

	"github.com/inancgumus/screen"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	db, _ := library.ConnectDb()
	defer db.Close()

	var menu bool
	for !menu {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("\n\n\n", "==={ ENIGMA LAUNDRY DATABASE }===")

		fmt.Println("Enigma Laundry Menu")
		fmt.Println("1 . Customer")
		fmt.Println("2 . Service")
		fmt.Println("3 . Order")
		fmt.Println("4 . Exit")

		switch library.InputNumber("Masukan pilihan anda :") {
		case 1:
			{

				screen.Clear()
				screen.MoveTopLeft()
				var subMenu bool
				for !subMenu {
					fmt.Println("\n\n\n", "==={ CUSTOMER MENU }===")
					fmt.Println("1 . Create Customer")
					fmt.Println("2 . View Of List Customer")
					fmt.Println("3 . View Details Customer By ID")
					fmt.Println("4 . Update Customer")
					fmt.Println("5 . Delete Customer")
					fmt.Println("6 . Back to Main Menu")

					switch library.InputNumber("Masukan pilihan anda :") {

					case 1:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.CreateCustomer(library.NewCustomer(), db)
						}
						fallthrough
					case 2:
						{
							screen.Clear()
							screen.MoveTopLeft()
							instance := library.ViewOfListCustomer(db)
							rowInstance := []table.Row{}
							for _, v := range instance {

								row := table.Row{v.Customer_id, v.Name, v.Phone, v.Address, v.Created_at.Format("2006-01-02"), v.Update_at.Format("2006-01-02")}
								rowInstance = append(rowInstance, row)
							}

							t := table.NewWriter()
							t.SetOutputMirror(os.Stdout)
							t.AppendHeader(table.Row{"id", "name", "phone", "address", "created_at", "update_at"})
							t.AppendRows(rowInstance)
							t.AppendSeparator()
							//t.AppendFooter(table.Row{})
							t.Render()

						}
					case 3:
						{
							screen.Clear()
							screen.MoveTopLeft()
							v := library.ViewDetailCustomerById(library.InputNumber("\t Masukan id :"), db)
							rowInstance := []table.Row{{v.Customer_id, v.Name, v.Phone, v.Address, v.Created_at.Format("2006-01-02"), v.Update_at.Format("2006-01-02")}}

							t := table.NewWriter()
							t.SetOutputMirror(os.Stdout)
							t.AppendHeader(table.Row{"id", "name", "phone", "address", "created_at", "update_at"})
							t.AppendRows(rowInstance)
							t.AppendSeparator()
							//t.AppendFooter(table.Row{})
							t.Render()

						}
					case 4:

						{
							screen.Clear()
							screen.MoveTopLeft()
							library.UpdateCustomer(library.NewCustomer(), db)
						}
					case 5:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.DeleteCustomer(library.InputText("\t Masukan id :"), db)
						}
					case 6:
						{
							screen.Clear()
							screen.MoveTopLeft()
							subMenu = true
						}
					default:
						{
							screen.Clear()
							screen.MoveTopLeft()
							fmt.Println("*Masukan Pilihan sesuai dengan menu!")
						}
					}
				}
			}
		case 2:
			{
				screen.Clear()
				screen.MoveTopLeft()
				var subMenu bool
				for !subMenu {
					fmt.Println("\n\n\n", "==={ SERVICE MENU }===")
					fmt.Println("1 . Create Service")
					fmt.Println("2 . View Of List Service")
					fmt.Println("3 . View Details Service By ID")
					fmt.Println("4 . Update Service")
					fmt.Println("5 . Delete Service")
					fmt.Println("6 . Back to Main Menu")

					var state bool = false
					switch library.InputNumber("Masukan pilihan anda :") {
					case 1:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.CreateService(library.NewService(), db)
							state = true
						}
						fallthrough
					case 2:
						{
							if state {
								screen.Clear()
								screen.MoveTopLeft()

							}
							instance := library.ViewOfListService(db)
							rowInstance := []table.Row{}
							for _, v := range instance {
								row := table.Row{v.Id, v.Name, v.Unit, v.Price, v.Created_at.Format("2006-01-02"), v.Update_at.Format("2006-01-02")}
								rowInstance = append(rowInstance, row)
							}

							t := table.NewWriter()
							t.SetOutputMirror(os.Stdout)
							t.AppendHeader(table.Row{"id", "name", "unit", "price", "created_at", "update_at"})
							t.AppendRows(rowInstance)
							t.AppendSeparator()
							//t.AppendFooter(table.Row{"TABLE OF ENIGMA SERVICE"})
							t.Render()

						}
					case 3:
						{
							screen.Clear()
							screen.MoveTopLeft()
							instances := library.ViewDetailServiceById(library.InputNumber("Masukan id :"), db)
							t := table.NewWriter()
							t.SetOutputMirror(os.Stdout)
							t.AppendHeader(table.Row{"id", "name", "unit", "price", "created_at", "update_at"})
							t.AppendRow(table.Row{instances.Id, instances.Name, instances.Unit, instances.Price, instances.Created_at.Format("2006-01-02"), instances.Update_at.Format("2006-01-02")})
							t.AppendSeparator()
							//t.AppendFooter(table.Row{"TABLE OF ENIGMA SERVICE"})
							t.Render()
						}
					case 4:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.UpdateService(library.NewService(), db)
						}
					case 5:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.DeleteService(library.InputNumber("Masukan Id :"), db)
						}
					case 6:
						{
							screen.Clear()
							screen.MoveTopLeft()
							subMenu = true
						}
					default:
						{
							screen.Clear()
							screen.MoveTopLeft()
							fmt.Println("*Masukan Pilihan sesuai dengan menu!")
						}
					}
				}

			}
		case 3:
			{
				screen.Clear()
				screen.MoveTopLeft()
				var subMenu bool
				for !subMenu {
					fmt.Println("\n\n\n", "==={ ORDER MENU }===")
					fmt.Println("1 . Create Order")
					fmt.Println("2 . Complete Order")
					fmt.Println("3 . View Of List Order")
					fmt.Println("4 . View Order Details By ID")
					fmt.Println("5 . Back to Main Menu")

					switch library.InputNumber("Masukan pilihan anda :") {
					case 1:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.CreateOrder(db)
						}
					case 2:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.CompleteOrder()
						}
					case 3:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.ViewOfListOrder()
						}
					case 4:
						{
							screen.Clear()
							screen.MoveTopLeft()
							library.ViewOfListOrderById(library.InputNumber("Input Order_Id :"))
						}
					case 5:
						{
							screen.Clear()
							screen.MoveTopLeft()
							subMenu = true
						}
					default:
						{
							screen.Clear()
							screen.MoveTopLeft()
							fmt.Println("*Masukan Pilihan sesuai dengan menu!")
						}
					}
				}

			}
		case 4:
			{
				menu = true
			}
		default:
			{
				screen.Clear()
				screen.MoveTopLeft()
				fmt.Println("Masukan Pilihan sesuai dengan menu")
			}
		}

	}
}
