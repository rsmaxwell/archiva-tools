package archivaClient

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"
)

const (
	restServices    = "restServices"
	redbackServices = "redbackServices"

	userService               = "userService"
	createAdminUserEndpoint   = restServices + "/" + redbackServices + "/" + userService + "/createAdminUser"
	createGuestUserEndpoint   = restServices + "/" + redbackServices + "/" + userService + "/createGuestUser"
	createUserEndpoint        = restServices + "/" + redbackServices + "/" + userService + "/createUser"
	deleteUserEndpoint        = restServices + "/" + redbackServices + "/" + userService + "/deleteUser"
	getGuestUserEndpoint      = restServices + "/" + redbackServices + "/" + userService + "/getGuestUser"
	getUserEndpoint           = restServices + "/" + redbackServices + "/" + userService + "/getUser"
	getUsersEndpoint          = restServices + "/" + redbackServices + "/" + userService + "/getUsers"
	updateUserEndpoint        = restServices + "/" + redbackServices + "/" + userService + "/updateUser"
	getUserOperationsEndpoint = restServices + "/" + redbackServices + "/" + userService + "/getUserOperations"

	loginService   = "loginService"
	loginEndpoint  = restServices + "/" + redbackServices + "/" + loginService + "/logIn"
	logoutEndpoint = restServices + "/" + redbackServices + "/" + loginService + "/logout"

	roleManagementService               = "roleManagementService"
	allRolesEndpoint                    = restServices + "/" + redbackServices + "/" + roleManagementService + "/allRoles"
	getApplicationRolesEndpoint         = restServices + "/" + redbackServices + "/" + roleManagementService + "/getApplicationRoles"
	getApplicationsEndpoint             = restServices + "/" + redbackServices + "/" + roleManagementService + "/getApplications"
	getEffectivelyAssignedRolesEndpoint = restServices + "/" + redbackServices + "/" + roleManagementService + "/getEffectivelyAssignedRoles"
	assignRoleByNameEndpoint            = restServices + "/" + redbackServices + "/" + roleManagementService + "/assignRoleByName"
	unassignRoleByNameEndpoint          = restServices + "/" + redbackServices + "/" + roleManagementService + "/unassignRoleByName"
	getRoleEndpoint                     = restServices + "/" + redbackServices + "/" + roleManagementService + "/getRole"

	archivaServices = "archivaServices"

	managedRepositoriesService      = "managedRepositoriesService"
	addManagedRepositoryEndpoint    = restServices + "/" + archivaServices + "/" + managedRepositoriesService + "/addManagedRepository"
	updateManagedRepositoryEndpoint = restServices + "/" + archivaServices + "/" + managedRepositoriesService + "/updateManagedRepository"
	getManagedRepositoriesEndpoint  = restServices + "/" + archivaServices + "/" + managedRepositoriesService + "/getManagedRepositories"
	getManagedRepositoryEndpoint    = restServices + "/" + archivaServices + "/" + managedRepositoriesService + "/getManagedRepository"
	deleteManagedRepositoryEndpoint = restServices + "/" + archivaServices + "/" + managedRepositoriesService + "/deleteManagedRepository"

	browseService        = "browseService"
	versionsListEndpoint = restServices + "/" + archivaServices + "/" + browseService + "/versionsList"

	archivaAdministrationService         = "archivaAdministrationService"
	enabledKnownContentConsumersEndpoint = restServices + "/" + archivaServices + archivaAdministrationService + "/enabledKnownContentConsumers"
	disabledKnownContentConsumerEndpoint = restServices + "/" + archivaServices + archivaAdministrationService + "/disabledKnownContentConsumer"

	repositoriesService    = "repositoriesService"
	projectVersionEndpoint = restServices + "/" + archivaServices + "/" + repositoriesService + "/projectVersion"
)

type ArchivaClient struct {
	scheme              string
	host                string
	port                *int
	base                string
	acceptInsecureCerts bool
}

func (c *ArchivaClient) AcceptInsecureCerts(value bool) {
	c.acceptInsecureCerts = value
}

