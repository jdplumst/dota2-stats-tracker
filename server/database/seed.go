package database

import "log"

func SeedDatabase() {
	db, err := NewConnection()
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	defer db.Close()

	// Clear hero table
	_, err = db.Exec("DELETE * FROM hero")
	if err != nil {
		log.Fatal("error clearing hero table in database", err)
	}

	// Seed hero table
	_, err = db.Exec(`INSERT INTO "hero" ("id", "name", "img") VALUES
		('1', 'Anti-Mage', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_antimage_png.png'),
		('2', 'Axe', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_axe_png.png'),
		('3', 'Bane', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_bane_png.png'),
		('4', 'Bloodseeker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_bloodseeker_png.png'),
		('5', 'Crystal Maiden', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_crystal_maiden_png.png'),
		('6', 'Drow Ranger', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_drow_ranger_png.png'),
		('7', 'Earthshaker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_earthshaker_png.png'),
		('8', 'Juggernaut', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_juggernaut_png.png'),
		('9', 'Mirana', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_mirana_png.png'),
		('10', 'Morphling', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_morphling_png.png'),
		('11', 'Shadow Fiend', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_nevermore_png.png'),
		('12', 'Phantom Lancer', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_phantom_lancer_png.png'),
		('13', 'Puck', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_puck_png.png'),
		('14', 'Pudge', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_pudge_png.png'),
		('15', 'Razor', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_razor_png.png'),
		('16', 'Sand King', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_sand_king_png.png'),
		('17', 'Storm Spirit', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_storm_spirit_png.png'),
		('18', 'Sven', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_sven_png.png'),
		('19', 'Tiny', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_tiny_png.png'),
		('20', 'Vengeful Spirit', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_vengefulspirit_png.png'),
		('21', 'Windranger', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_windrunner_png.png'),
		('22', 'Zeus', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_zuus_png.png'),
		('23', 'Kunkka', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_kunkka_png.png'),
		('24', 'Lina', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_lina_png.png'),
		('25', 'Lion', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_lion_png.png'),
		('26', 'Shadow Shaman', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_shadow_shaman_png.png'),
		('27', 'Slardar', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_slardar_png.png'),
		('28', 'Tidehunter', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_tidehunter_png.png'),
		('29', 'Witch Doctor', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_witch_doctor_png.png'),
		('30', 'Lich', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_lich_png.png'),
		('31', 'Riki', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_riki_png.png'),
		('32', 'Enigma', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_enigma_png.png'),
		('33', 'Tinker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_tinker_png.png'),
		('34', 'Sniper', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_sniper_png.png'),
		('35', 'Necrophos', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_necrolyte_png.png'),
		('36', 'Warlock', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_warlock_png.png'),
		('37', 'Beastmaster', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_beastmaster_png.png'),
		('38', 'Queen of Pain', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_queenofpain_png.png'),
		('39', 'Venomancer', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_venomancer_png.png'),
		('40', 'Faceless Void', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_faceless_void_png.png'),
		('41', 'Wraith King', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_skeleton_king_png.png'),
		('42', 'Death Prophet', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_death_prophet_png.png'),
		('43', 'Phantom Assassin', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_phantom_assassin_png.png'),
		('44', 'Pugna', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_pugna_png.png'),
		('45', 'Templar Assassin', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_templar_assassin_png.png'),
		('46', 'Viper', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_viper_png.png'),
		('47', 'Luna', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_luna_png.png'),
		('48', 'Dragon Knight', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_dragon_knight_png.png'),
		('49', 'Dazzle', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_dazzle_png.png'),
		('50', 'Clockwerk', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_rattletrap_png.png'),
		('51', 'Leshrac', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_leshrac_png.png'),
		('52', 'Nature''s Prophet', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_furion_png.png'),
		('53', 'Lifestealer', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_life_stealer_png.png'),
		('54', 'Dark Seer', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_dark_seer_png.png'),
		('55', 'Clinkz', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_clinkz_png.png'),
		('56', 'Omniknight', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_omniknight_png.png'),
		('57', 'Enchantress', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_enchantress_png.png'),
		('58', 'Huskar', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_huskar_png.png'),
		('59', 'Night Stalker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_night_stalker_png.png'),
		('60', 'Broodmother', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_broodmother_png.png'),
		('61', 'Bounty Hunter', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_bounty_hunter_png.png'),
		('62', 'Weaver', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_weaver_png.png'),
		('63', 'Jakiro', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_jakiro_png.png'),
		('64', 'Batrider', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_batrider_png.png'),
		('65', 'Chen', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_chen_png.png'),
		('66', 'Spectre', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_spectre_png.png'),
		('67', 'Ancient Apparition', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_ancient_apparition_png.png'),
		('68', 'Doom', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_doom_bringer_png.png'),
		('69', 'Ursa', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_ursa_png.png'),
		('70', 'Spirit Breaker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_spirit_breaker_png.png'),
		('71', 'Gyrocopter', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_gyrocopter_png.png'),
		('72', 'Alchemist', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_alchemist_png.png'),
		('73', 'Invoker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_invoker_png.png'),
		('74', 'Silencer', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_silencer_png.png'),
		('75', 'Outworld Destoyer', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_obsidian_destroyer_png.png'),
		('76', 'Lycan', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_lycan_png.png'),
		('77', 'Brewmaster', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_brewmaster_png.png'),
		('78', 'Shadow Demon', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_shadow_demon_png.png'),
		('79', 'Lone Druid', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_lone_druid_png.png'),
		('80', 'Chaos Knight', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_chaos_knight_png.png'),
		('81', 'Meepo', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_meepo_png.png'),
		('82', 'Treant Protector', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_treant_png.png'),
		('83', 'Ogre Magi', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_ogre_magi_png.png'),
		('84', 'Undying', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_undying_png.png'),
		('85', 'Rubick', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_rubick_png.png'),
		('86', 'Disruptor', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_disruptor_png.png'),
		('87', 'Nyx Assassin', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_nyx_assassin_png.png'),
		('88', 'Naga Siren', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_naga_siren_png.png'),
		('89', 'Keeper of the Light', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_keeper_of_the_light_png.png'),
		('90', 'Io', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_wisp_png.png'),
		('91', 'Visage', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_visage_png.png'),
		('92', 'Slark', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_slark_png.png'),
		('93', 'Medusa', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_medusa_png.png'),
		('94', 'Troll Warlord', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_troll_warlord_png.png'),
		('95', 'Centaur Warrunner', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_centaur_png.png'),
		('96', 'Magnus', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_magnataur_png.png'),
		('97', 'Timbersaw', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_shredder_png.png'),
		('98', 'Bristleback', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_bristleback_png.png'),
		('99', 'Tusk', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_tusk_png.png'),
		('100', 'Skywrath Mage', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_skywrath_mage_png.png'),
		('101', 'Abaddon', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_abaddon_png.png'),
		('102', 'Elder Titan', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_elder_titan_png.png'),
		('103', 'Legion Commander', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_legion_commander_png.png'),
		('104', 'Techies', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_techies_png.png'),
		('105', 'Ember Spirit', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_ember_spirit_png.png'),
		('106', 'Earth Spirit', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_earth_spirit_png.png'),
		('107', 'Underlord', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_abyssal_underlord_png.png'),
		('108', 'Terrorblade', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_terrorblade_png.png'),
		('109', 'Phoenix', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_phoenix_png.png'),
		('110', 'Oracle', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_oracle_png.png'),
		('111', 'Winter Wyvern', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_winter_wyvern_png.png'),
		('112', 'Arc Warden', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_arc_warden_png.png'),
		('113', 'Monkey King', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_monkey_king_png.png'),
		('114', 'Dark Willow', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_dark_willow_png.png'),
		('115', 'Pangolier', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_pangolier_png.png'),
		('116', 'Grimstroke', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_grimstroke_png.png'),
		('117', 'Hoodwink', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_hoodwink_png.png'),
		('118', 'Void Spirit', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_void_spirit_png.png'),
		('119', 'Snapfire', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_snapfire_png.png'),
		('120', 'Mars', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_mars_png.png'),
		('121', 'Ringmaster', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_ringmaster_png.png'),
		('122', 'Dawnbreaker', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_dawnbreaker_png.png'),
		('123', 'Marci', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_marci_png.png'),
		('124', 'Primal Beast', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_primal_beast_png.png'),
		('125', 'Muerta', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_muerta_png.png'),
		('126', 'Kez', 'https://dotabase.dillerm.io/vpk/panorama/images/heroes/icons/npc_dota_hero_kez_png.png');
	`)
	if err != nil {
		log.Fatal("error seeding hero table in database", err)
	}

	// Clear faction table
	_, err = db.Exec("DELETE * FROM faction")
	if err != nil {
		log.Fatal("error clearing faction table in database", err)
	}

	// Seed faction table
	_, err = db.Exec(`INSERT INTO "faction" ("id", "name") VALUES
		('1', 'Radiant'),
		('2', 'Dire');
	`)
	if err != nil {
		log.Fatal("error seeding faction table database", err)
	}

	// Clear position table
	_, err = db.Exec("DELETE * FROM position")
	if err != nil {
		log.Fatal("error clearing position table in database", err)
	}

	// Seed position table
	_, err = db.Exec(`INSERT INTO "position" ("id", "name") VALUES
		('1', 'Carry'),
		('2', 'Mid'),
		('3', 'Offlane'),
		('4', 'Soft Support'),
		('5', 'Hard Support');
	`)
	if err != nil {
		log.Fatal("error seeding position table in database", err)
	}
}
