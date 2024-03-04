package helper

import (
	"encoding/json"
	"net/http"
)

// Json godoc
// @Tags Json
// @Summary Make JSON response
// @Description Parses response value to JSON and writes in ResponseWriter.
// @Produce json
// @Param w body object true "ResponseWriter"
// @Param statusCode body int true "Status code"
// @Param message body interface{} true "Response message"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrResponse
// @Failure 404 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /json [put]
func JSONResponse(w http.ResponseWriter, statusCode int, message interface{}) error {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		return err
	}
	return nil
}
