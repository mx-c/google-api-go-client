// Package civicinfo provides access to the Google Civic Information API.
//
// See https://developers.google.com/civic-information
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/civicinfo/us_v1"
//   ...
//   civicinfoService, err := civicinfo.New(oauthHttpClient)
package civicinfo

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "civicinfo:us_v1"
const apiName = "civicinfo"
const apiVersion = "us_v1"
const basePath = "https://www.googleapis.com/civicinfo/us_v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Elections = NewElectionsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Elections *ElectionsService
}

func NewElectionsService(s *Service) *ElectionsService {
	rs := &ElectionsService{s: s}
	return rs
}

type ElectionsService struct {
	s *Service
}

type AdministrationRegion struct {
	// ElectionAdministrationBody: The election administration body for this
	// area.
	ElectionAdministrationBody *AdministrativeBody `json:"electionAdministrationBody,omitempty"`

	// Id: An ID for this object. IDs may change in future requests and
	// should not be cached. Access to this field requires special access
	// that can be requested from the Request more link on the Quotas page.
	Id string `json:"id,omitempty"`

	// Local_jurisdiction: The city or county that provides election
	// information for this voter. This object can have the same elements as
	// state.
	Local_jurisdiction *AdministrationRegion `json:"local_jurisdiction,omitempty"`

	// Name: The name of the jurisdiction.
	Name string `json:"name,omitempty"`

	// Sources: A list of sources for this area. If multiple sources are
	// listed the data has been aggregated from those sources.
	Sources []*Source `json:"sources,omitempty"`
}

type AdministrativeBody struct {
	// AbsenteeVotingInfoUrl: A URL provided by this administrative body for
	// information on absentee voting.
	AbsenteeVotingInfoUrl string `json:"absenteeVotingInfoUrl,omitempty"`

	// BallotInfoUrl: A URL provided by this administrative body to give
	// contest information to the voter.
	BallotInfoUrl string `json:"ballotInfoUrl,omitempty"`

	// CorrespondenceAddress: The mailing address of this administrative
	// body.
	CorrespondenceAddress *SimpleAddressType `json:"correspondenceAddress,omitempty"`

	// ElectionInfoUrl: A URL provided by this administrative body for
	// looking up general election information.
	ElectionInfoUrl string `json:"electionInfoUrl,omitempty"`

	// ElectionOfficials: The election officials for this election
	// administrative body.
	ElectionOfficials []*ElectionOfficial `json:"electionOfficials,omitempty"`

	// ElectionRegistrationConfirmationUrl: A URL provided by this
	// administrative body for confirming that the voter is registered to
	// vote.
	ElectionRegistrationConfirmationUrl string `json:"electionRegistrationConfirmationUrl,omitempty"`

	// ElectionRegistrationUrl: A URL provided by this administrative body
	// for looking up how to register to vote.
	ElectionRegistrationUrl string `json:"electionRegistrationUrl,omitempty"`

	// ElectionRulesUrl: A URL provided by this administrative body
	// describing election rules to the voter.
	ElectionRulesUrl string `json:"electionRulesUrl,omitempty"`

	// HoursOfOperation: A description of the hours of operation for this
	// administrative body.
	HoursOfOperation string `json:"hoursOfOperation,omitempty"`

	// Name: The name of this election administrative body.
	Name string `json:"name,omitempty"`

	// PhysicalAddress: The physical address of this administrative body.
	PhysicalAddress *SimpleAddressType `json:"physicalAddress,omitempty"`

	// Voter_services: A description of the services this administrative
	// body may provide.
	Voter_services []string `json:"voter_services,omitempty"`

	// VotingLocationFinderUrl: A URL provided by this administrative body
	// for looking up where to vote.
	VotingLocationFinderUrl string `json:"votingLocationFinderUrl,omitempty"`
}

