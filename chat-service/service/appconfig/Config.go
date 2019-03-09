package appconfig

const Config = `
userActiveWhenLastSeenMinutesAgo: 1
musicOnHoldUrl: https://actions.google.com/sounds/v1/cartoon/xylophone_tip_toe_scale_up.ogg
serviceDomain: strangerchat.suszczewicz.net
`

const MessagesEn = `
welcome: <prosody rate="90%">Hello! <break time="0.5s"/> Welcome to the stranger chat! <break time="0.5s"/>  This action contains user-generated content. <audio src="https://actions.google.com/sounds/v1/cartoon/cartoon_boing.ogg"/> We connect creative minds together by letting you chat with random strangers all over the world.<break time="0.5s"/> Meet new people online. <break time="0.5s"/> Share knowledge. <break time="0.5s"/> Get public opinion. <break time="0.5s"/> Enjoy and have fun! <break time="1s"/></prosody>
waitingForParticipant: <prosody rate="90%">Looking for a stranger. <break time="0.5s"/> Please wait!</prosody>
participantFoundInitiator: <prosody rate="90%"><audio src="https://actions.google.com/sounds/v1/cartoon/clang_and_wobble.ogg"/>We have found a stranger for you! <break time="0.5s"/> Please say hello!</prsody>
messageSent: <prosody rate="90%">Your message has been sent to your stranger. <break time="0.5s"/> Please wait for a reply!</prosody>
waitForReply:  <prosody rate="90%">Please wait for a reply!</prosody>
participantLeft: <audio src="https://actions.google.com/sounds/v1/cartoon/slide_whistle_to_drum.ogg"/><prosody rate="90%">Unfortunately your stranger has left the chat. No worries - we will find another one for you!</prosody>
participantSay: <audio src="https://actions.google.com/sounds/v1/cartoon/wood_plank_flicks.ogg"/><prosody rate="90%">Your stranger says:</prosody> <break time="1s"/>
participantReplyPrefix: <prosody rate="80%">
participantReplySufix: </prosody>
respondToParticipant: <prosody rate="90%">Please say something to respond or say next to find a new stranger.</prosody>
seeYouNextTime: <prosody rate="90%">See you next time!</prosody>
playManually: Please note that when using a keyboard as an input - you need to play the music on hold manually everytime. Please use your voice as an input for a better experience.
musicTitle: Music on hold
musicDescription: Please wait
checkStatus: Check status
nextParticipantWords: [ "Next", "Skip", "Find next", "Find another", "Find next stranger" ]
yesWords: [ "Absolutely", "Ack", "Acknowledged", "Affirmative", "Agree", "Aight", "All right", "All righty", "Alright", "Aye", "Certainly", "Cool", "Correct", "Definitely", "Exactly", "Fine", "Fo shizzle", "Fo sho", "Forizzle", "Gladly", "Good", "Indeed", "Indubitably", "Mhm", "Of course", "Ok", "Okay", "Precisely", "Right on", "Right", "Roger", "Sure thing", "Sure", "Surely", "True dat", "True", "Uh-huh", "Very well", "Way", "Word", "Ya", "Yaaasss", "Yah", "Ye", "Yea", "Yeah", "Yeh", "Yep", "Yeppers", "Yes", "Yessir", "Yessum", "Yis", "Yiss", "Yisss", "You bet", "Yup", "Yuppers", "Yus" ]
`

const MessagesPl = `
welcome: <prosody rate="90%">Cześć! <break time="0.5s"/> Witaj w czat ruletce! <break time="0.5s"/>  Ta akcja zawiera treści generowane przez użytkowników. <audio src="https://actions.google.com/sounds/v1/cartoon/cartoon_boing.ogg"/> Łączymy kreatywne umysły, pozwalając Ci rozmawiać z przypadkowymi ludźmi na całym świecie. <break time="0.5s"/> Poznaj nowych ludzi online. <break time="0.5s"/> Uzyskaj opinię publiczną. <break time="0.5s"/> Baw się dobrze! <break time="1s"/></prosody>
waitingForParticipant: <prosody rate="90%"> Szukam rozmówcy. <break time="0.5s"/> Proszę czekać!</prosody>
participantFoundInitiator: <prosody rate="90%"><audio src="https://actions.google.com/sounds/v1/cartoon/clang_and_wobble.ogg"/> Znaleźliśmy dla Ciebie rozmówcę. <break time="0.5s"/> Przywitaj się!</prsody>
messageSent: <prosody rate="90%"> Twoja wiadomość została wysłana. <break time="0.5s"/> Proszę czekać na odpowiedź!</prosody>
waitForReply:  <prosody rate="90%">Proszę czekać na odpowiedź!</prosody>
participantLeft: <audio src="https://actions.google.com/sounds/v1/cartoon/slide_whistle_to_drum.ogg"/><prosody rate="90%">Niestety Twój rozmówca opuścił czat - nie martw się znajdziemy dla Ciebie nowego.</prosody>
participantSay: <audio src="https://actions.google.com/sounds/v1/cartoon/wood_plank_flicks.ogg"/><prosody rate="90%">Twój rozmówca mówi:</prosody> <break time="1s"/>
participantReplyPrefix: <prosody rate="80%">
participantReplySufix: </prosody>
respondToParticipant: <prosody rate="90%">Proszę odpowiedzieć swojemu rozmówcy. Aby poszukać nowego rozmówcy powiedz "następny".</prosody>
seeYouNextTime: <prosody rate="90%">Do zobaczenia następnym razem!</prosody>
playManually: Jeżeli używasz klawiatury zamiast głosu - musisz ręcznie odgrywać muzykę na czekanie. Proszę używać głosu jeśli to możliwe.
musicTitle: Muzyka na czekanie
musicDescription: Proszę czekać
checkStatus: Sprawdź status
nextParticipantWords: [ "Następny", "Nastepny","Pomiń", "Pomin"]
yesWords: [ "OK", "Dobra", "Tak" ]
`
