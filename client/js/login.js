/** @fileoverview Gestion de la page login.  */

goog.provide('studash.page.login')

goog.require('goog.dom')
goog.require('studash.str')
goog.require('studash.Student')

studash.page.login.callback = null;

studash.page.login.OnEnter = function(callback) {
  goog.events.listen(
	  document.forms[0][3], 
		goog.events.EventType.CLICK, 
		studash.page.login.OnRequest
	);
	studash.page.login.callback = callback;
}

studash.page.login.OnRequest = function() {
  var f = document.forms[0];
	(new studash.Student.Credentials(f[0].value, f[1].value, f[2].value))
	  .tryAuth(studash.page.login.OnRequestSuccess, studash.page.login.OnRequestFailure);
}

studash.page.login.OnRequestFailure = function() {
  goog.dom.setTextContent(
	  goog.dom.getElementByClass('errorStr'),
		studash.str.login.error
  );
}

studash.page.login.OnRequestSuccess = function() {
  studash.page.login.callback();
}