func (c *ArchivaClient) NewHttpClientBasic() http.Client {

	httpClient := http.Client{Timeout: time.Duration(1) * time.Second}

	if c.acceptInsecureCerts {
		httpClient.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	return httpClient
}

func (c *ArchivaClient) NewHttpClient(session *Session) http.Client {
	httpClient := c.NewHttpClientBasic()
	httpClient.Jar = session.jar
	return httpClient
}

type Session struct {
	token string
	jar   *cookiejar.Jar
}

type AbstractRepository struct {
	Description          string `json:"description,omitempty"`
	Id                   string `json:"id,omitempty"`
	IndexDirectory       string `json:"indexDirectory,omitempty"`
	Layout               string `json:"layout,omitempty"`
	Name                 string `json:"name,omitempty"`
	PackedIndexDirectory string `json:"packedIndexDirectory,omitempty"`
}

type ManagedRepository struct {
	AbstractRepository

	BlockRedeployments      bool   `json:"blockRedeployments,omitempty"`
	CronExpression          string `json:"cronExpression,omitempty"`
	DeleteReleasedSnapshots bool   `json:"deleteReleasedSnapshots,omitempty"`
	Location                string `json:"location,omitempty"`
	Releases                bool   `json:"releases,omitempty"`
	ResetStats              bool   `json:"resetStats,omitempty"`
	RetentionCount          int    `json:"retentionCount,omitempty"`
	RetentionPeriod         int    `json:"retentionPeriod,omitempty"`
	Scanned                 bool   `json:"scanned,omitempty"`
	SkipPackedIndexCreation bool   `json:"skipPackedIndexCreation,omitempty"`
	Snapshots               bool   `json:"snapshots,omitempty"`
	StageRepoNeeded         bool   `json:"stageRepoNeeded,omitempty"`
	StagingRepository       bool   `json:"stagingRepository,omitempty"`
	DaysOlder               int    `json:"daysOlder,omitempty"`
}

type ManagedRepositoryX struct {
	AbstractRepository

	BlockRedeployments      bool   `json:"blockRedeployments,omitempty"`
	CronExpression          string `json:"cronExpression,omitempty"`
	DeleteReleasedSnapshots bool   `json:"deleteReleasedSnapshots,omitempty"`
	Location                string `json:"location,omitempty"`
	Releases                bool   `json:"releases,omitempty"`
	ResetStats              bool   `json:"resetStats,omitempty"`
	RetentionCount          string `json:"retentionCount,omitempty"`
	RetentionPeriod         string `json:"retentionPeriod,omitempty"`
	Scanned                 bool   `json:"scanned,omitempty"`
	SkipPackedIndexCreation bool   `json:"skipPackedIndexCreation,omitempty"`
	Snapshots               bool   `json:"snapshots,omitempty"`
	StageRepoNeeded         bool   `json:"stageRepoNeeded,omitempty"`
	StagingRepository       bool   `json:"stagingRepository,omitempty"`
	DaysOlder               string `json:"daysOlder,omitempty"`
}

func (r *ManagedRepository) ToX() *ManagedRepositoryX {
	return &ManagedRepositoryX{

		AbstractRepository: AbstractRepository{
			Description:          r.Description,
			Id:                   r.Id,
			IndexDirectory:       r.IndexDirectory,
			Layout:               r.Layout,
			Name:                 r.Name,
			PackedIndexDirectory: r.PackedIndexDirectory,
		},

		BlockRedeployments:      r.BlockRedeployments,
		CronExpression:          r.CronExpression,
		DeleteReleasedSnapshots: r.DeleteReleasedSnapshots,
		Location:                r.Location,
		Releases:                r.Releases,
		ResetStats:              r.ResetStats,
		RetentionCount:          strconv.Itoa(r.RetentionCount),
		RetentionPeriod:         strconv.Itoa(r.RetentionPeriod),
		Scanned:                 r.Scanned,
		SkipPackedIndexCreation: r.SkipPackedIndexCreation,
		Snapshots:               r.Snapshots,
		StageRepoNeeded:         r.StageRepoNeeded,
		StagingRepository:       r.StagingRepository,
		DaysOlder:               strconv.Itoa(r.DaysOlder),
	}
}

type User struct {
	Username                    string   `json:"username,omitempty"`
	FullName                    string   `json:"fullName,omitempty"`
	Email                       string   `json:"email,omitempty"`
	Validated                   bool     `json:"validated,omitempty"`
	Locked                      bool     `json:"locked,omitempty"`
	Password                    string   `json:"password,omitempty"`
	PasswordChangeRequired      bool     `json:"passwordChangeRequired,omitempty"`
	Permanent                   bool     `json:"permanent,omitempty"`
	ConfirmPassword             bool     `json:"confirmPassword,omitempty"`
	TimestampAccountCreation    string   `json:"timestampAccountCreation,omitempty"`
	TimestampLastLogin          string   `json:"timestampLastLogin,omitempty"`
	TimestampLastPasswordChange string   `json:"timestampLastPasswordChange,omitempty"`
	PreviousPassword            string   `json:"previousPassword,omitempty"`
	AssignedRoles               []string `json:"assignedRoles,omitempty"`
	ReadOnly                    bool     `json:"readOnly,omitempty"`
	UserManagerId               string   `json:"userManagerId,omitempty"`
	ValidationToken             string   `json:"validationToken,omitempty"`
}

func (u *User) Compare(o *User) bool {

	if u.Username != o.Username {
		return false
	}
	if u.FullName != o.FullName {
		return false
	}
	if u.Email != o.Email {
		return false
	}
	if u.Validated != o.Validated {
		return false
	}
	if u.Locked != o.Locked {
		return false
	}
	// if u.Password != o.Password {
	//	return false
	// }
	if u.PasswordChangeRequired != o.PasswordChangeRequired {
		return false
	}
	if u.Permanent != o.Permanent {
		return false
	}
	// if u.ConfirmPassword != o.ConfirmPassword {
	//	return false
	// }
	// if u.TimestampAccountCreation != o.TimestampAccountCreation {
	// 	return false
	// }
	// if u.TimestampLastLogin != o.TimestampLastLogin {
	//	return false
	//}
	// if u.TimestampLastPasswordChange != o.TimestampLastPasswordChange {
	// 	return false
	//}
	// if u.PreviousPassword != o.PreviousPassword {
	//	return false
	//}
	// if !basic.EqualStringSlices(u.AssignedRoles, o.AssignedRoles) {
	// 	return false
	//}
	if u.ReadOnly != o.ReadOnly {
		return false
	}
	// if u.UserManagerId != o.UserManagerId {
	//	return false
	// }
	// if u.ValidationToken != o.ValidationToken {
	//	return false
	// }

	return true
}

type Operation struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Permanent   bool   `json:"permanent,omitempty"`
}

