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
goog.provide('studash.Website')
goog.provide('studash.Website.Proxy')
 
goog.require('goog.net.XhrIo');
goog.require('goog.events');
goog.require('goog.json');
goog.require('studash.Student.Credentials')
 
/**
 * Actions principales du site.
 * @enum {string}
 */
studash.Actions = {
  Users: '/u/',
  Auth: '/auth'
}
 
/**
 * Un objet proxy pour l'utilisation des fonctions du site.
 * @constructor
 */
studash.Website = function() {}
 
/**
 * Fonction permettant de s'authentifier au site web.
 * @param {student.Credentials} credentials Infos de connection.
 */
studash.Website.prototype.Authenticate = function(credentials) {
	goog.net.XhrIo.send(studash.Actions.Auth, function(e) {
		var xhr = e.target;
    var obj = xhr.getResponseJson();
		console.log(obj);
  }, 'POST', credentials.serialize(), {'content-type':'application/json'}, 2000);
}

// function FetchActions() {
	// $.getJSON(Student.Uri(), function(actions) {
	  
		// var nActions = actions.List.length;
		// for (i = 0; i < nActions; i++) {
	    // var fetchUri = Student.Uri() + "/" + actions.List[i];
		  // $('#Informations').append("<div id=\"" + actions.List[i] + "\">Loading...</div>"); 
			// $.get(fetchUri)
       // .success(function(data) {
			   // var action = this.url.replace(Student.Uri() + '/', '');
			   // $('#' + action).text(data);
			 // });
		// }
  // });
// }