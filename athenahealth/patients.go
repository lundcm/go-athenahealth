package athenahealth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Patient represents a patient in athenahealth.
type Patient struct {
	Address1                           string             `json:"address1"`
	Address2                           string             `json:"address2"`
	AgriculturalWorker                 string             `json:"agriculturalworker"`
	AgriculturalWorkerType             string             `json:"agriculturalworkertype"`
	AllPatientStatuses                 []PatientStatus    `json:"allpatientstatuses"`
	AltFirstName                       string             `json:"altfirstname"`
	AssignedSexAtBirth                 string             `json:"assignedsexatbirth"`
	Balances                           []PatientBalance   `json:"balances"`
	CareSummaryDeliveryPreference      string             `json:"caresummarydeliverypreference"`
	City                               string             `json:"city"`
	ConfidentialityCode                string             `json:"confidentialitycode"`
	ConsentToCall                      bool               `json:"consenttocall"`
	ConsentToText                      bool               `json:"consenttotext"`
	ContactHomePhone                   string             `json:"contacthomephone"`
	ContactMobilePhone                 string             `json:"contactmobilephone"`
	ContactName                        string             `json:"contactname"`
	ContactPreference                  string             `json:"contactpreference"`
	ContactPreferenceAnnouncementEmail bool               `json:"contactpreference_announcement_email"`
	ContactPreferenceAnnouncementPhone bool               `json:"contactpreference_announcement_phone"`
	ContactPreferenceAnnouncementSMS   bool               `json:"contactpreference_announcement_sms"`
	ContactPreferenceAppointmentEmail  bool               `json:"contactpreference_appointment_email"`
	ContactPreferenceAppointmentPhone  bool               `json:"contactpreference_appointment_phone"`
	ContactPreferenceAppointmentSMS    bool               `json:"contactpreference_appointment_sms"`
	ContactPreferenceBillingEmail      bool               `json:"contactpreference_billing_email"`
	ContactPreferenceBillingPhone      bool               `json:"contactpreference_billing_phone"`
	ContactPreferenceBillingSMS        bool               `json:"contactpreference_billing_sms"`
	ContactPreferenceLabEmail          bool               `json:"contactpreference_lab_email"`
	ContactPreferenceLabPhone          bool               `json:"contactpreference_lab_phone"`
	ContactPreferenceLabSMS            bool               `json:"contactpreference_lab_sms"`
	ContactRelationship                string             `json:"contactrelationship"`
	CountryCode                        string             `json:"countrycode"`
	CountryCode3166                    string             `json:"countrycode3166"`
	CustomFields                       []CustomFieldValue `json:"customfields"`
	DeceasedDate                       string             `json:"deceaseddate"`
	DefaultPharmacyNCPDPID             string             `json:"defaultpharmacyncpdpid"`
	DepartmentID                       string             `json:"departmentid"`
	DOB                                string             `json:"dob"`
	DoNotCall                          bool               `json:"donotcall"`
	DriversLicense                     bool               `json:"driverslicense"`
	DriversLicenseExpirationDate       string             `json:"driverslicenseexpirationdate"`
	DriversLicenseNumber               string             `json:"driverslicensenumber"`
	DriversLicenseStateID              string             `json:"driverslicensestateid"`
	DriversLicenseURL                  string             `json:"driverslicenseurl"`
	Email                              string             `json:"email"`
	EmailExists                        string             `json:"emailexistsyn"`
	EmployerAddress                    string             `json:"employeraddress"`
	EmployerCity                       string             `json:"employercity"`
	EmployerFax                        string             `json:"employerfax"`
	EmployerID                         string             `json:"employerid"`
	EmployerName                       string             `json:"employername"`
	EmployerPhone                      string             `json:"employerphone"`
	EmployerState                      string             `json:"employerstate"`
	EmployerZip                        string             `json:"employerzip"`
	EthinicityCodes                    []string           `json:"ethnicitycodes"`
	EthnicityCode                      string             `json:"ethnicitycode"`
	FirstAppointment                   string             `json:"firstappointment"`
	FirstName                          string             `json:"firstname"`
	GenderIdentity                     string             `json:"genderidentity"`
	GenderIdentityOther                string             `json:"genderidentityother"`
	GuarantorAddress1                  string             `json:"guarantoraddress1"`
	GuarantorAddress2                  string             `json:"guarantoraddress2"`
	GuarantorAddressSameAsPatient      bool               `json:"guarantoraddresssameaspatient"`
	GuarantorCity                      string             `json:"guarantorcity"`
	GuarantorCountryCode               string             `json:"guarantorcountrycode"`
	GuarantorCountryCode3166           string             `json:"guarantorcountrycode3166"`
	GuarantorDOB                       string             `json:"guarantordob"`
	GuarantorEmail                     string             `json:"guarantoremail"`
	GuarantorEmployerID                string             `json:"guarantoremployerid"`
	GuarantorFirstName                 string             `json:"guarantorfirstname"`
	GuarantorLastName                  string             `json:"guarantorlastname"`
	GuarantorMiddleName                string             `json:"guarantormiddlename"`
	GuarantorPhone                     string             `json:"guarantorphone"`
	GuarantorRelationshipToPatient     string             `json:"guarantorrelationshiptopatient"`
	GuarantorSSN                       string             `json:"guarantorssn"`
	GuarantorState                     string             `json:"guarantorstate"`
	GuarantorSuffix                    string             `json:"guarantorsuffix"`
	GuarantorZip                       string             `json:"guarantorzip"`
	GuardianFirstName                  string             `json:"guardianfirstname"`
	GuardianLastName                   string             `json:"guardianlastname"`
	GuardianMiddleName                 string             `json:"guardianmiddlename"`
	GuardianSuffix                     string             `json:"guardiansuffix"`
	HasMobile                          bool               `json:"hasmobile"`
	Homebound                          bool               `json:"homebound"`
	Homeless                           string             `json:"homeless"`
	HomelessType                       string             `json:"homelesstype"`
	HomePhone                          string             `json:"homephone"`
	IndustryCode                       string             `json:"industrycode"`
	Insurances                         []Insurance        `json:"insurances"`
	Language6392Code                   string             `json:"language6392code"`
	LastAppointment                    string             `json:"lastappointment"`
	LastEmail                          string             `json:"lastemail"`
	LastName                           string             `json:"lastname"`
	LastUpdated                        string             `json:"lastupdated"`
	LastUpdatedBy                      string             `json:"lastupdatedby"`
	LocalPatientID                     string             `json:"localpatientid"`
	MaritalStatus                      string             `json:"maritalstatus"`
	MaritalStatusName                  string             `json:"maritalstatusname"`
	MedicationHistoryConsentVerified   bool               `json:"medicationhistoryconsentverified"`
	MiddleName                         string             `json:"middlename"`
	MobileCarrierID                    string             `json:"mobilecarrierid"`
	MobilePhone                        string             `json:"mobilephone"`
	NextKinName                        string             `json:"nextkinname"`
	NextKinPhone                       string             `json:"nextkinphone"`
	NextKinRelationship                string             `json:"nextkinrelationship"`
	Notes                              string             `json:"notes"`
	OccupationCode                     string             `json:"occupationcode"`
	OnlineStatementOnly                bool               `json:"onlinestatementonly"`
	PatientID                          string             `json:"patientid"`
	PatientPhoto                       bool               `json:"patientphoto"`
	PatientPhotoURL                    string             `json:"patientphotourl"`
	PortalAccessGiven                  bool               `json:"portalaccessgiven"`
	PortalSignatureOnFile              string             `json:"portalsignatureonfile"`
	PortalStatus                       PortalStatus       `json:"portalstatus"`
	PortalTermsOnFile                  bool               `json:"portaltermsonfile"`
	PovertyLevelCalculated             NumberString       `json:"povertylevelcalculated"`
	PovertyLevelFamilySize             string             `json:"povertylevelfamilysize"`
	PovertyLevelFamilySizeDeclined     bool               `json:"povertylevelfamilysizedeclined"`
	PovertyLevelIncomeDeclined         bool               `json:"povertylevelincomedeclined"`
	PovertyLevelIncomePayPeriod        string             `json:"povertylevelincomepayperiod"`
	PovertyLevelIncomePerPayPeriod     string             `json:"povertylevelincomeperpayperiod"`
	PovertyLevelIncomeRangeDeclined    bool               `json:"povertylevelincomerangedeclined"`
	PreferredName                      string             `json:"preferredname"`
	PreferredPronouns                  string             `json:"preferredpronouns"`
	PrimaryDepartmentID                string             `json:"primarydepartmentid"`
	PrimaryProviderID                  string             `json:"primaryproviderid"`
	PrivacyInformationVerified         bool               `json:"privacyinformationverified"`
	PublicHousing                      string             `json:"publichousing"`
	Race                               []string           `json:"race"`
	RaceCode                           string             `json:"racecode"`
	RaceName                           string             `json:"racename"`
	ReferralSourceID                   string             `json:"referralsourceid"`
	ReferralSourceOther                string             `json:"referralsourceother"`
	RegistrationDate                   string             `json:"registrationdate"`
	SchoolBasedHealthCenter            string             `json:"schoolbasedhealthcenter"`
	Sex                                string             `json:"sex"`
	SexualOrientation                  string             `json:"sexualorientation"`
	SexualOrientationOther             string             `json:"sexualorientationother"`
	SMSOptInDate                       string             `json:"smsoptindate"`
	SSN                                string             `json:"ssn"`
	State                              string             `json:"state"`
	Status                             string             `json:"status"`
	Suffix                             string             `json:"suffix"`
	Veteran                            string             `json:"veteran"`
	WorkPhone                          string             `json:"workphone"`
	Zip                                string             `json:"zip"`
}

