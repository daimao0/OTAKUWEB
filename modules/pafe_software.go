package modules

type SoftWare struct {
	SoftWareId          int64  `db:"software_id",json:"softWareId"`
	SoftWareTitle       string `db:"software_title",json:"softWareTitle"`
	SoftWareImg         string `db:"software_img",json:"softWareImg"`
	SoftWareUrl         string `db:"software_url",json:"softWareUrl"`
	SoftWareDescription string `db:"software_description",json:"softWareDescription"`
}