type Candidate struct {
	// CandidateUrl: The URL for the candidate's campaign web site.
	CandidateUrl string `json:"candidateUrl,omitempty"`

	// Channels: A list of known (social) media channels for this candidate.
	Channels []*Channel `json:"channels,omitempty"`

	// Email: The email address for the candidate's campaign.
	Email string `json:"email,omitempty"`

	// Name: The candidate's name.
	Name string `json:"name,omitempty"`

	// OrderOnBallot: The order the candidate appears on the ballot for this
	// contest.
	OrderOnBallot int64 `json:"orderOnBallot,omitempty,string"`

	// Party: The full name of the party the candidate is a member of.
	Party string `json:"party,omitempty"`

	// Phone: The voice phone number for the candidate's campaign office.
	Phone string `json:"phone,omitempty"`

	// PhotoUrl: A URL for a photo of the candidate.
	PhotoUrl string `json:"photoUrl,omitempty"`
}

type Channel struct {
	// Id: The unique public identifier for the candidate's channel.
	Id string `json:"id,omitempty"`

	// Type: The type of channel. The following is a list of types of
	// channels, but is not exhaustive. More channel types may be added at a
	// later time. One of: GooglePlus, YouTube, Facebook, Twitter
	Type string `json:"type,omitempty"`
}

type Contest struct {
	// BallotPlacement: A number specifying the position of this contest on
	// the voter's ballot.
	BallotPlacement int64 `json:"ballotPlacement,omitempty,string"`

	// Candidates: The candidate choices for this contest.
	Candidates []*Candidate `json:"candidates,omitempty"`

	// District: Information about the electoral district that this contest
	// is in.
	District *ElectoralDistrict `json:"district,omitempty"`

	// ElectorateSpecifications: A description of any additional eligibility
	// requirements for voting in this contest.
	ElectorateSpecifications string `json:"electorateSpecifications,omitempty"`

	// Id: An ID for this object. IDs may change in future requests and
	// should not be cached. Access to this field requires special access
	// that can be requested from the Request more link on the Quotas page.
	Id string `json:"id,omitempty"`

	// Level: The level of office for this contest. One of: federal, state,
	// county, city, other
	Level string `json:"level,omitempty"`

	// NumberElected: The number of candidates that will be elected to
	// office in this contest.
	NumberElected int64 `json:"numberElected,omitempty,string"`

	// NumberVotingFor: The number of candidates that a voter may vote for
	// in this contest.
	NumberVotingFor int64 `json:"numberVotingFor,omitempty,string"`

	// Office: The name of the office for this contest.
	Office string `json:"office,omitempty"`

	// PrimaryParty: If this is a partisan election, the name of the party
	// it is for.
	PrimaryParty string `json:"primaryParty,omitempty"`

	// ReferendumSubtitle: A brief description of the referendum. This field
	// is only populated for contests of type 'Referendum'.
	ReferendumSubtitle string `json:"referendumSubtitle,omitempty"`

	// ReferendumTitle: The title of the referendum. (e.g. 'Proposition 42')
	// This field is only populated for contests of type 'Referendum'.
	ReferendumTitle string `json:"referendumTitle,omitempty"`

	// ReferendumUrl: A link the referendum. This field is only populated
	// for contests of type 'Referendum'.
	ReferendumUrl string `json:"referendumUrl,omitempty"`

	// Sources: A list of sources for this contest. If multiple sources are
	// listed, the data has been aggregated from those sources.
	Sources []*Source `json:"sources,omitempty"`

	// Special: "Yes" or "No" depending on whether this a contest being held
	// outside the normal election cycle.
	Special string `json:"special,omitempty"`

	// Type: The type of contest. Usually this will be 'General', 'Primary',
	// or 'Run-off' for contests with candidates. For referenda this will be
	// 'Referendum'.
	Type string `json:"type,omitempty"`
}

type Election struct {
	// ElectionDay: Day of the election in YYYY-MM-DD format.
	ElectionDay string `json:"electionDay,omitempty"`

	// Id: The unique ID of this election.
	Id int64 `json:"id,omitempty,string"`

	// Name: A displayable name for the election.
	Name string `json:"name,omitempty"`
}

