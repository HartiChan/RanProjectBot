// coding=utf-8

// Copyright (c) 2020 HartiChan

// 8888888b.                          Y88b   d88P          888                                              .d8888b.                888
// 888   Y88b                          Y88b d88P           888                                             d88P  Y88b               888
// 888    888                           Y88o88P            888                                             888    888               888
// 888   d88P  8888b.  88888b.           Y888P     8888b.  888  888 888  888 88888b.d88b.   .d88b.         888         .d88b.   .d88888  .d88b.
// 8888888P"      "88b 888 "88b           888         "88b 888 .88P 888  888 888 "888 "88b d88""88b        888        d88""88b d88" 888 d8P  Y8b
// 888 T88b   .d888888 888  888           888     .d888888 888888K  888  888 888  888  888 888  888 888888 888    888 888  888 888  888 88888888
// 888  T88b  888  888 888  888           888     888  888 888 "88b Y88b 888 888  888  888 Y88..88P        Y88b  d88P Y88..88P Y88b 888 Y8b.
// 888   T88b "Y888888 888  888           888     "Y888888 888  888  "Y88888 888  888  888  "Y88P"          "Y8888P"   "Y88P"   "Y88888  "Y8888

package RanYakumo

func Run(polling bool) {
	var bot, _ = NewBot(ReadFileString("token.txt"))
	bot.MasterID = "masterID"

	registerCommands(bot)

	bot.SetUpdateHandler(updater)

	bot.Idle()
}