type PatientStatus struct {
	Status            string       `json:"status"`
	DepartmentID      int          `json:"departmentid"`
	PrimaryProviderID NumberString `json:"primaryproviderid"`
}

type PatientBalance struct {
	Balance         NumberString `json:"balance"`
	DepartmentList  string       `json:"departmentlist"`
	ProviderGroupID int          `json:"providergroupid"`
	CleanBalance    bool         `json:"cleanbalance"`
}

type Insurance struct {
	EligibilityLastChecked              string `json:"eligibilitylastchecked"`
	EligibilityReason                   string `json:"eligibilityreason"`
	EligibilityStatus                   string `json:"eligibilitystatus"`
	ID                                  string `json:"id"`
	InsuranceID                         string `json:"insuranceid"`
	InsuranceIDNumber                   string `json:"insuranceidnumber"`
	InsurancePackageAddress1            string `json:"insurancepackageaddress1"`
	InsurancePackageCity                string `json:"insurancepackagecity"`
	InsurancePackageID                  int    `json:"insurancepackageid"`
	InsurancePackageState               string `json:"insurancepackagestate"`
	InsurancePackageZip                 string `json:"insurancepackagezip"`
	InsurancePhone                      string `json:"insurancephone"`
	InsurancePlanDisplayName            string `json:"insuranceplandisplayname"`
	InsurancePlanName                   string `json:"insuranceplanname"`
	InsurancePolicyHolder               string `json:"insurancepolicyholder"`
	InsurancePolicyHolderAddress1       string `json:"insurancepolicyholderaddress1"`
	InsurancePolicyHolderCity           string `json:"insurancepolicyholdercity"`
	InsurancePolicyHolderCountryCode    string `json:"insurancepolicyholdercountrycode"`
	InsurancePolicyHolderCountryISO3166 string `json:"insurancepolicyholdercountryiso3166"`
	InsurancePolicyHolderDOB            string `json:"insurancepolicyholderdob"`
	InsurancePolicyHolderFirstName      string `json:"insurancepolicyholderfirstname"`
	InsurancePolicyHolderLastName       string `json:"insurancepolicyholderlastname"`
	InsurancePolicyHolderSex            string `json:"insurancepolicyholdersex"`
	InsurancePolicyHolderState          string `json:"insurancepolicyholderstate"`
	InsurancePolicyHolderZip            string `json:"insurancepolicyholderzip"`
	InsuranceType                       string `json:"insurancetype"`
	InsuredAddress                      string `json:"insuredaddress"`
	InsuredCity                         string `json:"insuredcity"`
	InsuredCountryCode                  string `json:"insuredcountrycode"`
	InsuredCountryISO3166               string `json:"insuredcountryiso3166"`
	InsuredDOB                          string `json:"insureddob"`
	InsuredEntityTypeID                 int    `json:"insuredentitytypeid"`
	InsuredFirstName                    string `json:"insuredfirstname"`
	InsuredLastName                     string `json:"insuredlastname"`
	InsuredSex                          string `json:"insuredsex"`
	InsuredState                        string `json:"insuredstate"`
	InsuredZip                          string `json:"insuredzip"`
	IRCName                             string `json:"ircname"`
	RelationshipToInsured               string `json:"relationshiptoinsured"`
	RelationshipToInsuredID             int    `json:"relationshiptoinsuredid"`
	SequenceNumber                      int    `json:"sequencenumber"`
}

