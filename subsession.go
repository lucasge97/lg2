package irstats

import (
	"net/http"
	"net/url"
)

type SubSessionResult struct {
	CategoryID            int                `json:"catid"`
	CautionType           int                `json:"cautiontype"`
	CornersPerLap         int                `json:"cornersperlap"`
	DriverChangeParam1    int                `json:"driver_change_param1"`
	DriverChangeParam2    int                `json:"driver_change_param2"`
	DriverChangeRule      int                `json:"driver_change_rule"`
	DriverChanges         int                `json:"driver_changes"`
	EventAverageLap       int                `json:"eventavglap"`
	EventLapsComplete     int                `json:"eventlapscomplete"`
	EventStrengthOfField  int                `json:"eventstrengthoffield"`
	EventType             int                `json:"evttype"`
	LeagueSeasonID        string             `json:"league_season_id"`
	LeagueID              string             `json:"leagueid"`
	LeaveMarbles          int                `json:"leavemarbles"`
	MaxTeamDrivers        int                `json:"max_team_drivers"`
	MaxWeeks              int                `json:"maxweeks"`
	MinTeamDrivers        int                `json:"min_team_drivers"`
	NcautionLaps          int                `json:"ncautionlaps"`
	Ncautions             int                `json:"ncautions"`
	NLapsForQualAVG       int                `json:"nlapsforqualavg"`
	NLapsForSoloAVG       int                `json:"nlapsforsoloavg"`
	NLeadChanges          int                `json:"nleadchanges"`
	PointsType            string             `json:"pointstype"`
	PrivateSessionID      int                `json:"privatesessionid"`
	RaceWeekNum           int                `json:"race_week_num"`
	Drivers               []SubSessionDriver `json:"rows"`
	RservStatus           string             `json:"rserv_status"`
	RubberLevelPractice   int                `json:"rubberlevel_practice"`
	RubberLevelQualify    int                `json:"rubberlevel_qualify"`
	RubberLevelRace       int                `json:"rubberlevel_race"`
	RubberLevelWarmup     int                `json:"rubberlevel_warmup"`
	SeasonName            string             `json:"season_name"`
	SeasonQuarter         int                `json:"season_quarter"`
	SeasonShortname       string             `json:"season_shortname"`
	SeasonYear            int                `json:"season_year"`
	SeasonID              int                `json:"seasonid"`
	SeriesName            string             `json:"series_name"`
	SeriesShortname       string             `json:"series_shortname"`
	SeriesID              int                `json:"seriesid"`
	SessionID             int                `json:"sessionid"`
	SessionName           string             `json:"sessionname"`
	SimSessionType        int                `json:"simsestype"`
	SimulatedStartTime    string             `json:"simulatedstarttime"`
	SpecialEventType      int                `json:"specialeventtype"`
	SpecialEventTypeText  string             `json:"specialeventtypetext"`
	StartTime             string             `json:"start_time"`
	SubsessionID          int                `json:"subsessionid"`
	TimeOfDay             int                `json:"timeofday"`
	TrackConfigName       string             `json:"track_config_name"`
	TrackName             string             `json:"track_name"`
	TrackID               int                `json:"trackid"`
	WeatherFogDensity     int                `json:"weather_fog_density"`
	WeatherRh             int                `json:"weather_rh"`
	WeatherSkies          int                `json:"weather_skies"`
	WeatherTempUnits      int                `json:"weather_temp_units"`
	WeatherTempValue      float64            `json:"weather_temp_value"`
	WeatherType           int                `json:"weather_type"`
	WeatherVarInitial     int                `json:"weather_var_initial"`
	WeatherVarOngoing     int                `json:"weather_var_ongoing"`
	WeatherWindDir        int                `json:"weather_wind_dir"`
	WeatherWindSpeedUnits int                `json:"weather_wind_speed_units"`
	WeatherWindSpeedValue float64            `json:"weather_wind_speed_value"`
}
type SubSessionDriver struct {
	AggChampPoints       int     `json:"aggchamppoints"`
	AvgLap               int     `json:"avglap"`
	BestLapNum           int     `json:"bestlapnum"`
	BestLapTime          int     `json:"bestlaptime"`
	BestNLapsNum         int     `json:"bestnlapsnum"`
	BestNLapsTime        int     `json:"bestnlapstime"`
	BestQualLapAt        int     `json:"bestquallapat"`
	BestQualLapNum       int     `json:"bestquallapnum"`
	BestQualLapTime      int     `json:"bestquallaptime"`
	CarColor1            string  `json:"car_color1"`
	CarColor2            string  `json:"car_color2"`
	CarColor3            string  `json:"car_color3"`
	CarNumberColor1      string  `json:"car_number_color1"`
	CarNumberColor2      string  `json:"car_number_color2"`
	CarNumberColor3      string  `json:"car_number_color3"`
	CarPattern           int     `json:"car_pattern"`
	CarClassID           int     `json:"carclassid"`
	CarID                int     `json:"carid"`
	CarNum               string  `json:"carnum"`
	CarNumberFont        int     `json:"carnumberfont"`
	CarNumberSlant       int     `json:"carnumberslant"`
	CarSponsor1          int     `json:"carsponsor1"`
	CarSponsor2          int     `json:"carsponsor2"`
	CCName               string  `json:"ccName"`
	CCNameShort          string  `json:"ccNameShort"`
	ChampPoints          int     `json:"champpoints"`
	ClassInterval        int     `json:"classinterval"`
	ClubID               int     `json:"clubid"`
	ClubName             string  `json:"clubname"`
	ClubPoints           int     `json:"clubpoints"`
	ClubShortName        string  `json:"clubshortname"`
	CustomerID           int     `json:"custid"`
	DamageModel          int     `json:"damage_model"`
	DisplayName          string  `json:"displayname"`
	Division             int     `json:"division"`
	DivisionName         string  `json:"divisionname"`
	DropRacePoints       int     `json:"dropracepoints"`
	EvtTypeName          string  `json:"evttypename"`
	FinishPos            int     `json:"finishpos"`
	FinishPosInClass     int     `json:"finishposinclass"`
	GripCompoundPractice int     `json:"grip_compound_practice"`
	GripCompoundQualify  int     `json:"grip_compound_qualify"`
	GripCompoundRace     int     `json:"grip_compound_race"`
	GripCompoundWarmup   int     `json:"grip_compound_warmup"`
	GroupID              int     `json:"groupid"`
	HeatInfoID           int     `json:"heatinfoid"`
	HelmColor1           string  `json:"helm_color1"`
	HelmColor2           string  `json:"helm_color2"`
	HelmColor3           string  `json:"helm_color3"`
	HelmPattern          int     `json:"helm_pattern"`
	HostID               string  `json:"hostid"`
	Incidents            int     `json:"incidents"`
	Interval             int     `json:"interval"`
	LapsComplete         int     `json:"lapscomplete"`
	LapsLead             int     `json:"lapslead"`
	LeaguePoints         string  `json:"league_points"`
	LeagueAggPoints      string  `json:"leagueaggpoints"`
	LicenseChangeOval    int     `json:"license_change_oval"`
	LicenseChangeRoad    int     `json:"license_change_road"`
	LicenseCategory      string  `json:"licensecategory"`
	LicenseGroup         int     `json:"licensegroup"`
	MaxPctFuelFill       int     `json:"max_pct_fuel_fill"`
	Multiplier           int     `json:"multiplier"`
	NewCPI               float64 `json:"newcpi"`
	NewIrating           int     `json:"newirating"`
	NewLicenseLevel      int     `json:"newlicenselevel"`
	NewSubLevel          int     `json:"newsublevel"`
	NewTTRating          int     `json:"newttrating"`
	OfficialSession      int     `json:"officialsession"`
	OldCPI               float64 `json:"oldcpi"`
	OldIrating           int     `json:"oldirating"`
	OldLicenseLevel      int     `json:"oldlicenselevel"`
	OldSubLevel          int     `json:"oldsublevel"`
	OldTTRating          int     `json:"oldttrating"`
	OptLapsComplete      int     `json:"optlapscomplete"`
	Pos                  int     `json:"pos"`
	QualLapTime          int     `json:"quallaptime"`
	ReasonOut            string  `json:"reasonout"`
	ReasonOutID          int     `json:"reasonoutid"`
	RestrictResults      string  `json:"restrictresults"`
	SessionStartTime     int64   `json:"sessionstarttime"`
	SimSesName           string  `json:"simsesname"`
	SimSesNum            int     `json:"simsesnum"`
	SimSesTypeName       string  `json:"simsestypename"`
	StartPos             int     `json:"startpos"`
	SubsessionFinishedAt int64   `json:"subsessionfinishedat"`
	SuitColor1           string  `json:"suit_color1"`
	SuitColor2           string  `json:"suit_color2"`
	SuitColor3           string  `json:"suit_color3"`
	SuitPattern          int     `json:"suit_pattern"`
	TrackCategory        string  `json:"track_category"`
	TrackCatID           int     `json:"track_catid"`
	VehicleKeyID         int     `json:"vehiclekeyid"`
	WeightPenaltyKG      int     `json:"weight_penalty_kg"`
	WheelChrome          int     `json:"wheel_chrome"`
	WheelColor           string  `json:"wheel_color"`
}

// SubSessionData returns extensive data about a session.
//
// This endpoint contains data points about a session that is unavailable anywhere else.
func (c *Client) SubSessionData(subSessionID *string, custID *string) (*SubSessionResult, *http.Response, error) {
	v := url.Values{
		"subsessionID": {*subSessionID},
		"custid":       {*custID},
	}
	sessionData := &SubSessionResult{}
	resp, err := c.do(URLPathSubSessionResults, &v, sessionData)
	return sessionData, resp, err
}
