/**
 * Utilitaire pour intéragir avec le webservice.
 */
 
goog.provide('website')
 
goog.require('goog.net.XhrIo');
goog.require('goog.events');
goog.require('goog.json');
goog.require('student.Credentials')
 
/**
 * Actions principales du site.
 * @enum {string}
 */
website.Actions = {
  Users: '/u/',
  Auth: '/auth'
}
 
/**
 * Fonction permettant de s'authentifier au site web.
 * @param {student.Credentials} credentials Infos de connection.
 */
website.prototype.Authenticate = function(credentials) {
   var request = new goog.net.XhrIo();
	 goog.events.listen(request, "complete", ajaxCallBack);
}