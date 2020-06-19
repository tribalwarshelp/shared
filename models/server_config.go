package models

import "encoding/xml"

type Config struct {
	XMLName   xml.Name `xml:"config" json:"xmlName" gqlgen:"xmlName"`
	Text      string   `xml:",chardata" json:"text" gqlgen:"text"`
	Speed     string   `xml:"speed" json:"speed" gqlgen:"speed"`
	UnitSpeed string   `xml:"unit_speed" json:"unitSpeed" gqlgen:"unitSpeed"`
	Moral     string   `xml:"moral" json:"moral" gqlgen:"moral"`
	Build     struct {
		Text    string `xml:",chardata" json:"text" gqlgen:"text"`
		Destroy string `xml:"destroy" json:"destroy" gqlgen:"destroy"`
	} `xml:"build" json:"build" gqlgen:"build"`
	Misc struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		KillRanking     string `xml:"kill_ranking" json:"killRanking" gqlgen:"killRanking"`
		Tutorial        string `xml:"tutorial" json:"tutorial" gqlgen:"tutorial"`
		TradeCancelTime string `xml:"trade_cancel_time" json:"tradeCancelTime" gqlgen:"tradeCancelTime"`
	} `xml:"misc" json:"misc" gqlgen:"misc"`
	Commands struct {
		Text              string `xml:",chardata" json:"text" gqlgen:"text"`
		MillisArrival     string `xml:"millis_arrival" json:"millisArrival" gqlgen:"millisArrival"`
		CommandCancelTime string `xml:"command_cancel_time" json:"commandCancelTime" gqlgen:"commandCancelTime"`
	} `xml:"commands" json:"commands" gqlgen:"commands"`
	Newbie struct {
		Text                 string `xml:",chardata" json:"text" gqlgen:"text"`
		Days                 string `xml:"days" json:"days" gqlgen:"days"`
		RatioDays            string `xml:"ratio_days" json:"ratioDays" gqlgen:"ratioDays"`
		Ratio                string `xml:"ratio" json:"ratio" gqlgen:"ratio"`
		RemoveNewbieVillages string `xml:"removeNewbieVillages" json:"removeNewbieVillages" gqlgen:"removeNewbieVillages"`
	} `xml:"newbie" json:"newbie" gqlgen:"newbie"`
	Game struct {
		Text               string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildtimeFormula   string `xml:"buildtime_formula" json:"buildtimeFormula" gqlgen:"buildtimeFormula"`
		Knight             string `xml:"knight" json:"knight" gqlgen:"knight"`
		KnightNewItems     string `xml:"knight_new_items" json:"knightNewItems" gqlgen:"knightNewItems"`
		Archer             string `xml:"archer" json:"archer" gqlgen:"archer"`
		Tech               string `xml:"tech" json:"tech" gqlgen:"tech"`
		FarmLimit          string `xml:"farm_limit" json:"farmLimit" gqlgen:"farmLimit"`
		Church             string `xml:"church" json:"church" gqlgen:"church"`
		Watchtower         string `xml:"watchtower" json:"watchtower" gqlgen:"watchtower"`
		Stronghold         string `xml:"stronghold" json:"stronghold" gqlgen:"stronghold"`
		FakeLimit          string `xml:"fake_limit" json:"fakeLimit" gqlgen:"fakeLimit"`
		BarbarianRise      string `xml:"barbarian_rise" json:"barbarianRise" gqlgen:"barbarianRise"`
		BarbarianShrink    string `xml:"barbarian_shrink" json:"barbarianShrink" gqlgen:"barbarianShrink"`
		BarbarianMaxPoints string `xml:"barbarian_max_points" json:"barbarianMaxPoints" gqlgen:"barbarianMaxPoints"`
		Hauls              string `xml:"hauls" json:"hauls" gqlgen:"hauls"`
		HaulsBase          string `xml:"hauls_base" json:"haulsBase" gqlgen:"haulsBase"`
		HaulsMax           string `xml:"hauls_max" json:"haulsMax" gqlgen:"haulsMax"`
		BaseProduction     string `xml:"base_production" json:"baseProduction" gqlgen:"baseProduction"`
		Event              string `xml:"event" json:"event" gqlgen:"event"`
		SuppressEvents     string `xml:"suppress_events" json:"suppressEvents" gqlgen:"suppressEvents"`
	} `xml:"game" json:"game" gqlgen:"game"`
	Buildings struct {
		Text             string `xml:",chardata" json:"text" gqlgen:"text"`
		CustomMain       string `xml:"custom_main" json:"customMain" gqlgen:"customMain"`
		CustomFarm       string `xml:"custom_farm" json:"customFarm" gqlgen:"customFarm"`
		CustomStorage    string `xml:"custom_storage" json:"customStorage" gqlgen:"customStorage"`
		CustomPlace      string `xml:"custom_place" json:"customPlace" gqlgen:"customPlace"`
		CustomBarracks   string `xml:"custom_barracks" json:"customBarracks" gqlgen:"customBarracks"`
		CustomChurch     string `xml:"custom_church" json:"customChurch" gqlgen:"customChurch"`
		CustomSmith      string `xml:"custom_smith" json:"customSmith" gqlgen:"customSmith"`
		CustomWood       string `xml:"custom_wood" json:"customWood" gqlgen:"customWood"`
		CustomStone      string `xml:"custom_stone" json:"customStone" gqlgen:"customStone"`
		CustomIron       string `xml:"custom_iron" json:"customIron" gqlgen:"customIron"`
		CustomMarket     string `xml:"custom_market" json:"customMarket" gqlgen:"customMarket"`
		CustomStable     string `xml:"custom_stable" json:"customStable" gqlgen:"customStable"`
		CustomWall       string `xml:"custom_wall" json:"customWall" gqlgen:"customWall"`
		CustomGarage     string `xml:"custom_garage" json:"customGarage" gqlgen:"customGarage"`
		CustomHide       string `xml:"custom_hide" json:"customHide" gqlgen:"customHide"`
		CustomSnob       string `xml:"custom_snob" json:"customSnob" gqlgen:"customSnob"`
		CustomStatue     string `xml:"custom_statue" json:"customStatue" gqlgen:"customStatue"`
		CustomWatchtower string `xml:"custom_watchtower" json:"customWatchtower" gqlgen:"customWatchtower"`
	} `xml:"buildings" json:"buildings" gqlgen:"buildings"`
	Snob struct {
		Text          string `xml:",chardata" json:"text" gqlgen:"text"`
		Gold          string `xml:"gold" json:"gold" gqlgen:"gold"`
		CheapRebuild  string `xml:"cheap_rebuild" json:"cheapRebuild" gqlgen:"cheapRebuild"`
		Rise          string `xml:"rise" json:"rise" gqlgen:"rise"`
		MaxDist       string `xml:"max_dist" json:"maxDist" gqlgen:"maxDist"`
		Factor        string `xml:"factor" json:"factor" gqlgen:"factor"`
		CoinWood      string `xml:"coin_wood" json:"coinWood" gqlgen:"coinWood"`
		CoinStone     string `xml:"coin_stone" json:"coinStone" gqlgen:"coinStone"`
		CoinIron      string `xml:"coin_iron" json:"coinIron" gqlgen:"coinIron"`
		NoBarbConquer string `xml:"no_barb_conquer" json:"noBarbConquer" gqlgen:"noBarbConquer"`
	} `xml:"snob" json:"snob" gqlgen:"snob"`
	Ally struct {
		Text                  string `xml:",chardata" json:"text" gqlgen:"text"`
		NoHarm                string `xml:"no_harm" json:"noHarm" gqlgen:"noHarm"`
		NoOtherSupport        string `xml:"no_other_support" json:"noOtherSupport" gqlgen:"noOtherSupport"`
		AllytimeSupport       string `xml:"allytime_support" json:"allytimeSupport" gqlgen:"allytimeSupport"`
		NoLeave               string `xml:"no_leave" json:"noLeave" gqlgen:"noLeave"`
		NoJoin                string `xml:"no_join" json:"noJoin" gqlgen:"noJoin"`
		Limit                 string `xml:"limit" json:"limit" gqlgen:"limit"`
		FixedAllies           string `xml:"fixed_allies" json:"fixedAllies" gqlgen:"fixedAllies"`
		PointsMemberCount     string `xml:"points_member_count" json:"pointsMemberCount" gqlgen:"pointsMemberCount"`
		WarsMemberRequirement string `xml:"wars_member_requirement" json:"warsMemberRequirement" gqlgen:"warsMemberRequirement"`
		WarsPointsRequirement string `xml:"wars_points_requirement" json:"warsPointsRequirement" gqlgen:"warsPointsRequirement"`
		WarsAutoacceptDays    string `xml:"wars_autoaccept_days" json:"warsAutoacceptDays" gqlgen:"warsAutoacceptDays"`
		Levels                string `xml:"levels" json:"levels" gqlgen:"levels"`
		XpRequirements        string `xml:"xp_requirements" json:"xpRequirements" gqlgen:"xpRequirements"`
	} `xml:"ally" json:"ally" gqlgen:"ally"`
	Coord struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MapSize         string `xml:"map_size" json:"mapSize" gqlgen:"mapSize"`
		Func            string `xml:"func" json:"func" gqlgen:"func"`
		EmptyVillages   string `xml:"empty_villages" json:"emptyVillages" gqlgen:"emptyVillages"`
		BonusVillages   string `xml:"bonus_villages" json:"bonusVillages" gqlgen:"bonusVillages"`
		Inner           string `xml:"inner" json:"inner" gqlgen:"inner"`
		SelectStart     string `xml:"select_start" json:"selectStart" gqlgen:"selectStart"`
		VillageMoveWait string `xml:"village_move_wait" json:"villageMoveWait" gqlgen:"villageMoveWait"`
		NobleRestart    string `xml:"noble_restart" json:"nobleRestart" gqlgen:"nobleRestart"`
		StartVillages   string `xml:"start_villages" json:"startVillages" gqlgen:"startVillages"`
	} `xml:"coord" json:"coord" gqlgen:"coord"`
	Sitter struct {
		Text  string `xml:",chardata" json:"text" gqlgen:"text"`
		Allow string `xml:"allow" json:"allow" gqlgen:"allow"`
	} `xml:"sitter" json:"sitter" gqlgen:"sitter"`
	Sleep struct {
		Text     string `xml:",chardata" json:"text" gqlgen:"text"`
		Active   string `xml:"active" json:"active" gqlgen:"active"`
		Delay    string `xml:"delay" json:"delay" gqlgen:"delay"`
		Min      string `xml:"min" json:"min" gqlgen:"min"`
		Max      string `xml:"max" json:"max" gqlgen:"max"`
		MinAwake string `xml:"min_awake" json:"minAwake" gqlgen:"minAwake"`
		MaxAwake string `xml:"max_awake" json:"maxAwake" gqlgen:"maxAwake"`
		WarnTime string `xml:"warn_time" json:"warnTime" gqlgen:"warnTime"`
	} `xml:"sleep" json:"sleep" gqlgen:"sleep"`
	Night struct {
		Text      string `xml:",chardata" json:"text" gqlgen:"text"`
		Active    string `xml:"active" json:"active" gqlgen:"active"`
		StartHour string `xml:"start_hour" json:"startHour" gqlgen:"startHour"`
		EndHour   string `xml:"end_hour" json:"endHour" gqlgen:"endHour"`
		DefFactor string `xml:"def_factor" json:"defFactor" gqlgen:"defFactor"`
		Duration  string `xml:"duration" json:"duration" gqlgen:"duration"`
	} `xml:"night" json:"night" gqlgen:"night"`
	Win struct {
		Text  string `xml:",chardata" json:"text" gqlgen:"text"`
		Check string `xml:"check" json:"check" gqlgen:"check"`
	} `xml:"win" json:"win" gqlgen:"win"`
}
