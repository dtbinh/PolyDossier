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
		var xhr = e.target;
    var obj = xhr.getResponseJson();
		console.log(obj);
  }, 'POST', credentials.serialize(), {'content-type':'application/json'}, 2000);
}

document.getElementById('Sync').addEventListener('click', function(e) {
 console.debug(document.forms[0][0].value)
  var credentials = new studash.Student.Credentials(document.forms[0][0].value,
			document.forms[0][1].value, document.forms[0][2].value);
	console.debug(credentials)
  studash.authenticate(credentials);
})

goog.exportSymbol('studash.Actions', studash.Actions)
goog.exportSymbol('studash.authenticate', studash.authenticate)

