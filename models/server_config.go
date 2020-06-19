package models

import "encoding/xml"

type ServerConfig struct {
	XMLName   xml.Name `xml:"config" json:"-" gqlgen:"xmlName"`
	Text      string   `xml:",chardata" json:"-" gqlgen:"text"`
	Speed     float64  `xml:"speed" json:"speed" gqlgen:"speed"`
	UnitSpeed float64  `xml:"unit_speed" json:"unit_speed" gqlgen:"unitSpeed"`
	Moral     float64  `xml:"moral" json:"moral" gqlgen:"moral"`
	Build     struct {
		Text    string  `xml:",chardata" json:"-" gqlgen:"text"`
		Destroy float64 `xml:"destroy" json:"destroy" gqlgen:"destroy"`
	} `xml:"build" json:"build" gqlgen:"build"`
	Misc struct {
		Text            string `xml:",chardata" json:"-" gqlgen:"text"`
		KillRanking     int    `xml:"kill_ranking" json:"kill_ranking" gqlgen:"killRanking"`
		Tutorial        int    `xml:"tutorial" json:"tutorial" gqlgen:"tutorial"`
		TradeCancelTime int    `xml:"trade_cancel_time" json:"trade_cancel_time" gqlgen:"tradeCancelTime"`
	} `xml:"misc" json:"misc" gqlgen:"misc"`
	Commands struct {
		Text              string `xml:",chardata" json:"-" gqlgen:"text"`
		MillisArrival     int    `xml:"millis_arrival" json:"millis_arrival" gqlgen:"millisArrival"`
		CommandCancelTime int    `xml:"command_cancel_time" json:"command_cancel_time" gqlgen:"commandCancelTime"`
	} `xml:"commands" json:"commands" gqlgen:"commands"`
	Newbie struct {
		Text                 string `xml:",chardata" json:"-" gqlgen:"text"`
		Days                 int    `xml:"days" json:"days" gqlgen:"days"`
		RatioDays            int    `xml:"ratio_days" json:"ratio_days" gqlgen:"ratioDays"`
		Ratio                int    `xml:"ratio" json:"ratio" gqlgen:"ratio"`
		RemoveNewbieVillages int    `xml:"removeNewbieVillages" json:"remove_newbie_villages" gqlgen:"removeNewbieVillages"`
	} `xml:"newbie" json:"newbie" gqlgen:"newbie"`
	Game struct {
		Text               string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildtimeFormula   int     `xml:"buildtime_formula" json:"buildtime_formula" gqlgen:"buildtimeFormula"`
		Knight             int     `xml:"knight" json:"knight" gqlgen:"knight"`
		KnightNewItems     bool    `xml:"knight_new_items" json:"knight_new_items" gqlgen:"knightNewItems"`
		Archer             int     `xml:"archer" json:"archer" gqlgen:"archer"`
		Tech               int     `xml:"tech" json:"tech" gqlgen:"tech"`
		FarmLimit          int     `xml:"farm_limit" json:"farm_limit" gqlgen:"farmLimit"`
		Church             int     `xml:"church" json:"church" gqlgen:"church"`
		Watchtower         int     `xml:"watchtower" json:"watchtower" gqlgen:"watchtower"`
		Stronghold         int     `xml:"stronghold" json:"stronghold" gqlgen:"stronghold"`
		FakeLimit          int     `xml:"fake_limit" json:"fake_limit" gqlgen:"fakeLimit"`
		BarbarianRise      float64 `xml:"barbarian_rise" json:"barbarian_rise" gqlgen:"barbarianRise"`
		BarbarianShrink    int     `xml:"barbarian_shrink" json:"barbarian_shrink" gqlgen:"barbarianShrink"`
		BarbarianMaxPoints int     `xml:"barbarian_max_points" json:"barbarian_max_points" gqlgen:"barbarianMaxPoints"`
		Hauls              int     `xml:"hauls" json:"hauls" gqlgen:"hauls"`
		HaulsBase          int     `xml:"hauls_base" json:"hauls_base" gqlgen:"haulsBase"`
		HaulsMax           int     `xml:"hauls_max" json:"hauls_max" gqlgen:"haulsMax"`
		BaseProduction     int     `xml:"base_production" json:"base_production" gqlgen:"baseProduction"`
		Event              int     `xml:"event" json:"event" gqlgen:"event"`
		SuppressEvents     int     `xml:"suppress_events" json:"suppress_events" gqlgen:"suppressEvents"`
	} `xml:"game" json:"game" gqlgen:"game"`
	Buildings struct {
		Text             string `xml:",chardata" json:"-" gqlgen:"text"`
		CustomMain       int    `xml:"custom_main" json:"custom_main" gqlgen:"customMain"`
		CustomFarm       int    `xml:"custom_farm" json:"custom_farm" gqlgen:"customFarm"`
		CustomStorage    int    `xml:"custom_storage" json:"custom_storage" gqlgen:"customStorage"`
		CustomPlace      int    `xml:"custom_place" json:"custom_place" gqlgen:"customPlace"`
		CustomBarracks   int    `xml:"custom_barracks" json:"custom_barracks" gqlgen:"customBarracks"`
		CustomChurch     int    `xml:"custom_church" json:"custom_church" gqlgen:"customChurch"`
		CustomSmith      int    `xml:"custom_smith" json:"custom_smith" gqlgen:"customSmith"`
		CustomWood       int    `xml:"custom_wood" json:"custom_wood" gqlgen:"customWood"`
		CustomStone      int    `xml:"custom_stone" json:"custom_stone" gqlgen:"customStone"`
		CustomIron       int    `xml:"custom_iron" json:"custom_iron" gqlgen:"customIron"`
		CustomMarket     int    `xml:"custom_market" json:"custom_market" gqlgen:"customMarket"`
		CustomStable     int    `xml:"custom_stable" json:"custom_stable" gqlgen:"customStable"`
		CustomWall       int    `xml:"custom_wall" json:"custom_wall" gqlgen:"customWall"`
		CustomGarage     int    `xml:"custom_garage" json:"custom_garage" gqlgen:"customGarage"`
		CustomHide       int    `xml:"custom_hide" json:"custom_hide" gqlgen:"customHide"`
		CustomSnob       int    `xml:"custom_snob" json:"custom_snob" gqlgen:"customSnob"`
		CustomStatue     int    `xml:"custom_statue" json:"custom_statue" gqlgen:"customStatue"`
		CustomWatchtower int    `xml:"custom_watchtower" json:"custom_watchtower" gqlgen:"customWatchtower"`
	} `xml:"buildings" json:"buildings" gqlgen:"buildings"`
	Snob struct {
		Text          string  `xml:",chardata" json:"-" gqlgen:"text"`
		Gold          int     `xml:"gold" json:"gold" gqlgen:"gold"`
		CheapRebuild  int     `xml:"cheap_rebuild" json:"cheap_rebuild" gqlgen:"cheapRebuild"`
		Rise          int     `xml:"rise" json:"rise" gqlgen:"rise"`
		MaxDist       int     `xml:"max_dist" json:"max_dist" gqlgen:"maxDist"`
		Factor        float64 `xml:"factor" json:"factor" gqlgen:"factor"`
		CoinWood      int     `xml:"coin_wood" json:"coin_wood" gqlgen:"coinWood"`
		CoinStone     int     `xml:"coin_stone" json:"coin_stone" gqlgen:"coinStone"`
		CoinIron      int     `xml:"coin_iron" json:"coin_iron" gqlgen:"coinIron"`
		NoBarbConquer bool    `xml:"no_barb_conquer" json:"no_barb_conquer" gqlgen:"noBarbConquer"`
	} `xml:"snob" json:"snob" gqlgen:"snob"`
	Ally struct {
		Text                  string `xml:",chardata" json:"-" gqlgen:"text"`
		NoHarm                int    `xml:"no_harm" json:"no_harm" gqlgen:"noHarm"`
		NoOtherSupport        int    `xml:"no_other_support" json:"no_other_support" gqlgen:"noOtherSupport"`
		AllytimeSupport       int    `xml:"allytime_support" json:"allytime_support" gqlgen:"allytimeSupport"`
		NoLeave               int    `xml:"no_leave" json:"no_leave" gqlgen:"noLeave"`
		NoJoin                int    `xml:"no_join" json:"no_join" gqlgen:"noJoin"`
		Limit                 int    `xml:"limit" json:"limit" gqlgen:"limit"`
		FixedAllies           int    `xml:"fixed_allies" json:"fixed_allies" gqlgen:"fixedAllies"`
		PointsMemberCount     int    `xml:"points_member_count" json:"points_member_count" gqlgen:"pointsMemberCount"`
		WarsMemberRequirement int    `xml:"wars_member_requirement" json:"wars_member_requirement" gqlgen:"warsMemberRequirement"`
		WarsPointsRequirement int    `xml:"wars_points_requirement" json:"wars_points_requirement" gqlgen:"warsPointsRequirement"`
		WarsAutoacceptDays    int    `xml:"wars_autoaccept_days" json:"wars_autoaccept_days" gqlgen:"warsAutoacceptDays"`
		Levels                int    `xml:"levels" json:"levels" gqlgen:"levels"`
		XpRequirements        string `xml:"xp_requirements" json:"xp_requirements" gqlgen:"xpRequirements"`
	} `xml:"ally" json:"ally" gqlgen:"ally"`
	Coord struct {
		Text            string `xml:",chardata" json:"-" gqlgen:"text"`
		MapSize         int    `xml:"map_size" json:"map_size" gqlgen:"mapSize"`
		Func            int    `xml:"func" json:"func" gqlgen:"func"`
		EmptyVillages   int    `xml:"empty_villages" json:"empty_villages" gqlgen:"emptyVillages"`
		BonusVillages   int    `xml:"bonus_villages" json:"bonus_villages" gqlgen:"bonusVillages"`
		BonusNew        int    `xml:"bonus_new" json:"bonus_new" gqlgen:"bonusNew"`
		Inner           int    `xml:"inner" json:"inner" gqlgen:"inner"`
		SelectStart     int    `xml:"select_start" json:"select_start" gqlgen:"selectStart"`
		VillageMoveWait int    `xml:"village_move_wait" json:"village_move_wait" gqlgen:"villageMoveWait"`
		NobleRestart    int    `xml:"noble_restart" json:"noble_restart" gqlgen:"nobleRestart"`
		StartVillages   int    `xml:"start_villages" json:"start_villages" gqlgen:"startVillages"`
	} `xml:"coord" json:"coord" gqlgen:"coord"`
	Sitter struct {
		Text  string `xml:",chardata" json:"-" gqlgen:"text"`
		Allow int    `xml:"allow" json:"allow" gqlgen:"allow"`
	} `xml:"sitter" json:"sitter" gqlgen:"sitter"`
	Sleep struct {
		Text     string `xml:",chardata" json:"-" gqlgen:"text"`
		Active   int    `xml:"active" json:"active" gqlgen:"active"`
		Delay    int    `xml:"delay" json:"delay" gqlgen:"delay"`
		Min      int    `xml:"min" json:"min" gqlgen:"min"`
		Max      int    `xml:"max" json:"max" gqlgen:"max"`
		MinAwake int    `xml:"min_awake" json:"min_awake" gqlgen:"minAwake"`
		MaxAwake int    `xml:"max_awake" json:"max_awake" gqlgen:"maxAwake"`
		WarnTime int    `xml:"warn_time" json:"warn_time" gqlgen:"warnTime"`
	} `xml:"sleep" json:"sleep" gqlgen:"sleep"`
	Night struct {
		Text      string `xml:",chardata" json:"-" gqlgen:"text"`
		Active    int    `xml:"active" json:"active" gqlgen:"active"`
		StartHour int    `xml:"start_hour" json:"start_hour" gqlgen:"startHour"`
		EndHour   int    `xml:"end_hour" json:"end_hour" gqlgen:"endHour"`
		DefFactor int    `xml:"def_factor" json:"def_factor" gqlgen:"defFactor"`
	} `xml:"night" json:"night" gqlgen:"night"`
	Win struct {
		Text  string `xml:",chardata" json:"-" gqlgen:"text"`
		Check int    `xml:"check" json:"check" gqlgen:"check"`
	} `xml:"win" json:"win" gqlgen:"win"`
}
