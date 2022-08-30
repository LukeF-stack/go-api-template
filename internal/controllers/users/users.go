package users

//
//func Register(c *fiber.Ctx) error {
//	database := utils.GetLocal[*gorm.DB](c, "db")
//	var status = 500
//	var e []error
//	var bookModel book.Book
//	err := c.BodyParser(&bookModel)
//	if err != nil {
//		e = append(e, err)
//	}
//	statement := database.Create(&bookModel)
//	if statement.Error != nil {
//		e = append(e, statement.Error)
//	} else {
//		status = 201
//	}
//	c.Status(status)
//	return c.JSON(types.Response{Error: e})
//}

//
//func GetBook(c *fiber.Ctx) error {
//	database := utils.GetLocal[*gorm.DB](c, "db")
//	var status = 500
//	var e []error
//	var data types.Data = nil
//	var bookModel book.Book
//	query := database.First(&bookModel, c.Params("id"))
//	if query.Error != nil {
//		e = append(e, query.Error)
//	} else {
//		status = 200
//	}
//	data = make(map[string]any)
//	data["book"] = bookModel
//	c.Status(status)
//	return c.JSON(types.Response{Error: e, Data: data})
//}