type Resource struct {
	Identifier string `json:"identifier,omitempty"`
	Pattern    bool   `json:"pattern,omitempty"`
	Permanent  bool   `json:"permanent,omitempty"`
}

type Permission struct {
	Name      string    `json:"name,omitempty"`
	Operation Operation `json:"operation,omitempty"`
	Resource  Resource  `json:"resource,omitempty"`
}

type Role struct {
	Name              string       `json:"name,omitempty"`
	Description       string       `json:"description,omitempty"`
	Assignable        bool         `json:"assignable,omitempty"`
	ChildRoleNames    []string     `json:"childRoleNames,omitempty"`
	Permissions       []Permission `json:"permissions,omitempty"`
	Permanent         bool         `json:"permanent,omitempty"`
	ParentRoleNames   []string     `json:"parentRoleNames,omitempty"`
	ParentsRolesUsers []User       `json:"parentsRolesUsers,omitempty"`
	Users             []User       `json:"users,omitempty"`
	OtherUsers        []User       `json:"otherUsers,omitempty"`
	RemovedUsers      []User       `json:"removedUsers,omitempty"`
}

type RoleTemplate struct {
	Id          string `json:"id,omitempty"`
	NamePrefix  string `json:"namePrefix,omitempty"`
	Delimiter   string `json:"delimiter,omitempty"`
	Description string `json:"description,omitempty"`
	Resource    string `json:"resource,omitempty"`
	Roles       string `json:"roles,omitempty"`
}

type ApplicationRoles struct {
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	GlobalRoles   []string       `json:"globalRoles,omitempty"`
	RoleTemplates []RoleTemplate `json:"roleTemplates,omitempty"`
	Resources     []string       `json:"resources,omitempty"`
}

type Application struct {
	Version         string `json:"version,omitempty"`
	Id              string `json:"id,omitempty"`
	Description     string `json:"description,omitempty"`
	LongDescription string `json:"longDescription,omitempty"`
}

func NewClient(scheme string, host string, port *int, base string) (*ArchivaClient, error) {

	client := &ArchivaClient{
		scheme: scheme,
		host:   host,
		port:   port,
		base:   base,
	}

	return client, nil
}

func (c *ArchivaClient) domainUrl() string {
	portString := ""
	if c.port != nil {
		if (strings.ToLower(c.scheme) == "http") && (*c.port == 80) {
			portString = ""
		} else if (strings.ToLower(c.scheme) == "https") && (*c.port == 443) {
			portString = ""
		} else {
			portString = fmt.Sprintf(":%d", *c.port)
		}
	}

	return fmt.Sprintf("%s://%s%s", c.scheme, c.host, portString)
}

func (c *ArchivaClient) baseUrl() string {
	return c.domainUrl() + "/" + c.base
}

func (r *AbstractRepository) Compare(o *AbstractRepository) bool {

	if r.Description != o.Description {
		return false
	}
	if r.Id != o.Id {
		return false
	}
	if r.IndexDirectory != o.IndexDirectory {
		return false
	}
	if r.Layout != o.Layout {
		return false
	}
	if r.Name != o.Name {
		return false
	}
	if r.PackedIndexDirectory != o.PackedIndexDirectory {
		return false
	}

	return true
}

func (r *ManagedRepository) Compare(o *ManagedRepository) bool {

	r2 := &r.AbstractRepository
	o2 := &o.AbstractRepository

	if !r2.Compare(o2) {
		return false
	}

	if r.BlockRedeployments != o.BlockRedeployments {
		return false
	}

	if r.CronExpression != o.CronExpression {
		return false
	}

	if r.DeleteReleasedSnapshots != o.DeleteReleasedSnapshots {
		return false
	}

	if r.Location != o.Location {
		return false
	}

	if r.Releases != o.Releases {
		return false
	}

	if r.ResetStats != o.ResetStats {
		return false
	}

	if r.RetentionCount != o.RetentionCount {
		return false
	}

	if r.RetentionPeriod != o.RetentionPeriod {
		return false
	}

	if r.Scanned != o.Scanned {
		return false
	}

	if r.SkipPackedIndexCreation != o.SkipPackedIndexCreation {
		return false
	}

	if r.Snapshots != o.Snapshots {
		return false
	}

	if r.StageRepoNeeded != o.StageRepoNeeded {
		return false
	}

	if r.StageRepoNeeded != o.StageRepoNeeded {
		return false
	}

	if r.DaysOlder != o.DaysOlder {
		return false
	}

	return true
}
