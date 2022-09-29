package todos

type DataRequest struct {
	Task string `json:"task" binding:"required"`
}
type InputRequest struct {
	ID    string `json:"id" binding:"required"`
	Check string `json:"check" binding:"required"`
}
type InputRequestOnlyId struct {
	ID string `json:"id" binding:"required"`
}
type InputLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
