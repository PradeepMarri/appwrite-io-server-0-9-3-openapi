package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// TeamList represents the TeamList schema from the OpenAPI specification
type TeamList struct {
	Sum int `json:"sum"` // Total sum of items in the list.
	Teams []Team `json:"teams"` // List of teams.
}

// Country represents the Country schema from the OpenAPI specification
type Country struct {
	Name string `json:"name"` // Country name.
	Code string `json:"code"` // Country two-character ISO 3166-1 alpha code.
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Registration int `json:"registration"` // User registration date in Unix timestamp.
	Status int `json:"status"` // User status. 0 for Unactivated, 1 for active and 2 is blocked.
	Id string `json:"$id"` // User ID.
	Email string `json:"email"` // User email address.
	Emailverification bool `json:"emailVerification"` // Email verification status.
	Name string `json:"name"` // User name.
	Passwordupdate int `json:"passwordUpdate"` // Unix timestamp of the most recent password update
	Prefs map[string]interface{} `json:"prefs"` // User preferences as a key-value object
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
	Collection string `json:"$collection"` // Collection ID.
	Id string `json:"$id"` // Document ID.
	Permissions map[string]interface{} `json:"$permissions"` // Document permissions.
}

// Preferences represents the Preferences schema from the OpenAPI specification
type Preferences struct {
}

// ExecutionList represents the ExecutionList schema from the OpenAPI specification
type ExecutionList struct {
	Executions []Execution `json:"executions"` // List of executions.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Provideruid string `json:"providerUid"` // Session Provider User ID.
	Devicemodel string `json:"deviceModel"` // Device model name.
	Clientname string `json:"clientName"` // Client name.
	Clientengineversion string `json:"clientEngineVersion"` // Client engine name.
	Clienttype string `json:"clientType"` // Client type.
	Osname string `json:"osName"` // Operating system name.
	Providertoken string `json:"providerToken"` // Session Provider Token.
	Id string `json:"$id"` // Session ID.
	Clientversion string `json:"clientVersion"` // Client version.
	Expire int `json:"expire"` // Session expiration date in Unix timestamp.
	Clientcode string `json:"clientCode"` // Client code name. View list of [available options](https://github.com/appwrite/appwrite/blob/master/docs/lists/clients.json).
	Devicebrand string `json:"deviceBrand"` // Device brand name.
	Countryname string `json:"countryName"` // Country name.
	Osversion string `json:"osVersion"` // Operating system version.
	Current bool `json:"current"` // Returns true if this the current user session.
	Oscode string `json:"osCode"` // Operating system code name. View list of [available options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
	Userid string `json:"userId"` // User ID.
	Countrycode string `json:"countryCode"` // Country two-character ISO 3166-1 alpha code.
	Ip string `json:"ip"` // IP in use when the session was created.
	Clientengine string `json:"clientEngine"` // Client engine name.
	Provider string `json:"provider"` // Session Provider.
	Devicename string `json:"deviceName"` // Device name.
}

// Execution represents the Execution schema from the OpenAPI specification
type Execution struct {
	Stdout string `json:"stdout"` // The script stdout output string. Logs the last 4,000 characters of the execution stdout output.
	Datecreated int `json:"dateCreated"` // The execution creation date in Unix timestamp.
	Status string `json:"status"` // The status of the function execution. Possible values can be: `waiting`, `processing`, `completed`, or `failed`.
	Stderr string `json:"stderr"` // The script stderr output string. Logs the last 4,000 characters of the execution stderr output
	Id string `json:"$id"` // Execution ID.
	Functionid string `json:"functionId"` // Function ID.
	Time float32 `json:"time"` // The script execution time in seconds.
	Trigger string `json:"trigger"` // The trigger that caused the function to execute. Possible values can be: `http`, `schedule`, or `event`.
	Exitcode int `json:"exitCode"` // The script exit code.
}

// Function represents the Function schema from the OpenAPI specification
type Function struct {
	Datecreated int `json:"dateCreated"` // Function creation date in Unix timestamp.
	Name string `json:"name"` // Function name.
	Schedule string `json:"schedule"` // Function execution schedult in CRON format.
	Vars string `json:"vars"` // Function environment variables.
	Id string `json:"$id"` // Function ID.
	Dateupdated int `json:"dateUpdated"` // Function update date in Unix timestamp.
	Events []string `json:"events"` // Function trigger events.
	Runtime string `json:"runtime"` // Function execution runtime.
	Schedulenext int `json:"scheduleNext"` // Function next scheduled execution date in Unix timestamp.
	Scheduleprevious int `json:"schedulePrevious"` // Function next scheduled execution date in Unix timestamp.
	Tag string `json:"tag"` // Function active tag ID.
	Status string `json:"status"` // Function status. Possible values: disabled, enabled
	Timeout int `json:"timeout"` // Function execution timeout in seconds.
	Permissions map[string]interface{} `json:"$permissions"` // Function permissions.
}

