package go-ocr-reporter 

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID                         int         `json:"Id"`
	WorkspaceID                string      `json:"WorkspaceId"`
	CustomerName               string      `json:"CustomerName"`
	CustomerInternalID         string      `json:"CustomerInternalId"` // should be an int ?
	CustomerLabel              string      `json:"CustomerLabel"`
	CustomerStatus             string      `json:"CustomerStatus"`
	Device                     string      `json:"Device"`
	ReferenceID                string      `json:"ReferenceId"`
	DistinguishedName          string      `json:"DistinguishedName"`
	CanonicalName              string      `json:"CanonicalName"`
	AdGroups                   string      `json:"AdGroups"`
	LocalGroups                string      `json:"LocalGroups"`
	Domain                     string      `json:"Domain"`
	DNSSuffix                  string      `json:"DnsSuffix"`
	Username                   string      `json:"Username"`
	FullName                   string      `json:"FullName"`
	Description                string      `json:"Description"`
	EMail                      string      `json:"EMail"`
	Phone                      string      `json:"Phone"`
	Administrator              SpecialBool `json:"Administrator"`
	IsLocalAdministrator       SpecialBool `json:"IsLocalAdministrator"`
	RdsDenied                  string      `json:"RdsDenied"`
	Country                    string      `json:"Country"`
	Source                     string      `json:"Source"`
	AddedInSourceAt            SpecialDate `json:"AddedInSourceAt"`
	UpdatedInSourceAt          SpecialDate `json:"UpdatedInSourceAt"`
	LastLogin                  SpecialDate `json:"LastLogin"`
	LicensingStartDate         string      `json:"LicensingStartDate"`
	LicensingEndDate           string      `json:"LicensingEndDate"`
	AccountExpiresOn           string      `json:"AccountExpiresOn"`
	OtherAttributes            string      `json:"OtherAttributes"`
	ExcludeFromSplaUsageReport SpecialBool `json:"ExcludeFromSplaUsageReport"`
	ReasonForExcluding         string      `json:"ReasonForExcluding"`
	AdditionalRemarks          string      `json:"AdditionalRemarks"`
	ExecutedScript             string      `json:"ExecutedScript"`
	InitStatus                 string      `json:"InitStatus"`
	Notes                      string      `json:"Notes"`
	Status                     string      `json:"Status"`
	DataSource                 string      `json:"DataSource"`
	ScannerID                  string      `json:"ScannerId"`
	Created                    SpecialDate `json:"Created"`
	Updated                    SpecialDate `json:"Updated"`
	RealUpdated                SpecialDate `json:"RealUpdated"`
	Deleted                    interface{} `json:"Deleted"`
}

type TotalCustomers struct {
	Items      []Customer `json:"items"`
	TotalCount int        `json:"total_count"`
}

type Customer struct {
	ID                     int         `json:"id"`
	WsID                   string      `json:"ws_id"`
	CustomerGroupID        interface{} `json:"customer_group_id"`
	GroupLabel             interface{} `json:"group_label"`
	RefID                  interface{} `json:"ref_id"`
	InternalID             string      `json:"internal_id"`
	ExtraField1            string      `json:"extra_field_1"`
	ExtraField2            string      `json:"extra_field_2"`
	Label                  string      `json:"label"`
	OwnUse                 int         `json:"own_use"`
	EnrollmentNumber       string      `json:"enrollment_number"`
	ContactPerson          string      `json:"contact_person"`
	Email                  string      `json:"email"`
	Phone                  string      `json:"phone"`
	PricecategoryID        string      `json:"pricecategory_id"`
	CurrencyID             string      `json:"currency_id"`
	CountryID              string      `json:"country_id"`
	TestingStartDate       SpecialDate `json:"testing_start_date"`
	LicensingStartDate     SpecialDate `json:"licensing_start_date"`
	LicensingEndDate       SpecialDate `json:"licensing_end_date"`
	AfterImport            string      `json:"after_import"`
	Status                 string      `json:"status"`
	DataSource             interface{} `json:"data_source"` // Should be string
	ScannerID              interface{} `json:"scanner_id"`  // Should be string
	Updated                SpecialUnix `json:"updated"`
	RealUpdated            SpecialUnix `json:"real_updated"`
	Deleted                interface{} `json:"deleted"`
	CountUsers             SpecialInt  `json:"count_users"`               // Should be int
	CountServers           SpecialInt  `json:"count_servers"`             // Should be int
	CountDirectScanServers SpecialInt  `json:"count_direct_scan_servers"` // Should be int
}