type ElectionOfficial struct {
	// EmailAddress: The email address of the election official.
	EmailAddress string `json:"emailAddress,omitempty"`

	// FaxNumber: The fax number of the election official.
	FaxNumber string `json:"faxNumber,omitempty"`

	// Name: The full name of the election official.
	Name string `json:"name,omitempty"`

	// OfficePhoneNumber: The office phone number of the election official.
	OfficePhoneNumber string `json:"officePhoneNumber,omitempty"`

	// Title: The title of the election official.
	Title string `json:"title,omitempty"`
}

type ElectionsQueryResponse struct {
	// Elections: A list of available elections
	Elections []*Election `json:"elections,omitempty"`

	// Kind: The kind, fixed to "civicinfo#electionsQueryResponse".
	Kind string `json:"kind,omitempty"`
}

type ElectoralDistrict struct {
	// Id: An identifier for this district, relative to its scope. For
	// example, the 34th State Senate district would have id "34" and a
	// scope of stateUpper.
	Id string `json:"id,omitempty"`

	// Name: The name of the district.
	Name string `json:"name,omitempty"`

	// Scope: The geographic scope of this district. If unspecified the
	// district's geography is not known. One of: statewide, congressional,
	// stateUpper, stateLower, countywide, judicial, schoolBoard, cityWide,
	// special
	Scope string `json:"scope,omitempty"`
}

type PollingLocation struct {
	// Address: The address of the location
	Address *SimpleAddressType `json:"address,omitempty"`

	// EndDate: The last date that this early vote site may be used. This
	// field is not populated for polling locations.
	EndDate string `json:"endDate,omitempty"`

	// Id: An ID for this object. IDs may change in future requests and
	// should not be cached. Access to this field requires special access
	// that can be requested from the Request more link on the Quotas page.
	Id string `json:"id,omitempty"`

	// Name: The name of the early vote site. This field is not populated
	// for polling locations.
	Name string `json:"name,omitempty"`

	// Notes: Notes about this location (e.g. accessibility ramp or entrance
	// to use)
	Notes string `json:"notes,omitempty"`

	// PollingHours: A description of when this location is open.
	PollingHours string `json:"pollingHours,omitempty"`

	// Sources: A list of sources for this location. If multiple sources are
	// listed the data has been aggregated from those sources.
	Sources []*Source `json:"sources,omitempty"`

	// StartDate: The first date that this early vote site may be used. This
	// field is not populated for polling locations.
	StartDate string `json:"startDate,omitempty"`

	// VoterServices: The services provided by this early vote site. This
	// field is not populated for polling locations.
	VoterServices string `json:"voterServices,omitempty"`
}

type SimpleAddressType struct {
	// City: The city or town for the address.
	City string `json:"city,omitempty"`

	// Line1: The street name and number of this address.
	Line1 string `json:"line1,omitempty"`

	// Line2: The second line the address, if needed.
	Line2 string `json:"line2,omitempty"`

	// Line3: The third line of the address, if needed.
	Line3 string `json:"line3,omitempty"`

	// LocationName: The name of the location.
	LocationName string `json:"locationName,omitempty"`

	// State: The US two letter state abbreviation of the address.
	State string `json:"state,omitempty"`

	// Zip: The US Postal Zip Code of the address.
	Zip string `json:"zip,omitempty"`
}

type Source struct {
	// Name: The name of the data source.
	Name string `json:"name,omitempty"`

	// Official: Whether this data comes from an official government source.
	Official bool `json:"official,omitempty"`
}

type VoterInfoRequest struct {
	// Address: The registered address of the voter to look up.
	Address string `json:"address,omitempty"`
}

