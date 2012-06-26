/**
 * Base code for the student object.
 */
 goog.require('goog.json');
 
 goog.provide('student')
 goog.provide('student.Credentials')
 
/**
 * Donnés de connection de l'utilisateur.
 * @param {string} username Le nom d'utilisateur.
 * @param {string} password Le mot de passe.
 * @param {string} dateOfBirth La date de naissance. (jj/mm/aa)
 * @constructor
 */
student.Credentials = function(username, password, dateOfBirth) {
	this.Username = username;
	this.Password = password;
	this.DateOfBirth = dateOfBirth;
};

/**
 * Fonction sérialisant les credentials.
 * @return {string} Serialized Json Object.
 */
student.Credentials.prototype.serialize =  function() {
  return goog.json.serialize({ 'username': this.username, 'password': this.password, 'dateOfBirth': this.dateOfBirth})
};

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
 