// Membership represents the Membership schema from the OpenAPI specification
type Membership struct {
	Confirm bool `json:"confirm"` // User confirmation status, true if the user has joined the team or false otherwise.
	Email string `json:"email"` // User email address.
	Joined int `json:"joined"` // Date, the user has accepted the invitation to join the team in Unix timestamp.
	Userid string `json:"userId"` // User ID.
	Invited int `json:"invited"` // Date, the user has been invited to join the team in Unix timestamp.
	Name string `json:"name"` // User name.
	Roles []string `json:"roles"` // User list of roles
	Teamid string `json:"teamId"` // Team ID.
	Id string `json:"$id"` // Membership ID.
}

// Permissions represents the Permissions schema from the OpenAPI specification
type Permissions struct {
	Read []string `json:"read"` // Read permissions.
	Write []string `json:"write"` // Write permissions.
}

// Phone represents the Phone schema from the OpenAPI specification
type Phone struct {
	Code string `json:"code"` // Phone code.
	Countrycode string `json:"countryCode"` // Country two-character ISO 3166-1 alpha code.
	Countryname string `json:"countryName"` // Country name.
}

// CollectionList represents the CollectionList schema from the OpenAPI specification
type CollectionList struct {
	Collections []Collection `json:"collections"` // List of collections.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// SessionList represents the SessionList schema from the OpenAPI specification
type SessionList struct {
	Sessions []Session `json:"sessions"` // List of sessions.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// Continent represents the Continent schema from the OpenAPI specification
type Continent struct {
	Code string `json:"code"` // Continent two letter code.
	Name string `json:"name"` // Continent name.
}

// DocumentList represents the DocumentList schema from the OpenAPI specification
type DocumentList struct {
	Sum int `json:"sum"` // Total sum of items in the list.
	Documents []Document `json:"documents"` // List of documents.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code string `json:"code"` // Error code.
	Message string `json:"message"` // Error message.
	Version string `json:"version"` // Server version number.
}

// Rule represents the Rule schema from the OpenAPI specification
type Rule struct {
	Id string `json:"$id"` // Rule ID.
	List []string `json:"list"` // List of allowed values
	Collection string `json:"$collection"` // Rule Collection.
	DefaultField string `json:"default"` // Rule default value.
	Label string `json:"label"` // Rule label.
	TypeField string `json:"type"` // Rule type. Possible values:
	Array bool `json:"array"` // Is array?
	Key string `json:"key"` // Rule key.
	Required bool `json:"required"` // Is required?
}

// TagList represents the TagList schema from the OpenAPI specification
type TagList struct {
	Tags []Tag `json:"tags"` // List of tags.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// File represents the File schema from the OpenAPI specification
type File struct {
	Id string `json:"$id"` // File ID.
	Permissions map[string]interface{} `json:"$permissions"` // File permissions.
	Datecreated int `json:"dateCreated"` // File creation date in Unix timestamp.
	Mimetype string `json:"mimeType"` // File mime type.
	Name string `json:"name"` // File name.
	Signature string `json:"signature"` // File MD5 signature.
	Sizeoriginal int `json:"sizeOriginal"` // File original size in bytes.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Id string `json:"$id"` // Tag ID.
	Command string `json:"command"` // The entrypoint command in use to execute the tag code.
	Datecreated int `json:"dateCreated"` // The tag creation date in Unix timestamp.
	Functionid string `json:"functionId"` // Function ID.
	Size string `json:"size"` // The code size in bytes.
}

// Locale represents the Locale schema from the OpenAPI specification
type Locale struct {
	Continentcode string `json:"continentCode"` // Continent code. A two character continent code "AF" for Africa, "AN" for Antarctica, "AS" for Asia, "EU" for Europe, "NA" for North America, "OC" for Oceania, and "SA" for South America.
	Country string `json:"country"` // Country name. This field support localization.
	Countrycode string `json:"countryCode"` // Country code in [ISO 3166-1](http://en.wikipedia.org/wiki/ISO_3166-1) two-character format
	Currency string `json:"currency"` // Currency code in [ISO 4217-1](http://en.wikipedia.org/wiki/ISO_4217) three-character format
	Eu bool `json:"eu"` // True if country is part of the Europian Union.
	Ip string `json:"ip"` // User IP address.
	Continent string `json:"continent"` // Continent name. This field support localization.
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
	Name string `json:"name"` // Language name.
	Nativename string `json:"nativeName"` // Language native name.
	Code string `json:"code"` // Language two-character ISO 639-1 codes.
}

// CountryList represents the CountryList schema from the OpenAPI specification
type CountryList struct {
	Countries []Country `json:"countries"` // List of countries.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// Currency represents the Currency schema from the OpenAPI specification
type Currency struct {
	Symbolnative string `json:"symbolNative"` // Currency native symbol.
	Code string `json:"code"` // Currency code in [ISO 4217-1](http://en.wikipedia.org/wiki/ISO_4217) three-character format.
	Decimaldigits int `json:"decimalDigits"` // Number of decimal digits.
	Name string `json:"name"` // Currency name.
	Nameplural string `json:"namePlural"` // Currency plural name
	Rounding float32 `json:"rounding"` // Currency digit rounding.
	Symbol string `json:"symbol"` // Currency symbol.
}

// Collection represents the Collection schema from the OpenAPI specification
type Collection struct {
	Id string `json:"$id"` // Collection ID.
	Permissions map[string]interface{} `json:"$permissions"` // Collection permissions.
	Datecreated int `json:"dateCreated"` // Collection creation date in Unix timestamp.
	Dateupdated int `json:"dateUpdated"` // Collection creation date in Unix timestamp.
	Name string `json:"name"` // Collection name.
	Rules []Rule `json:"rules"` // Collection rules.
}

// LogList represents the LogList schema from the OpenAPI specification
type LogList struct {
	Logs []Log `json:"logs"` // List of logs.
}

// FileList represents the FileList schema from the OpenAPI specification
type FileList struct {
	Files []File `json:"files"` // List of files.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// Token represents the Token schema from the OpenAPI specification
type Token struct {
	Id string `json:"$id"` // Token ID.
	Expire int `json:"expire"` // Token expiration date in Unix timestamp.
	Secret string `json:"secret"` // Token secret key. This will return an empty string unless the response is returned using an API key or as part of a webhook payload.
	Userid string `json:"userId"` // User ID.
}

// Log represents the Log schema from the OpenAPI specification
type Log struct {
	Clientengineversion string `json:"clientEngineVersion"` // Client engine name.
	Clientversion string `json:"clientVersion"` // Client version.
	Ip string `json:"ip"` // IP session in use when the session was created.
	Devicemodel string `json:"deviceModel"` // Device model name.
	Clientcode string `json:"clientCode"` // Client code name. View list of [available options](https://github.com/appwrite/appwrite/blob/master/docs/lists/clients.json).
	Event string `json:"event"` // Event name.
	Devicebrand string `json:"deviceBrand"` // Device brand name.
	Countrycode string `json:"countryCode"` // Country two-character ISO 3166-1 alpha code.
	Osname string `json:"osName"` // Operating system name.
	Time int `json:"time"` // Log creation time in Unix timestamp.
	Clienttype string `json:"clientType"` // Client type.
	Countryname string `json:"countryName"` // Country name.
	Clientname string `json:"clientName"` // Client name.
	Devicename string `json:"deviceName"` // Device name.
	Oscode string `json:"osCode"` // Operating system code name. View list of [available options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
	Osversion string `json:"osVersion"` // Operating system version.
	Clientengine string `json:"clientEngine"` // Client engine name.
}

// PhoneList represents the PhoneList schema from the OpenAPI specification
type PhoneList struct {
	Phones []Phone `json:"phones"` // List of phones.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Sum int `json:"sum"` // Total sum of team members.
	Id string `json:"$id"` // Team ID.
	Datecreated int `json:"dateCreated"` // Team creation date in Unix timestamp.
	Name string `json:"name"` // Team name.
}

// UserList represents the UserList schema from the OpenAPI specification
type UserList struct {
	Sum int `json:"sum"` // Total sum of items in the list.
	Users []User `json:"users"` // List of users.
}

// LanguageList represents the LanguageList schema from the OpenAPI specification
type LanguageList struct {
	Languages []Language `json:"languages"` // List of languages.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// CurrencyList represents the CurrencyList schema from the OpenAPI specification
type CurrencyList struct {
	Currencies []Currency `json:"currencies"` // List of currencies.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// ContinentList represents the ContinentList schema from the OpenAPI specification
type ContinentList struct {
	Sum int `json:"sum"` // Total sum of items in the list.
	Continents []Continent `json:"continents"` // List of continents.
}

// FunctionList represents the FunctionList schema from the OpenAPI specification
type FunctionList struct {
	Functions []Function `json:"functions"` // List of functions.
	Sum int `json:"sum"` // Total sum of items in the list.
}

// MembershipList represents the MembershipList schema from the OpenAPI specification
type MembershipList struct {
	Sum int `json:"sum"` // Total sum of items in the list.
	Memberships []Membership `json:"memberships"` // List of memberships.
}
