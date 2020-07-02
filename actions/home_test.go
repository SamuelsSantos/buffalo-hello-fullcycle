package actions

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/hello").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Hello Full Cycle")
}
