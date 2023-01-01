package predict

// @Summary     predict
// @Description Provide the data of a car and get the cost prediction.
// @Tags        Predict
// @Accept      json
// @Produce     json
// @Param       carData body     predictRequest true "carData"
// @Success     200     {object} transport.Response{Data=int}
// @Failure     400     {object} map[string]string{error=string} "Invalid request"
// @Failure     500     {object} map[string]string{error=string} "Internal server error"
// @Router      /costs/predict [post]
func _() {}
