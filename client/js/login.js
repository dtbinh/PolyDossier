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

studash.page.login.ShowThrobber = function() {
  goog.dom.getElementByClass('throbber').style.visibility = 'visible';
}

studash.page.login.HideThrobber = function() {
  goog.dom.getElementByClass('throbber').style.visibility = 'hidden';
}

studash.page.login.SetErrorStr = function (str) {
  goog.dom.setTextContent(
	  goog.dom.getElementByClass('errorStr'),	str
  );
}

studash.page.login.OnRequest = function() {
  var f = document.forms[0];
	
	studash.page.login.ShowThrobber();
	studash.page.login.SetErrorStr('');
	
	(new studash.Student.Credentials(f[0].value, f[1].value, f[2].value))
	  .tryAuth(studash.page.login.OnRequestSuccess, studash.page.login.OnRequestFailure);
}

studash.page.login.OnRequestFailure = function() {
  studash.page.login.HideThrobber();
	studash.page.login.SetErrorStr(studash.str.login.error);
}

studash.page.login.OnRequestSuccess = function() {
  studash.page.login.callback();
}