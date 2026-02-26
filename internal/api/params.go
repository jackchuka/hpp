package api

// GourmetSearchParams contains all parameters for /gourmet/v1/
type GourmetSearchParams struct {
	ID       []string `url:"id,omitempty"`
	Name     *string  `url:"name,omitempty"`
	NameKana *string  `url:"name_kana,omitempty"`
	NameAny  *string  `url:"name_any,omitempty"`
	Tel      *string  `url:"tel,omitempty"`
	Address  *string  `url:"address,omitempty"`
	Keyword  *string  `url:"keyword,omitempty"`

	// Location
	Lat   *float64 `url:"lat,omitempty"`
	Lng   *float64 `url:"lng,omitempty"`
	Range *int     `url:"range,omitempty"`
	Datum *string  `url:"datum,omitempty"`

	// Area filters
	LargeServiceArea *string  `url:"large_service_area,omitempty"`
	ServiceArea      []string `url:"service_area,omitempty"`
	LargeArea        []string `url:"large_area,omitempty"`
	MiddleArea       []string `url:"middle_area,omitempty"`
	SmallArea        []string `url:"small_area,omitempty"`

	// Category filters
	Genre             []string `url:"genre,omitempty"`
	Budget            []string `url:"budget,omitempty"`
	CreditCardFilter  []string `url:"credit_card,omitempty"`
	Special           []string `url:"special,omitempty"`
	SpecialOr         []string `url:"special_or,omitempty"`
	SpecialCategory   []string `url:"special_category,omitempty"`
	SpecialCategoryOr []string `url:"special_category_or,omitempty"`

	// Capacity
	PartyCapacity *int `url:"party_capacity,omitempty"`

	// Boolean filters (serialized as 0/1)
	WiFi         bool `url:"wifi,int,omitempty"`
	Wedding      bool `url:"wedding,int,omitempty"`
	Course       bool `url:"course,int,omitempty"`
	FreeDrink    bool `url:"free_drink,int,omitempty"`
	FreeFood     bool `url:"free_food,int,omitempty"`
	PrivateRoom  bool `url:"private_room,int,omitempty"`
	Horigotatsu  bool `url:"horigotatsu,int,omitempty"`
	Tatami       bool `url:"tatami,int,omitempty"`
	Cocktail     bool `url:"cocktail,int,omitempty"`
	Shochu       bool `url:"shochu,int,omitempty"`
	Sake         bool `url:"sake,int,omitempty"`
	Wine         bool `url:"wine,int,omitempty"`
	Card         bool `url:"card,int,omitempty"`
	NonSmoking   bool `url:"non_smoking,int,omitempty"`
	Charter      bool `url:"charter,int,omitempty"`
	Ktai         bool `url:"ktai,int,omitempty"`
	Parking      bool `url:"parking,int,omitempty"`
	BarrierFree  bool `url:"barrier_free,int,omitempty"`
	Sommelier    bool `url:"sommelier,int,omitempty"`
	NightView    bool `url:"night_view,int,omitempty"`
	OpenAir      bool `url:"open_air,int,omitempty"`
	Show         bool `url:"show,int,omitempty"`
	Equipment    bool `url:"equipment,int,omitempty"`
	Karaoke      bool `url:"karaoke,int,omitempty"`
	Band         bool `url:"band,int,omitempty"`
	TV           bool `url:"tv,int,omitempty"`
	Lunch        bool `url:"lunch,int,omitempty"`
	Midnight     bool `url:"midnight,int,omitempty"`
	MidnightMeal bool `url:"midnight_meal,int,omitempty"`
	English      bool `url:"english,int,omitempty"`
	Pet          bool `url:"pet,int,omitempty"`
	Child        bool `url:"child,int,omitempty"`
	KtaiCoupon   *int `url:"ktai_coupon,omitempty"`

	// Output control
	Type  *string `url:"type,omitempty"`
	Order *int    `url:"order,omitempty"`
	Start *int    `url:"start,omitempty"`
	Count *int    `url:"count,omitempty"`
}

// ShopSearchParams contains all parameters for /shop/v1/
type ShopSearchParams struct {
	Keyword *string `url:"keyword,omitempty"`
	Tel     *string `url:"tel,omitempty"`
	Start   *int    `url:"start,omitempty"`
	Count   *int    `url:"count,omitempty"`
}

// LargeAreaParams contains parameters for /large_area/v1/
type LargeAreaParams struct {
	LargeArea []string `url:"large_area,omitempty"`
	Keyword   *string  `url:"keyword,omitempty"`
}

// MiddleAreaParams contains parameters for /middle_area/v1/
type MiddleAreaParams struct {
	MiddleArea []string `url:"middle_area,omitempty"`
	LargeArea  []string `url:"large_area,omitempty"`
	Keyword    *string  `url:"keyword,omitempty"`
	Start      *int     `url:"start,omitempty"`
	Count      *int     `url:"count,omitempty"`
}

// SmallAreaParams contains parameters for /small_area/v1/
type SmallAreaParams struct {
	SmallArea  []string `url:"small_area,omitempty"`
	MiddleArea []string `url:"middle_area,omitempty"`
	Keyword    *string  `url:"keyword,omitempty"`
	Start      *int     `url:"start,omitempty"`
	Count      *int     `url:"count,omitempty"`
}

// GenreParams contains parameters for /genre/v1/
type GenreParams struct {
	Code    []string `url:"code,omitempty"`
	Keyword *string  `url:"keyword,omitempty"`
}

// SpecialParams contains parameters for /special/v1/
type SpecialParams struct {
	Special         []string `url:"special,omitempty"`
	SpecialCategory []string `url:"special_category,omitempty"`
}

// SpecialCategoryParams contains parameters for /special_category/v1/
type SpecialCategoryParams struct {
	SpecialCategory []string `url:"special_category,omitempty"`
}