type VoterInfoResponse struct {
	// Contests: Contests that will appear on the voter's ballot
	Contests []*Contest `json:"contests,omitempty"`

	// EarlyVoteSites: Locations where the voter is eligible to vote early,
	// prior to election day
	EarlyVoteSites []*PollingLocation `json:"earlyVoteSites,omitempty"`

	// Election: The election that was queried.
	Election *Election `json:"election,omitempty"`

	// Kind: The kind, fixed to "civicinfo#voterInfoResponse".
	Kind string `json:"kind,omitempty"`

	// NormalizedInput: The normalized version of the requested address
	NormalizedInput *SimpleAddressType `json:"normalizedInput,omitempty"`

	// PollingLocations: Locations where the voter is eligible to vote on
	// election day. For states with mail-in voting only, these locations
	// will be nearby drop box locations. Drop box locations are free to the
	// voter and may be used instead of placing the ballot in the mail.
	PollingLocations []*PollingLocation `json:"pollingLocations,omitempty"`

	// State: Local Election Information for the state that the voter votes
	// in. For the US, there will only be one element in this array.
	State []*AdministrationRegion `json:"state,omitempty"`

	// Status: The result of the request. One of: success,
	// noStreetSegmentFound, addressUnparseable, noAddressParameter,
	// multipleStreetSegmentsFound, electionOver, electionUnknown,
	// internalLookupFailure
	Status string `json:"status,omitempty"`
}

// method id "civicinfo.elections.electionQuery":

type ElectionsElectionQueryCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// ElectionQuery: List of available elections to query.
func (r *ElectionsService) ElectionQuery() *ElectionsElectionQueryCall {
	c := &ElectionsElectionQueryCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *ElectionsElectionQueryCall) Do() (*ElectionsQueryResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/civicinfo/us_v1/", "elections")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(ElectionsQueryResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List of available elections to query.",
	//   "httpMethod": "GET",
	//   "id": "civicinfo.elections.electionQuery",
	//   "path": "elections",
	//   "response": {
	//     "$ref": "ElectionsQueryResponse"
	//   }
	// }

}

// method id "civicinfo.elections.voterInfoQuery":

type ElectionsVoterInfoQueryCall struct {
	s                *Service
	electionId       int64
	voterinforequest *VoterInfoRequest
	opt_             map[string]interface{}
}

// VoterInfoQuery: Looks up information relevant to a voter based on the
// voter's registered address.
func (r *ElectionsService) VoterInfoQuery(electionId int64, voterinforequest *VoterInfoRequest) *ElectionsVoterInfoQueryCall {
	c := &ElectionsVoterInfoQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.electionId = electionId
	c.voterinforequest = voterinforequest
	return c
}

// OfficialOnly sets the optional parameter "officialOnly": If set to
// true, only data from official state sources will be returned.
func (c *ElectionsVoterInfoQueryCall) OfficialOnly(officialOnly bool) *ElectionsVoterInfoQueryCall {
	c.opt_["officialOnly"] = officialOnly
	return c
}

func (c *ElectionsVoterInfoQueryCall) Do() (*VoterInfoResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.voterinforequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["officialOnly"]; ok {
		params.Set("officialOnly", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/civicinfo/us_v1/", "voterinfo/{electionId}/lookup")
	urls = strings.Replace(urls, "{electionId}", strconv.FormatInt(c.electionId, 10), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(VoterInfoResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Looks up information relevant to a voter based on the voter's registered address.",
	//   "httpMethod": "POST",
	//   "id": "civicinfo.elections.voterInfoQuery",
	//   "parameterOrder": [
	//     "electionId"
	//   ],
	//   "parameters": {
	//     "electionId": {
	//       "description": "The unique ID of the election to look up. A list of election IDs can be obtained at.https://www.googleapis.com/civicinfo/{version}/elections",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "officialOnly": {
	//       "default": "false",
	//       "description": "If set to true, only data from official state sources will be returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "voterinfo/{electionId}/lookup",
	//   "request": {
	//     "$ref": "VoterInfoRequest"
	//   },
	//   "response": {
	//     "$ref": "VoterInfoResponse"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x2d && r <= 0x7a || r == '~' {
			return r
		}
		return -1
	}, s)
}
