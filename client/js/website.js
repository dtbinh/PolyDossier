/**
 * @fileoverview Outils important ayant lien au concept d'utilisateur.
 *
 * Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been
 * the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley
 * of type and scrambled it to make a type specimen book. It has survived not only five centuries,
 * but also the leap into electronic typesetting, remaining essentially unchanged. It was
 * popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
 * and more recently with desktop publishing software like Aldus PageMaker including versions
 * of Lorem Ipsum.
 */
goog.provide('studash')
 
goog.require('goog.net.XhrIo');
goog.require('goog.json');
goog.require('goog.events')

goog.require('studash.Dashboard')
 
/**
 * Actions principales du site.
 * @enum {string}
 */
studash.Actions = {
  Users: '/u/',
  Auth: '/auth'
}
 
/**
 * Fonction permettant de s'authentifier au site web.
 * @param {student.Credentials} credentials Infos de connection.
 */
studash.authenticate = function(credentials) {
	goog.net.XhrIo.send(studash.Actions.Auth, function(e) {
		var resp = e.target.getResponseJson();
    if (resp.AuthResponse == true) studash.Dashboard.Create();
		else {
		
		  var span = document.getElementById('errorStr');
			while (span.firstChild) {
        span.removeChild(span.firstChild);
      }
      span.appendChild(document.createTextNode("Vos identifiants sont invalides."));
		}
		console.log(resp);
	
  }, 'POST', credentials.serialize(), {'content-type':'application/json'}, 2000);
}

goog.events.listen(document.getElementById('Sync'), goog.events.EventType.CLICK, function(e) {
  var credentials = new studash.Student.Credentials(document.forms[0][0].value,
			document.forms[0][1].value, document.forms[0][2].value);
  studash.authenticate(credentials);
});

goog.exportSymbol('studash.Actions', studash.Actions)
goog.exportSymbol('studash.authenticate', studash.authenticate)

