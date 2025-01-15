package database

type User struct {
	Id            int    `db:"id" json:"id"`
	Email         string `db:"email" json:"email"`
	EmailVerified bool   `db:"email_verified" json:"email_verified"`
}

type Hero struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Img  string `db:"img" json:"img"`
}

type Match struct {
	Id         int    `db:"id" json:"id"`
	HeroName   string `db:"hero_name" json:"hero_name"`
	HeroImg    string `db:"hero_img" json:"hero_img"`
	MatchId    int    `db:"match_id" json:"match_id"`
	Date       string `db:"date" json:"date"`
	Faction    string `db:"faction" json:"faction"`
	Position   string `db:"position" json:"position"`
	Result     bool   `db:"result" json:"result"`
	LaneResult bool   `db:"lane_result" json:"lane_result"`
}