type PortalStatus struct {
	BlockedFailedLogins       bool   `json:"blockedfailedlogins"`
	EntityToDisplay           string `json:"entitytodisplay"`
	FamilyBlockedFailedLogins bool   `json:"familyblockedfailedlogins"`
	FamilyRegistered          bool   `json:"familyregistered"`
	NoPortal                  bool   `json:"noportal"`
	PortalRegistrationDate    string `json:"portalregistrationdate"`
	Registered                bool   `json:"registered"`
	Status                    string `json:"status"`
	TermsAccepted             bool   `json:"termsaccepted"`
}

type GetPatientOptions struct {
	ShowCustomFields               bool
	ShowInsurance                  bool
	ShowPortalStatus               bool
	ShowLocalPatientID             bool
	DepartmentID                   int
	LimitLocalPatientID            bool
	ShowAllPatientDepartmentStatus bool
}

// GetPatient - Get data for specific patient, only looks at first returned patient.
//
// GET /v1/{practiceid}/patients/{patientid}
//
// https://docs.athenahealth.com/api/api-ref/patient#Get-specific-patient-record
func (h *HTTPClient) GetPatient(ctx context.Context, id string, opts *GetPatientOptions) (*Patient, error) {
	out, q := []*Patient{}, url.Values{}

	if opts != nil {
		if opts.ShowCustomFields {
			q.Add("showcustomfields", "true")
		}

		if opts.ShowInsurance {
			q.Add("showinsurance", "true")
		}

		if opts.ShowPortalStatus {
			q.Add("showportalstatus", "true")
		}

		if opts.ShowLocalPatientID {
			q.Add("showlocalpatientid", "true")
		}

		if opts.DepartmentID != 0 {
			q.Add("departmentid", fmt.Sprintf("%d", opts.DepartmentID))
		}

		if opts.LimitLocalPatientID {
			q.Add("limitlocalpatientid", "true")
		}

		if opts.ShowAllPatientDepartmentStatus {
			q.Add("showallpatientdepartmentstatus", "true")
		}
	}

	_, err := h.Get(ctx, fmt.Sprintf("/patients/%s", id), q, &out)
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, errors.New("Unexpected length returned")
	}

	return out[0], nil
}

