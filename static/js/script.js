function search(entry) {
	if (entry) {
		$.ajax({
		    url: '/search/quick/',
		    type: 'POST',
		    data: {q: entry}
		}).done(function(res) { $('#results').html(res); });
	} else {
		$('#results').html('<i>&nbsp;Suchbegriff eingeben</i>');
	};
}

$(function () {
  $('[data-toggle="popover"]').popover()
})

function showInfo(id) {
    if (id) {
		 $('#countryInfo').find('.modal-title').text(id);
                 $('#countryInfo').find('.modal-body').html($('#inf_'+id).html());
	} else {
		$('#countryInfo').find('.modal-title').text('Informationen zum Land');
        	$('#countryInfo').find('.modal-body').html('Informationen zum Land');
	};
    $('#countryInfo').modal();
}

/*$(document).on("submit", "form", function(event) {
    // Das eigentliche Absenden verhindern
    event.preventDefault();
    
    // Das sendende Formular und die Metadaten bestimmen
    var form = $(this); // Dieser Zeiger $(this) oder $("form"), falls die ID form im HTML exisitiert, klappt übrigens auch ohne jQuery ;)
    var action = form.attr("action"), // attr() kann enweder den aktuellen Inhalt des gennanten Attributs auslesen, oder setzt ein neuen Wert, falls ein zweiter Parameter gegeben ist
        method = form.attr("method"),
        data   = form.serialize(); // baut die Daten zu einem String nach dem Muster vorname=max&nachname=Müller&alter=42 ... zusammen
        
    // Der eigentliche AJAX Aufruf
    $.ajax({
        url : action,
        type : method,
        data : data
    }).done(function (data) {
        // Bei Erfolg
        $('#legende').html(data);
    }).fail(function() {
        // Bei Fehler
        alert("Fehler!");
    }).always(function() {
        // Immer
        //alert("Beendet!");
    });
});
*/
