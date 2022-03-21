package dotago

type Client struct {
	Token string
}

type CommonParams struct {
	Key      string `url:"key,omitempty"`
	Format   string `url:"format,omitempty"`
	Language string `url:"language,omitempty"`
}

type MatchHistoryParams struct {
	CommonParams
	HeroID    int32 `url:"hero_id,omitempty"`
	GameMode  int32 `url:"game_mode,omitempty"`
	AccountID int   `url:"account_id,omitempty"`
}

type MatchHistory struct {
	MatchID       int   `json:"match_id"`
	MatchSeqNum   int64 `json:"match_seq_num"`
	StartTime     int   `json:"start_time"`
	LobbyType     int   `json:"lobby_type"`
	RadiantTeamID int   `json:"radiant_team_id"`
	DireTeamID    int   `json:"dire_team_id"`
	Players       []struct {
		AccountID  int `json:"account_id,omitempty"`
		PlayerSlot int `json:"player_slot"`
		HeroID     int `json:"hero_id"`
	}
}

type MatchHistoryResult struct {
	Result struct {
		Status           int            `json:"status"`
		NumResults       int            `json:"num_results"`
		TotalResults     int            `json:"total_results"`
		ResultsRemaining int            `json:"results_remaining"`
		Matches          []MatchHistory `json:"matches"`
	} `json:"result"`
}

type MatchDetailsParams struct {
	CommonParams
	MatchID int `url:"match_id,omitempty"`
}

type Player struct {
	AccountID         int `json:"account_id"`
	PlayerSlot        int `json:"player_slot"`
	HeroID            int `json:"hero_id"`
	Item0             int `json:"item_0"`
	Item1             int `json:"item_1"`
	Item2             int `json:"item_2"`
	Item3             int `json:"item_3"`
	Item4             int `json:"item_4"`
	Item5             int `json:"item_5"`
	Backpack0         int `json:"backpack_0"`
	Backpack1         int `json:"backpack_1"`
	Backpack2         int `json:"backpack_2"`
	ItemNeutral       int `json:"item_neutral"`
	Kills             int `json:"kills"`
	Deaths            int `json:"deaths"`
	Assists           int `json:"assists"`
	LeaverStatus      int `json:"leaver_status"`
	LastHits          int `json:"last_hits"`
	Denies            int `json:"denies"`
	GoldPerMin        int `json:"gold_per_min"`
	XpPerMin          int `json:"xp_per_min"`
	Level             int `json:"level"`
	NetWorth          int `json:"net_worth"`
	AghanimsScepter   int `json:"aghanims_scepter"`
	AghanimsShard     int `json:"aghanims_shard"`
	Moonshard         int `json:"moonshard"`
	HeroDamage        int `json:"hero_damage"`
	TowerDamage       int `json:"tower_damage"`
	HeroHealing       int `json:"hero_healing"`
	Gold              int `json:"gold"`
	GoldSpent         int `json:"gold_spent"`
	ScaledHeroDamage  int `json:"scaled_hero_damage"`
	ScaledTowerDamage int `json:"scaled_tower_damage"`
	ScaledHeroHealing int `json:"scaled_hero_healing"`
	AbilityUpgrades   []struct {
		Ability int `json:"ability"`
		Time    int `json:"time"`
		Level   int `json:"level"`
	} `json:"ability_upgrades"`
}

type MatchDetailsResult struct {
	Result struct {
		Players               []Player `json:"players"`
		RadiantWin            bool     `json:"radiant_win"`
		Duration              int      `json:"duration"`
		PreGameDuration       int      `json:"pre_game_duration"`
		StartTime             int      `json:"start_time"`
		MatchID               int64    `json:"match_id"`
		MatchSeqNum           int64    `json:"match_seq_num"`
		TowerStatusRadiant    int      `json:"tower_status_radiant"`
		TowerStatusDire       int      `json:"tower_status_dire"`
		BarracksStatusRadiant int      `json:"barracks_status_radiant"`
		BarracksStatusDire    int      `json:"barracks_status_dire"`
		Cluster               int      `json:"cluster"`
		FirstBloodTime        int      `json:"first_blood_time"`
		LobbyType             int      `json:"lobby_type"`
		HumanPlayers          int      `json:"human_players"`
		Leagueid              int      `json:"leagueid"`
		PositiveVotes         int      `json:"positive_votes"`
		NegativeVotes         int      `json:"negative_votes"`
		GameMode              int      `json:"game_mode"`
		Flags                 int      `json:"flags"`
		Engine                int      `json:"engine"`
		RadiantScore          int      `json:"radiant_score"`
		DireScore             int      `json:"dire_score"`
		PicksBans             []struct {
			IsPick bool `json:"is_pick"`
			HeroID int  `json:"hero_id"`
			Team   int  `json:"team"`
			Order  int  `json:"order"`
		} `json:"picks_bans"`
	} `json:"result"`
}

type HeroData map[string]struct {
	Name string
}
type PlayerData struct {
	Name      string
	ID        int
	DiscordID string
}

type GiphyRandomSearch struct {
	Data struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			FixedWidthStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			PreviewGif struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_gif"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_downsampled"`
			Preview struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"preview"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_small"`
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidth struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width"`
			DownsizedStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_still"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_medium"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"original_mp4"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_large"`
			PreviewWebp struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_webp"`
			Original struct {
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"original"`
			OriginalStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original_still"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_small_still"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_small"`
			Looping struct {
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
			} `json:"looping"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"downsized_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height"`
			Four80WStill struct {
				URL    string `json:"url"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"480w_still"`
		} `json:"images"`
		User struct {
			AvatarURL    string `json:"avatar_url"`
			BannerImage  string `json:"banner_image"`
			BannerURL    string `json:"banner_url"`
			ProfileURL   string `json:"profile_url"`
			Username     string `json:"username"`
			DisplayName  string `json:"display_name"`
			Description  string `json:"description"`
			IsVerified   bool   `json:"is_verified"`
			WebsiteURL   string `json:"website_url"`
			InstagramURL string `json:"instagram_url"`
		} `json:"user"`
	} `json:"data"`
	Meta struct {
		Msg        string `json:"msg"`
		Status     int    `json:"status"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}
