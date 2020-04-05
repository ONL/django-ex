# Interaktive Karten für den Geographieunterricht

Diese Anwendung dient der Unterstützung im Geographieunterricht. Sie stellt Informationen und Aufgaben in interaktiven Karten bereit.

Der Container benötigt keine Volumes oder Config.

Das Admin-Passwort wird mit der Umgebungsvariable ``PASSWORD`` gesetzt, der Container lauscht auf Port ``8080``

``docker run -d -e PASSWORD=password -p 80:8080 paulritter/interactive-maps``

## License

This code is dedicated to the public domain to the maximum extent permitted by applicable law, pursuant to [CC0](http://creativecommons.org/publicdomain/zero/1.0/).
