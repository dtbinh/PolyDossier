/**
 * @fileoverview Base de la partie client du code.
 */
goog.provide('studash')

goog.require('studash.page.login')
goog.require('studash.Dashboard')

/**
 * Actions principales du site.
 * @enum {string}
 */
studash.pages = {
  users: '/u/',
  auth: '/auth'
}

studash.OnEnter = function() {
  studash.DoBeforeLogin();
}

studash.DoBeforeLogin = function() {
  studash.page.login.OnEnter(studash.DoAfterLogin);
}

studash.DoAfterLogin = function() {
  studash.Dashboard.Create();
}
 
studash.OnEnter();