// GetPatients - Get data for specific patient
//
// GET /v1/{practiceid}/patients/{patientid}
//
// https://docs.athenahealth.com/api/api-ref/patient#Get-specific-patient-record
func (h *HTTPClient) GetPatients(ctx context.Context, id string, opts *GetPatientOptions) ([]*Patient, error) {
	out, q := []*Patient{}, url.Values{}

	if opts != nil {
		if opts.ShowCustomFields {
			q.Add("showcustomfields", "true")
		}

		if opts.ShowInsurance {
			q.Add("showinsurance", "true")
		}

		if opts.ShowPortalStatus {
			q.Add("showportalstatus", "true")
		}

		if opts.ShowLocalPatientID {
			q.Add("showlocalpatientid", "true")
		}

		if opts.DepartmentID != 0 {
			q.Add("departmentid", fmt.Sprintf("%d", opts.DepartmentID))
		}

		if opts.LimitLocalPatientID {
			q.Add("limitlocalpatientid", "true")
		}

		if opts.ShowAllPatientDepartmentStatus {
			q.Add("showallpatientdepartmentstatus", "true")
		}
	}

	_, err := h.Get(ctx, fmt.Sprintf("/patients/%s", id), q, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type ListPatientsOptions struct {
	FirstName    string
	LastName     string
	DepartmentID int
	Status       string

	Pagination *PaginationOptions
}

type ListPatientsResult struct {
	Patients []*Patient

	Pagination *PaginationResult
}

type listPatientsResponse struct {
	Patients []*Patient `json:"patients"`

	PaginationResponse
}

// ListPatients - Retrieves a list of patients in the given practice based on search criteria
//
// GET /v1/{practiceid}/patients
//
// https://docs.athenahealth.com/api/api-ref/patient#Get-list-of-patients-for-a-practice
func (h *HTTPClient) ListPatients(ctx context.Context, opts *ListPatientsOptions) (*ListPatientsResult, error) {
	out := &listPatientsResponse{}

	q := url.Values{}

	if opts != nil {
		if len(opts.FirstName) > 0 {
			q.Add("firstname", opts.FirstName)
		}

		if len(opts.LastName) > 0 {
			q.Add("lastname", opts.LastName)
		}

		if opts.DepartmentID != 0 {
			q.Add("departmentid", strconv.Itoa(opts.DepartmentID))
		}

		if len(opts.Status) > 0 {
			q.Add("status", opts.Status)
		}

		if opts.Pagination != nil {
			if opts.Pagination.Limit > 0 {
				q.Add("limit", strconv.Itoa(opts.Pagination.Limit))
			}

			if opts.Pagination.Offset > 0 {
				q.Add("offset", strconv.Itoa(opts.Pagination.Offset))
			}
		}
	}

	_, err := h.Get(ctx, "/patients", q, out)
	if err != nil {
		return nil, err
	}

	return &ListPatientsResult{
		Patients:   out.Patients,
		Pagination: makePaginationResult(out.Next, out.Previous, out.TotalCount),
	}, nil
}

type (
	UpdatePatientOptions struct {
		Address1            *string
		Address2            *string
		AltFirstName        *string
		AssignedSexAtBirth  *string
		City                *string
		ConsentToCall       *bool
		ConsentToText       *bool
		ContactHomePhone    *string
		ContactMobilePhone  *string
		ContactName         *string
		ContactPreference   *string
		ContactRelationship *string
		DepartmentID        *string
		DOB                 *string
		Email               *string
		EthnicityCode       *string
		FirstName           *string
		GenderIdentity      *string
		GenderIdentityOther *string
		HasMobile           *bool
		HomePhone           *string
		Language6392Code    *string
		LastName            *string
		MaritalStatus       *string
		MobilePhone         *string
		Notes               *string
		OccupationCode      *string
		PreferredName       *string
		PreferredPronouns   *string
		PrimaryDepartmentID *string
		Race                []string
		State               *string
		Status              *string
		Zip                 *string
	}

	UpdatePatientResult struct {
		PatientID string
	}

	updatePatientResponse struct {
		PatientID string `json:"patientid"`
	}
)

var (
	UpdatePatientStatusInactiveOption    = "i"
	UpdatePatientStatusActiveOption      = "a"
	UpdatePatientStatusProspectiveOption = "p"
)

// UpdatePatient - Modifies data of a specific patient
//
// PUT /v1/{practiceid}/patients/{patientid}
//
// https://docs.athenahealth.com/api/api-ref/patient#Update-specific-patient-record
func (h *HTTPClient) UpdatePatient(ctx context.Context, patientID string, opts *UpdatePatientOptions) (*UpdatePatientResult, error) {
	out := []*updatePatientResponse{}

	form := url.Values{}

	if opts != nil {
		if opts.Address1 != nil {
			form.Add("address1", *opts.Address1)
		}

		if opts.Address2 != nil {
			form.Add("address2", *opts.Address2)
		}

		if opts.AltFirstName != nil {
			form.Add("altfirstname", *opts.AltFirstName)
		}

		if opts.AssignedSexAtBirth != nil {
			form.Add("assignedsexatbirth", *opts.AssignedSexAtBirth)
		}

		if opts.ConsentToCall != nil {
			form.Add("consenttocall", strconv.FormatBool(*opts.ConsentToCall))
		}

		if opts.ConsentToText != nil {
			form.Add("consenttotext", strconv.FormatBool(*opts.ConsentToText))
		}

		if opts.ContactPreference != nil {
			form.Add("contactpreference", *opts.ContactPreference)
		}

		if opts.ContactName != nil {
			form.Add("contactname", *opts.ContactName)
		}

		if opts.ContactMobilePhone != nil {
			form.Add("contactmobilephone", *opts.ContactMobilePhone)
		}

		if opts.HasMobile != nil {
			form.Add("hasmobileyn", strconv.FormatBool(*opts.HasMobile))
		}

		if opts.ContactHomePhone != nil {
			form.Add("contacthomephone", *opts.ContactHomePhone)
		}

		if opts.ContactRelationship != nil {
			form.Add("contactrelationship", *opts.ContactRelationship)
		}

		if opts.City != nil {
			form.Add("city", *opts.City)
		}

		if opts.DepartmentID != nil {
			form.Add("departmentid", *opts.DepartmentID)
		}

		if opts.DOB != nil {
			form.Add("dob", *opts.DOB)
		}

		if opts.Email != nil {
			form.Add("email", *opts.Email)
		}

		if opts.EthnicityCode != nil {
			form.Add("ethnicitycode", *opts.EthnicityCode)
		}

		if opts.FirstName != nil {
			form.Add("firstname", *opts.FirstName)
		}

		if opts.GenderIdentity != nil {
			form.Add("genderidentity", *opts.GenderIdentity)
		}

		if opts.GenderIdentityOther != nil {
			form.Add("genderidentityother", *opts.GenderIdentityOther)
		}

		if opts.HomePhone != nil {
			form.Add("homephone", *opts.HomePhone)
		}

		if opts.LastName != nil {
			form.Add("lastname", *opts.LastName)
		}

		if opts.Language6392Code != nil {
			form.Add("language6392code", *opts.Language6392Code)
		}

		if opts.MaritalStatus != nil {
			form.Add("maritalstatus", *opts.MaritalStatus)
		}

		if opts.MobilePhone != nil {
			form.Add("mobilephone", *opts.MobilePhone)
		}

		if opts.Notes != nil {
			form.Add("notes", *opts.Notes)
		}

		if opts.OccupationCode != nil {
			form.Add("occupationcode", *opts.OccupationCode)
		}

		if opts.PreferredName != nil {
			form.Add("preferredname", *opts.PreferredName)
		}

		if opts.PreferredPronouns != nil {
			form.Add("preferredpronouns", *opts.PreferredPronouns)
		}

		if opts.PrimaryDepartmentID != nil {
			form.Add("primarydepartmentid", *opts.PrimaryDepartmentID)
		}

		if opts.Race != nil {
			form.Add("race", strings.Join(opts.Race, "\t"))
		}

		if opts.State != nil {
			form.Add("state", *opts.State)
		}

		if opts.Status != nil {
			form.Add("status", *opts.Status)
		}

		if opts.Zip != nil {
			form.Add("zip", *opts.Zip)
		}
	}

	_, err := h.PutForm(ctx, fmt.Sprintf("/patients/%s", patientID), form, &out)
	if err != nil {
		return nil, err
	}

	if len(out) != 1 {
		return nil, errors.New("unexpected response")
	}

	return &UpdatePatientResult{
		PatientID: out[0].PatientID,
	}, nil
}

type GetPatientPhotoOptions struct {
	JPEGOutput bool
}

type patientPhoto struct {
	Image string `json:"image"`
}

// GetPatientPhoto - Get a patient's photo
//
// GET /v1/{practiceid}/patients/{patientid}/photo
//
// https://docs.athenahealth.com/api/api-ref/patient-photo#Get-patient's-photo
func (h *HTTPClient) GetPatientPhoto(ctx context.Context, patientID string, opts *GetPatientPhotoOptions) (string, error) {
	out := &patientPhoto{}

	q := url.Values{}

	if opts != nil {
		if opts.JPEGOutput {
			return "", errors.New("JPEGOutput is not supported")
		}
	}

	_, err := h.Get(ctx, fmt.Sprintf("/patients/%s/photo", patientID), q, &out)
	if err != nil {
		return "", err
	}

	return out.Image, nil
}

// UpdatePatientPhoto - Update a patient's photo.
//
// POST /v1/{practiceid}/patients/{patientid}/photo
//
// https://docs.athenahealth.com/api/api-ref/patient-photo#Update-patient's-photo
func (h *HTTPClient) UpdatePatientPhoto(ctx context.Context, patientID string, data []byte) error {
	form := url.Values{}
	form.Add("image", base64.StdEncoding.EncodeToString(data))

	_, err := h.PostForm(ctx, fmt.Sprintf("/patients/%s/photo", patientID), form, nil)
	return err
}

// UpdatePatientPhotoReader - performs the same operation as UpdatePatientPhoto except is more memory efficient
// by streaming the image contents into the request, assuming you haven't already read the
// entire image contents into memory
// POST /v1/{practiceid}/patients/{patientid}/photo
// https://developer.athenahealth.com/docs/read/forms_and_documents/Patient_Photo#section-1
func (h *HTTPClient) UpdatePatientPhotoReader(ctx context.Context, patientID string, r io.Reader) error {
	form := NewFormURLEncoder()
	form.AddReader("image", r)

	_, err := h.PostFormReader(ctx, fmt.Sprintf("/patients/%s/photo", patientID), form, nil)
	return err
}

type ListChangedPatientOptions struct {
	DepartmentID               string
	IgnoreRestrictions         bool
	LeaveUnprocessed           bool
	PatientID                  string
	ReturnGlobalID             bool
	ShowProcessedEndDatetime   time.Time
	ShowProcessedStartDatetime time.Time
}

type listChangedPatientsResponse struct {
	ChangedPatients []*Patient `json:"patients"`
}

// ListChangedPatients - Gets list of changes made to the patient record
//
// GET /v1/{practiceid}/patients/changed
//
// https://docs.athenahealth.com/api/api-ref/patient#Get-list-of-changes-in-patient-records
func (h *HTTPClient) ListChangedPatients(ctx context.Context, opts *ListChangedPatientOptions) ([]*Patient, error) {
	out := &listChangedPatientsResponse{}

	q := url.Values{}

	if opts != nil {
		if len(opts.DepartmentID) > 0 {
			q.Add("departmentid", opts.DepartmentID)
		}

		if opts.IgnoreRestrictions {
			q.Add("ignorerestrictions", strconv.FormatBool(opts.IgnoreRestrictions))
		}

		if opts.LeaveUnprocessed {
			q.Add("leaveunprocessed", strconv.FormatBool(opts.LeaveUnprocessed))
		}

		if len(opts.PatientID) > 0 {
			q.Add("patientid", opts.PatientID)
		}

		if opts.ReturnGlobalID {
			q.Add("returnglobalid", strconv.FormatBool(opts.ReturnGlobalID))
		}

		if !opts.ShowProcessedEndDatetime.IsZero() {
			q.Add("showprocessedenddatetime", opts.ShowProcessedEndDatetime.Format("01/02/2006 15:04:05"))
		}

		if !opts.ShowProcessedStartDatetime.IsZero() {
			q.Add("showprocessedstartdatetime", opts.ShowProcessedStartDatetime.Format("01/02/2006 15:04:05"))
		}
	}

	_, err := h.Get(ctx, "/patients/changed", q, out)
	if err != nil {
		return nil, err
	}

	return out.ChangedPatients, nil
}

type UpdatePatientInformationVerificationDetailsOptions struct {
	DepartmentID                int
	ExpirationDate              *time.Time
	InsuredSignature            *string
	PatientSignature            *string
	PrivacyNotice               *string
	ReasonPatientUnableToSign   *string
	SignatureDatetime           time.Time
	SignatureName               string
	SignerRelationshipToPatient *string
}

type updatePatientInformationVerificationDetailsResponse struct {
	Success bool `json:"success"`
}

// UpdatePatientInformationVerificationDetails - Update a patient's verified privacy information
//
// POST /v1/{practiceid}/patients/{patientid}/privacyinformationverified
//
// https://docs.athenahealth.com/api/api-ref/privacy-information-verification#Update-patient's-privacy-information-verification-details
func (h *HTTPClient) UpdatePatientInformationVerificationDetails(ctx context.Context, patientID string, opts *UpdatePatientInformationVerificationDetailsOptions) error {
	out := []*updatePatientInformationVerificationDetailsResponse{}
	var form url.Values

	if opts != nil {
		form = url.Values{}

		form.Add("departmentid", strconv.Itoa(opts.DepartmentID))

		if opts.ExpirationDate != nil {
			form.Add("expirationdate", opts.ExpirationDate.Format("01/02/2006"))
		}

		if opts.InsuredSignature != nil {
			form.Add("insuredsignature", *opts.InsuredSignature)
		}

		if opts.PatientSignature != nil {
			form.Add("patientsignature", *opts.PatientSignature)
		}

		if opts.PrivacyNotice != nil {
			form.Add("privacynotice", *opts.PrivacyNotice)
		}

		if opts.ReasonPatientUnableToSign != nil {
			form.Add("reasonpatientunabletosign", *opts.ReasonPatientUnableToSign)
		}

		form.Add("signaturedatetime", opts.SignatureDatetime.Format("01/02/2006 15:04:05"))
		form.Add("signaturename", opts.SignatureName)

		if opts.SignerRelationshipToPatient != nil {
			form.Add("signerrelationshiptopatientid", *opts.SignerRelationshipToPatient)
		}
	}

	_, err := h.PostForm(ctx, fmt.Sprintf("/patients/%s/privacyinformationverified", patientID), form, &out)
	if err != nil {
		return err
	}

	if len(out) != 1 || !out[0].Success {
		return errors.New("unexpected response")
	}

	return nil
}

type UpdatePatientMedicationHistoryConsentOptions struct {
	DepartmentID      int
	SignatureDatetime time.Time
	SignatureName     string
}

type updatePatientMedicationHistoryConsentResponse struct {
	Success string `json:"success"`
}

// UpdatePatientMedicationHistoryConsent - Update patient's medication history consent flag as having been verified
//
// POST /v1/{practiceid}/patients/{patientid}/medicationhistoryconsentverified
//
// https://docs.athenahealth.com/api/api-ref/medication-history-consent
func (h *HTTPClient) UpdatePatientMedicationHistoryConsent(ctx context.Context, patientID string, opts *UpdatePatientMedicationHistoryConsentOptions) error {
	out := []*updatePatientMedicationHistoryConsentResponse{}
	var form url.Values

	if opts != nil {
		form = url.Values{}

		form.Add("departmentid", strconv.Itoa(opts.DepartmentID))

		form.Add("signaturedatetime", opts.SignatureDatetime.Format("01/02/2006 15:04:05"))
		form.Add("signaturename", opts.SignatureName)
	}

	_, err := h.PostForm(ctx, fmt.Sprintf("/patients/%s/medicationhistoryconsentverified", patientID), form, &out)
	if err != nil {
		return err
	}

	if len(out) != 1 || out[0].Success == "false" {
		return errors.New("unexpected response")
	}

	return nil
}

// GetPatientCustomFields - Get custom fields information for a specific patient
//
// GET /v1/{practiceid}/patients/{patientid}/customfields
//
// https://docs.athenahealth.com/api/api-ref/patient-custom-fields#Get-custom-field-information-from-patient's-records
func (h *HTTPClient) GetPatientCustomFields(ctx context.Context, patientID, departmentID string) ([]*CustomFieldValue, error) {
	out := []*CustomFieldValue{}

	query := url.Values{}
	query.Add("departmentid", departmentID)

	_, err := h.Get(ctx, fmt.Sprintf("/patients/%s/customfields", patientID), query, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type updatePatientCustomFieldsResponse struct {
	Success         bool `json:"success"`
	UpdatedCount    int  `json:"updatedCount"`
	DisallowedCount int  `json:"disallowedCount"`
}

// UpdatePatientCustomFields - Update custom fields data for a specific patient
//
// PUT /v1/{practiceid}/patients/{patientid}/customfields
//
// https://docs.athenahealth.com/api/api-ref/patient-custom-fields#Update-custom-field-information-from-patient's-records
func (h *HTTPClient) UpdatePatientCustomFields(ctx context.Context, patientID, departmentID string, customFields []*CustomFieldValue) error {
	out := &updatePatientCustomFieldsResponse{}

	customFieldsJSON, err := json.Marshal(customFields)
	if err != nil {
		return errors.Wrap(err, "marshaling custom fields")
	}

	form := url.Values{}
	form.Add("departmentid", departmentID)
	form.Add("customfields", string(customFieldsJSON))

	_, err = h.PutForm(ctx, fmt.Sprintf("/patients/%s/customfields", patientID), form, out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New("unexpected response")
	}

	return nil
}

type ListPatientsMatchingCustomFieldOptions struct {
	CustomFieldID    string
	CustomFieldValue string

	Pagination *PaginationOptions
}

type ListPatientsMatchingCustomFieldResult struct {
	Patients []*Patient

	Pagination *PaginationResult
}

type listPatientsMatchingCustomFieldResponse struct {
	Patients []*Patient `json:"patients"`

	PaginationResponse
}

// ListPatientsMatchingCustomField - Get matching patients based on custom field value
//
// GET /v1/{practiceid}/patients/customfields/{customfieldid}/{customfieldvalue}
//
// https://docs.athenahealth.com/api/api-ref/patient#Get-list-of-patients---matching-custom-field-criteria
func (h *HTTPClient) ListPatientsMatchingCustomField(ctx context.Context, opts *ListPatientsMatchingCustomFieldOptions) (*ListPatientsMatchingCustomFieldResult, error) {
	if opts == nil {
		panic("opts is nil")
	}

	out := &listPatientsMatchingCustomFieldResponse{}
	q := url.Values{}

	if opts.Pagination != nil {
		if opts.Pagination.Limit > 0 {
			q.Add("limit", strconv.Itoa(opts.Pagination.Limit))
		}

		if opts.Pagination.Offset > 0 {
			q.Add("offset", strconv.Itoa(opts.Pagination.Offset))
		}
	}

	_, err := h.Get(ctx, fmt.Sprintf("/patients/customfields/%s/%s", opts.CustomFieldID, opts.CustomFieldValue), q, out)
	if err != nil {
		return nil, err
	}

	return &ListPatientsMatchingCustomFieldResult{
		Patients:   out.Patients,
		Pagination: makePaginationResult(out.Next, out.Previous, out.TotalCount),
	}, nil
}

type CreatePatientOptions struct {
	Address1              string
	Address2              string
	City                  string
	DepartmentID          string
	DOB                   time.Time
	Email                 string
	FirstName             string
	HomePhone             string
	LastName              string
	MiddleName            string
	MobilePhone           string
	Notes                 string
	Sex                   string
	SSN                   string
	State                 string
	Status                string
	Zip                   string
	BypassPatientMatching bool
}

type createPatientResponse struct {
	ErrorMessage string `json:"errormessage"`
	PatientID    string `json:"patientid"`
}

// CreatePatient - Create new patient record
//
// POST /v1/{practiceid}/patients
//
// https://docs.athenahealth.com/api/api-ref/patient#Create-new-patient-record
func (h *HTTPClient) CreatePatient(ctx context.Context, opts *CreatePatientOptions) (string, error) {
	if opts == nil {
		panic("opts is nil")
	}

	out := []*createPatientResponse{}

	form := url.Values{}

	form.Add("address1", opts.Address1)
	form.Add("address2", opts.Address2)
	form.Add("city", opts.City)
	form.Add("departmentid", opts.DepartmentID)
	form.Add("dob", opts.DOB.Format("01/02/2006"))
	form.Add("email", opts.Email)
	form.Add("firstname", opts.FirstName)
	form.Add("homephone", opts.HomePhone)
	form.Add("lastname", opts.LastName)
	form.Add("middlename", opts.MiddleName)
	form.Add("mobilephone", opts.MobilePhone)
	form.Add("notes", opts.Notes)
	form.Add("sex", opts.Sex)
	form.Add("ssn", opts.SSN)
	form.Add("state", opts.State)
	form.Add("status", opts.Status)
	form.Add("zip", opts.Zip)

	if opts.BypassPatientMatching {
		form.Add("bypasspatientmatching", "true")
	}

	_, err := h.PostForm(ctx, "/patients", form, &out)
	if err != nil {
		return "", err
	}

	if len(out) != 1 {
		return "", errors.New("unexpected response")
	}

	if len(out[0].ErrorMessage) > 0 {
		return "", errors.New(out[0].ErrorMessage)
	}

	return out[0].PatientID, nil
}
