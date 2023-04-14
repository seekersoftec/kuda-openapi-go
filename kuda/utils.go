package kuda

// func (c *Controller) NewUser(ctx *gin.Context) {
// 	var request NewUserRequest

// 	if err := controllers.BindJSON(ctx, &request); err != nil {
// 		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
// 		_ = ctx.Error(appError)
// 		return
// 	}

// 	userModel, err := c.UserService.Create(toUsecaseMapper(&request))
// 	// fmt.Println(userModel)
// 	if err != nil {
// 		_ = ctx.Error(err)
// 		return
// 	}
// 	userResponse := domainToResponseMapper(userModel)
// 	ctx.JSON(http.StatusOK, userResponse)
// }

// BindJSON is a function that binds the request body to the given struct and rewrite the request body on the context
// func BindJSON(c *http.Request, request interface{}) (err error) {
// 	buf := make([]byte, 5120)
// 	num, _ := c.Body.Read(buf)
// 	reqBody := string(buf[0:num])
// 	c.Body = io.NopCloser(bytes.NewBuffer([]byte(reqBody)))
// 	err = c.ShouldBindJSON(request)
// 	c.Body = io.NopCloser(bytes.NewBuffer([]byte(reqBody)))
// 	return
// }

// BindJSONMap is a function that binds the request body to the given map and rewrite the request body on the context
// func BindJSONMap(c *http.Response, response interface{}) (err error) {
// 	buf := make([]byte, 5120)
// 	num, _ := c.Body.Read(buf)
// 	reqBody := buf[0:num]
// 	c.Body = io.NopCloser(bytes.NewBuffer(reqBody))
// 	err = json.Unmarshal(reqBody, &response)
// 	if err != nil {
// 		return err
// 	}
// 	c.Body = io.NopCloser(bytes.NewBuffer(reqBody))
// 	return
// }

//
//
// // BindJSON is a function that binds the request body to the given struct and rewrite the request body on the context
// func BindJSON(c *http.Response, response interface{}) (err error) {
// 	buf := make([]byte, 5120)
// 	num, _ := c.Body.Read(buf)
// 	responseBody := string(buf[0:num])
// 	c.Body = io.NopCloser(bytes.NewBuffer([]byte(responseBody)))
// 	// err = c.ShouldBindJSON(response)
// 	c.Body = io.NopCloser(bytes.NewBuffer([]byte(responseBody)))
// 	return
// }

// func responseToJson(resp *http.Response) (map[string]interface{}, error) {
// 	defer resp.Body.Close()

// 	var data map[string]interface{}
// 	err := json.NewDecoder(resp.Body).Decode(&data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func responseToStruct(resp *http.Response, v interface{}) error {
// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("request failed, status code: %d", resp.StatusCode)
// 		return errors.New("request failed")
// 	}
// 	defer resp.Body.Close()

// 	err := json.NewDecoder(resp.Body).Decode(v)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(resp.Body)

// 	return nil
// }