type Software struct {
	ID                                      int         `json:"Id"`
	WorkspaceID                             string      `json:"WorkspaceId"`
	UserReferenceID                         string      `json:"UserReferenceId"`
	User                                    string      `json:"User"`
	Domain                                  string      `json:"Domain"`
	Username                                string      `json:"Username"`
	EMail                                   string      `json:"EMail"`
	AccountExpiresOn                        interface{} `json:"AccountExpiresOn"`
	SoftwareLicense                         string      `json:"SoftwareLicense"`
	Version                                 string      `json:"Version"`
	SoftwareName                            string      `json:"SoftwareName"`
	PlusEdition                             string      `json:"PlusEdition"`
	SalForSa                                string      `json:"SalForSa"`
	EffectiveNumberOfSalMultiplexingPooling string      `json:"EffectiveNumberOfSalMultiplexingPooling"`
	MailboxSize                             string      `json:"MailboxSize"`
	UsedMailboxSize                         string      `json:"UsedMailboxSize"`
	IsSharedMailbox                         string      `json:"IsSharedMailbox"`
	RemoteAccessGroup                       string      `json:"RemoteAccessGroup"`
	AccessAuthorizationDate                 string      `json:"AccessAuthorizationDate"`
	LicensingStartDate                      string      `json:"LicensingStartDate"`
	ExcludeFromSplaUsageReport              SpecialBool `json:"ExcludeFromSplaUsageReport"`
	ReasonForExcluding                      interface{} `json:"ReasonForExcluding"`
	AdditionalRemarks                       interface{} `json:"AdditionalRemarks"`
	ExecutedScript                          string      `json:"ExecutedScript"`
	ForceSplaLicense                        string      `json:"ForceSplaLicense"`
	AdditionalInformation                   string      `json:"AdditionalInformation"`
	Status                                  string      `json:"Status"`
	DataSource                              string      `json:"DataSource"`
	ScannerID                               string      `json:"ScannerId"`
	Created                                 SpecialDate `json:"Created"`
	Updated                                 SpecialDate `json:"Updated"`
	Deleted                                 interface{} `json:"Deleted"`
}

type OCR struct {
	BaseUrl     string
	WorkspaceID string
	AccessToken string
	Ocid        string
	Client      *http.Client
}

// NewGitlab generates a new gitlab service
func NewOCR(baseUrl, workspaceID, accessToken, ocid string) *OCR {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return &OCR{
		BaseUrl:     baseUrl,
		WorkspaceID: workspaceID,
		AccessToken: accessToken,
		Ocid:        ocid,
		Client:      client,
	}
}

func (o *OCR) GetCustomers() ([]Customer, []byte, error) {
	var body []byte
	var err error
	var c []Customer

	url := "/customer/customers?_format=json&_token=%s&_tenant=%s&_language=en"
	resp, err := o.execRequest("GET", url, body)
	if err != nil {
		fmt.Println(err)
	}

	customersfile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// Use a local file for development
	// customersfile, _ = ioutil.ReadFile("/Users/chris/go/src/github.com/chrisjoyce911/ocr/customers.json")

	var totalCustomers TotalCustomers
	err = json.Unmarshal([]byte(customersfile), &totalCustomers)
	if err != nil {
		fmt.Println(err)
	}
	c = totalCustomers.Items

	return c, body, err
}

func (o *OCR) GetUsers() ([]User, []byte, error) {
	var body []byte
	var err error
	var u []User

	url := "/report/data-export?_format=json&_token=%s&_tenant=%s&_language=en&payload[category]=WsUser"
	resp, err := o.execRequest("GET", url, body)
	if err != nil {
		fmt.Println(err)
	}

	usersfile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// // Use a local file for development
	// usersfile, _ = ioutil.ReadFile("/Users/chris/go/src/github.com/chrisjoyce911/ocr/users.json")
	var users map[string]User
	err = json.Unmarshal([]byte(usersfile), &users)
	if err != nil {
		fmt.Println(err)
	}

	for _, us := range users {
		u = append(u, us)
	}

	return u, body, err
}

func (o *OCR) GetUsersSoftware() ([]Software, []byte, error) {
	var body []byte
	var err error
	var s []Software

	url := "/report/data-export?_format=json&_token=%s&_tenant=%s&_language=en&payload[category]=WsInstalledSoftware-user"
	resp, err := o.execRequest("GET", url, body)
	if err != nil {
		fmt.Println(err)
	}

	usersoftwarefile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Use a local file for development
	// usersoftwarefile, _ = ioutil.ReadFile("/Users/chris/go/src/github.com/chrisjoyce911/ocr/usersoftware.json")
	var software map[string]Software
	err = json.Unmarshal([]byte(usersoftwarefile), &software)
	if err != nil {
		fmt.Println(err)
	}

	for _, us := range software {
		s = append(s, us)
	}

	return s, body, err
}

func (o *OCR) execRequest(method, url string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error

	url = fmt.Sprintf(url, o.AccessToken, o.Ocid)
	url = fmt.Sprintf("%s/api/workspace/%s%s", o.BaseUrl, o.WorkspaceID, url)

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	if err != nil {
		panic("Error while building gitlab request")
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do error: %q", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("*ocr.buildAndExecRequest failed: <%d> %s %s", resp.StatusCode, req.Method, req.URL)
	}

	return resp, err
}
