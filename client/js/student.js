/**
 * @fileoverview Outils important ayant lien au concept d'utilisateur.
 *
 */
goog.provide('studash.Student')
	
goog.require('goog.json');
goog.require('goog.net.XhrIo');

/**
 * Donnés de connection de l'utilisateur.
 * @param {string=} opt_username Le nom d'utilisateur. (optional)
 * @param {string=} opt_password Le mot de passe. (optional)
 * @param {string=} opt_dateOfBirth La date de naissance. (jj/mm/aa) (optional)
 * @constructor
 */
studash.Student.Credentials = function(opt_username, opt_password, opt_dateOfBirth) {
  /**
	 * Le code utilisateur demandé par le dossier.
	 * @type {string|undefined} 
	 */
	this.Username = opt_username;
	/**
	 * Le mot de passe demandé par le dossier.
	 * @type {string|undefined} 
	 */
	this.Password = opt_password;
	/**
	 * La date de naissance demandé par le dossier.
	 * @type {string|undefined} 
	 */
	this.DateOfBirth = opt_dateOfBirth;
};

/**
 * Fonction sérialisant les credentials en json.
 * @return {string} Serialized Json Object.
 */
studash.Student.Credentials.prototype.serialize =  function() {
  return goog.json.serialize({ 'username': this.Username,
      'password': this.Password, 'dateOfBirth': this.DateOfBirth})
};

/**
 * Fonction qui test l'auhtentifiation.
 */
studash.Student.Credentials.prototype.tryAuth =  function(onSuccess, onFailure) {
  if (!this.Username || !this.Password || !this.DateOfBirth) {
    onFailure();
	  return;
	}
	
	goog.net.XhrIo.send(
	  studash.pages.auth,
		function(e) {
	    if (e.target.getResponseJson().AuthResponse) onSuccess();
		  else onFailure();	
    },
		'POST', this.serialize(), {'content-type':'application/json'}, 2000
	);
};
 