//Created Sujith P D

package viewmodels

type UserViewModel struct {
	UserName           string
	Password           string
	LoginStatus        string
	LoginType          string
	LoginCreatedOn     string
	LoginCreatedBy     int
	LoginUpdatedOn     string
	LoginUpdatedBy     string
}
type UserDetailsViewModel struct {
	Values             [][]string
	Keys               []string
	UserName           string
	LoginType          string
	LoginStatus        string
	SessionProject     int
	SessionProjectName string
}
type EditUserDetailsViewModel struct {
	Id                 int
	UserName           string
	Uname              string
	Name               string
	Password           string
	UserType           string
	LoginType          string
	SessionProject     int
	SessionProjectName string
}
