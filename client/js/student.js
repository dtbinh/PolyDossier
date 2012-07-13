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
goog.provide('studash.Student')
goog.provide('studash.Student.Credentials')
	
goog.require('goog.json');
 
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
 * Fonction sérialisant les credentials.
 * @return {string} Serialized Json Object.
 */
studash.Student.Credentials.prototype.serialize =  function() {
  return goog.json.serialize({ 'username': this.Username, 'password': this.Password, 'dateOfBirth': this.DateOfBirth})
};

goog.exportSymbol('studash.Student.Credentials', studash.Student.Credentials)
goog.exportSymbol('studash.Student.Credentials.serialize', studash.Student.Credentials.serialize)

 // var Student = {
   // name        : "John Doe",
	 // username    : "malesd",
	 // password    : "rr11ee22",
	 // dateOfBirth : "890822",
	 // Credentials : function() { return { code : this.username, nip : this.password, naissance : this.dateOfBirth }},
	 // Uri         : function() { return "/u/" + this.username; },
	 // TryAuth     : function() {
	   // $.post("/auth", this.Credentials()).success(function(data) { console.log(data); });
	 // }
 // }
 