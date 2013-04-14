# Overview
get 
http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe

# history der saisons
post
http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe
param: 
* filter_saison_id: 1
* task: veranstaltungen 

# Begegnungen
get 
http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=veranstaltung&veranstaltungid=8

# Begegnung
get
http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=8&id=2

# Spieler
get
http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=spieler_details&id=992
